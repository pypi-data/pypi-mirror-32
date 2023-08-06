/* file: odbc_data_source.h */
/*******************************************************************************
* Copyright 2014-2018 Intel Corporation.
*
* This software and the related documents are Intel copyrighted  materials,  and
* your use of  them is  governed by the  express license  under which  they were
* provided to you (License).  Unless the License provides otherwise, you may not
* use, modify, copy, publish, distribute,  disclose or transmit this software or
* the related documents without Intel's prior written permission.
*
* This software and the related documents  are provided as  is,  with no express
* or implied  warranties,  other  than those  that are  expressly stated  in the
* License.
*******************************************************************************/

/*
//++
//  Implementation of the ODBC data source class
//--
*/
#ifndef __ODBC_DATA_SOURCE_H__
#define __ODBC_DATA_SOURCE_H__

#include <sstream>
#include <fstream>
#include "services/daal_memory.h"
#include "data_management/data_source/data_source.h"
#include "data_management/data/data_dictionary.h"
#include "data_management/data/numeric_table.h"
#include "data_management/data/homogen_numeric_table.h"
#include <sql.h>
#include <sqltypes.h>
#include <sqlext.h>

#include "mysql_feature_manager.h"

namespace daal
{
namespace data_management
{

/**
 * \brief Contains version 1.0 of Intel(R) Data Analytics Acceleration Library (Intel(R) DAAL) interface.
 */
namespace interface1
{
/**
 * @ingroup data_sources
 * @{
 */
/**
 * <a name="DAAL-CLASS-DATA_MANAGEMENT__ODBCDATASOURCE"></a>
 * \brief Connects to data sources with the ODBC API.
 *
 * \tparam _featureManager       Type of a data source, supports only \ref MySQLFeatureManager
 */

template<typename _featureManager, typename summaryStatisticsType = DAAL_SUMMARY_STATISTICS_TYPE>
class ODBCDataSource : public DataSourceTemplate<data_management::HomogenNumericTable<DAAL_DATA_TYPE>, summaryStatisticsType>
{
public:
    typedef _featureManager FeatureManager;

    using DataSourceIface::NumericTableAllocationFlag;
    using DataSourceIface::DictionaryCreationFlag;
    using DataSourceIface::DataSourceStatus;

    using DataSource::checkDictionary;
    using DataSource::checkNumericTable;
    using DataSource::freeNumericTable;
    using DataSource::_dict;
    using DataSource::_initialMaxRows;

protected:
    typedef data_management::HomogenNumericTable<DAAL_DATA_TYPE> DefaultNumericTableType;

    FeatureManager featureManager;

public:
    /**
     * Constructor for the ODBCDataSource class
     *
     * \param[in] dbname                Data Source Name  as  configured in settings of the ODBC driver
     * \param[in] tablename             Name of a table to export from a data source
     * \param[in] username              (optional) Username for the data source
     * \param[in] password              (optional) Password for the Username in the data source
     * \param[in] doAllocateNumericTable        (optional) Flag that specifies whether a Numeric Table
     *                                                     associated with an ODBC Data Source is allocated inside the Data Source
     * \param[in] doCreateDictionaryFromContext (optional) Flag that specifies whether a Data Dictionary
     *                                                     is created from the context of the ODBC Data Source
     * \param[in]  initialMaxRows                          Initial value of maximum number of rows in Numeric Table allocated in
     *                                                     loadDataBlock() method
     *
     */
    ODBCDataSource(const std::string &dbname, const std::string &tablename, const std::string &username = "",
                   const std::string &password = "",
                   DataSourceIface::NumericTableAllocationFlag doAllocateNumericTable    = DataSource::notAllocateNumericTable,
                   DataSourceIface::DictionaryCreationFlag doCreateDictionaryFromContext = DataSource::notDictionaryFromContext,
                   size_t initialMaxRows = 10) :
        DataSourceTemplate<DefaultNumericTableType, summaryStatisticsType>(doAllocateNumericTable, doCreateDictionaryFromContext),
        _dbname(dbname), _username(username), _password(password), _tablename(tablename),
        _idx_last_read(0), _hdlDbc(SQL_NULL_HDBC), _hdlEnv(SQL_NULL_HENV)
    {
        _query = "SELECT * FROM " + _tablename;
        _connectionStatus = DataSource::notReady;
        _initialMaxRows = initialMaxRows;
    }

    /*! \private */
    ~ODBCDataSource()
    {
        _freeHandles();
    }

    /**
     *  Frees ODBC connection handles
     */
    void freeHandles()
    {
        SQLRETURN ret = _freeHandles();
        if (!SQL_SUCCEEDED(ret))
        {
            this->_status.add(services::ErrorODBC);
        }
    }

    virtual size_t loadDataBlock(size_t maxRows) DAAL_C11_OVERRIDE
    {
        services::Status s = checkDictionary();
        if(!s) { return 0; }

        s = checkNumericTable();
        if(!s) { return 0; }

        return loadDataBlock(maxRows, this->DataSource::_spnt.get());
    }

    /**
     *  Loads a data block of a specified size into an externally allocated Numeric Table
     *  \param[in] maxRows Maximum number of rows to load from a Data Source into the Numeric Table
     *  \param nt Externally allocated Numeric Table
     *  \return Actual number of rows loaded from the Data Source
     */
    virtual size_t loadDataBlock(size_t maxRows, NumericTable *nt)
    {
        services::Status s = checkDictionary();
        if(!s) { return 0; }

        if( nt == NULL ) { this->_status.add(services::throwIfPossible(services::ErrorNullInputNumericTable)); return 0; }

        DataSourceTemplate<DefaultNumericTableType, summaryStatisticsType>::resizeNumericTableImpl( maxRows, nt );

        if(nt->getDataMemoryStatus() == NumericTableIface::userAllocated)
        {
            if(nt->getNumberOfRows() < maxRows)
            {
                this->_status.add(services::throwIfPossible(services::ErrorIncorrectNumberOfObservations));
                return 0;
            }
            if(nt->getNumberOfColumns() != _dict->getNumberOfFeatures())
            {
                this->_status.add(services::throwIfPossible(services::ErrorIncorrectNumberOfFeatures));
                return 0;
            }
        }

        std::string query_exec = featureManager.setLimitQuery(_query, _idx_last_read, maxRows);
        SQLRETURN ret;
        ret = _establishHandles();
        if (!SQL_SUCCEEDED(ret)) { this->_status.add(services::throwIfPossible(services::ErrorHandlesSQL)); return 0; }

        SQLHSTMT hdlStmt = SQL_NULL_HSTMT;
        ret = SQLAllocHandle(SQL_HANDLE_STMT, _hdlDbc, &hdlStmt);
        if (!SQL_SUCCEEDED(ret)) { this->_status.add(services::throwIfPossible(services::ErrorSQLstmtHandle)); return 0; }

        ret = SQLExecDirect(hdlStmt, (SQLCHAR *)query_exec.c_str(), SQL_NTS);
        if (!SQL_SUCCEEDED(ret)) { this->_status.add(services::throwIfPossible(services::ErrorODBC)); return 0; }

        DataSourceIface::DataSourceStatus dataSourceStatus;

        dataSourceStatus = featureManager.statementResultsNumericTable(hdlStmt, nt, maxRows);

        size_t nRead = nt->getNumberOfRows();
        _idx_last_read += nRead;

        if(nt->basicStatistics.get(NumericTableIface::minimum   ).get() != NULL &&
           nt->basicStatistics.get(NumericTableIface::maximum   ).get() != NULL &&
           nt->basicStatistics.get(NumericTableIface::sum       ).get() != NULL &&
           nt->basicStatistics.get(NumericTableIface::sumSquares).get() != NULL)
        {
            for(size_t i = 0; i < nRead; i++)
            {
                DataSourceTemplate<DefaultNumericTableType, summaryStatisticsType>::updateStatistics( i, nt );
            }
        }

        ret = SQLFreeHandle(SQL_HANDLE_STMT, hdlStmt);
        if (!SQL_SUCCEEDED(ret)) { this->_status.add(services::throwIfPossible(services::ErrorSQLstmtHandle)); return 0; }

        if (dataSourceStatus == DataSource::endOfData) { _connectionStatus = DataSource::endOfData; }

        NumericTableDictionaryPtr ntDict = nt->getDictionarySharedPtr();
        size_t nFeatures = _dict->getNumberOfFeatures();
        ntDict->setNumberOfFeatures(nFeatures);
        for (size_t i = 0; i < nFeatures; i++)
        {
            ntDict->setFeature((*_dict)[i].ntFeature, i);
        }

        return nRead;
    }

    size_t loadDataBlock() DAAL_C11_OVERRIDE
    {
        services::Status s = checkDictionary();
        if(!s) { return 0; }

        s = checkNumericTable();
        if(!s) { return 0; }

        return loadDataBlock(this->DataSource::_spnt.get());
    }

    size_t loadDataBlock(NumericTable* nt) DAAL_C11_OVERRIDE
    {
        services::Status s = checkDictionary();
        if(!s) { return 0; }

        if( nt == NULL ) {this->_status.add(services::throwIfPossible(services::ErrorNullInputNumericTable)); return 0; }

        size_t maxRows = (_initialMaxRows > 0 ? _initialMaxRows : 10);
        size_t nrows = 0;
        size_t ncols = _dict->getNumberOfFeatures();

        DataCollection tables;

        for( ; ; )
        {
            NumericTablePtr ntCurrent = HomogenNumericTable<DAAL_DATA_TYPE>::create(ncols, maxRows, NumericTableIface::doAllocate, &s);
            if (!s)
            {
                this->_status.add(services::throwIfPossible(services::ErrorNumericTableNotAllocated));
                break;
            }
            tables.push_back(ntCurrent);
            size_t rows = loadDataBlock(maxRows, ntCurrent.get());
            nrows += rows;
            if (rows < maxRows) { break; }
            maxRows *= 2;
        }

        DataSourceTemplate<DefaultNumericTableType, summaryStatisticsType>::resizeNumericTableImpl( nrows, nt );
        nt->setNormalizationFlag(NumericTable::nonNormalized);

        BlockDescriptor<DAAL_DATA_TYPE> blockCurrent, block;

        size_t pos = 0;

        for (size_t i = 0; i < tables.size(); i++) {
            NumericTable *ntCurrent = (NumericTable*)(tables[i].get());
            size_t rows = ntCurrent->getNumberOfRows();

            if (rows == 0) { continue; }

            ntCurrent->getBlockOfRows(0, rows, readOnly, blockCurrent);
            nt->getBlockOfRows(pos, rows, writeOnly, block);

            services::daal_memcpy_s(block.getBlockPtr(), rows * ncols * sizeof(DAAL_DATA_TYPE), blockCurrent.getBlockPtr(), rows * ncols * sizeof(DAAL_DATA_TYPE));

            ntCurrent->releaseBlockOfRows(blockCurrent);
            nt->releaseBlockOfRows(block);

            DataSourceTemplate<DefaultNumericTableType, summaryStatisticsType>::combineStatistics( ntCurrent, nt, pos == 0);
            pos += rows;
        }

        NumericTableDictionaryPtr ntDict = nt->getDictionarySharedPtr();
        size_t nFeatures = _dict->getNumberOfFeatures();
        ntDict->setNumberOfFeatures(nFeatures);
        for (size_t i = 0; i < nFeatures; i++)
        {
            ntDict->setFeature((*_dict)[i].ntFeature, i);
        }

        return nrows;
    }

    services::Status createDictionaryFromContext() DAAL_C11_OVERRIDE
    {
        if (_dict) { return services::throwIfPossible(services::Status(services::ErrorDictionaryAlreadyAvailable)); }

        std::string query_exec = _query + ";";

        SQLRETURN ret;
        ret = _establishHandles();
        if (!SQL_SUCCEEDED(ret)) { return services::throwIfPossible(services::Status(services::ErrorHandlesSQL)); }

        SQLHSTMT hdlStmt = SQL_NULL_HSTMT;
        ret = SQLAllocHandle(SQL_HANDLE_STMT, _hdlDbc, &hdlStmt);
        if (!SQL_SUCCEEDED(ret)) { return services::throwIfPossible(services::Status(services::ErrorSQLstmtHandle)); }

        ret = SQLPrepare(hdlStmt, (SQLCHAR *)query_exec.c_str(), SQL_NTS);
        if (!SQL_SUCCEEDED(ret))
        {
            return services::throwIfPossible(services::Status(services::ErrorODBC));
        }

        services::Status status;
        _dict = DataSourceDictionary::create(&status);
        if (!status) return status;

        featureManager.createDictionary(hdlStmt, this->_dict.get());

        ret = SQLFreeHandle(SQL_HANDLE_STMT, hdlStmt);
        if (!SQL_SUCCEEDED(ret)) { return services::throwIfPossible(services::Status(services::ErrorSQLstmtHandle)); }

        if (featureManager.getErrors()->size() == 0) { _connectionStatus = DataSource::readyForLoad; }
        return status;
    }

    DataSourceIface::DataSourceStatus getStatus() DAAL_C11_OVERRIDE
    {
        return _connectionStatus;
    }

    size_t getNumberOfAvailableRows() DAAL_C11_OVERRIDE
    {
        SQLLEN nRows = 0;
        SQLRETURN ret;

        ret = _establishHandles();
        if (!SQL_SUCCEEDED(ret)) { return 0; }

        SQLHSTMT hdlStmt = SQL_NULL_HSTMT;
        ret = SQLAllocHandle(SQL_HANDLE_STMT, _hdlDbc, &hdlStmt);
        if (!SQL_SUCCEEDED(ret)) { return 0; }

        std::string query_exec = "SELECT COUNT(*) FROM " + _tablename + " ;";
        ret = SQLExecDirect(hdlStmt, (SQLCHAR *)query_exec.c_str(), SQL_NTS);
        if (!SQL_SUCCEEDED(ret)) { return 0; }

        ret = SQLBindCol(hdlStmt, 1, SQL_C_ULONG, (SQLPOINTER)&nRows, 0, NULL);
        if (!SQL_SUCCEEDED(ret)) { return 0; }

        ret = SQLFetchScroll(hdlStmt, SQL_FETCH_NEXT, 1);
        if (!SQL_SUCCEEDED(ret)) { return 0; }

        ret = SQLFreeHandle(SQL_HANDLE_STMT, hdlStmt);
        if (!SQL_SUCCEEDED(ret)) { return ret; }

        return nRows;
    }

    FeatureManager &getFeatureManager()
    {
        return featureManager;
    }

private:
    std::string      _dbname;
    std::string      _username;
    std::string      _password;
    std::string      _tablename;
    std::string      _query;
    size_t           _idx_last_read;
    DataSourceIface::DataSourceStatus _connectionStatus;

    SQLHENV _hdlEnv;
    SQLHDBC _hdlDbc;

    SQLRETURN _establishHandles()
    {
        if (_hdlEnv == SQL_NULL_HENV || _hdlDbc == SQL_NULL_HDBC)
        {
            SQLRETURN ret;
            ret = SQLAllocHandle(SQL_HANDLE_ENV, SQL_NULL_HANDLE, &_hdlEnv);
            if (!SQL_SUCCEEDED(ret)) { return ret; }

            ret = SQLSetEnvAttr(_hdlEnv, SQL_ATTR_ODBC_VERSION, (SQLPOINTER) SQL_OV_ODBC3, SQL_IS_UINTEGER);
            if (!SQL_SUCCEEDED(ret)) { return ret; }

            ret = SQLAllocHandle(SQL_HANDLE_DBC, _hdlEnv, &_hdlDbc);
            if (!SQL_SUCCEEDED(ret)) { return ret; }

            SQLCHAR *usernamePtr;
            SQLCHAR *passwordPtr;
            SQLSMALLINT usernameLen;
            SQLSMALLINT passwordLen;
            if (_username.empty())
            {
                usernamePtr = NULL;
                usernameLen = 0;
            }
            else
            {
                usernamePtr = (SQLCHAR *)_username.c_str();
                usernameLen = SQL_NTS;
            }
            if (_password.empty())
            {
                passwordPtr = NULL;
                passwordLen = 0;
            }
            else
            {
                passwordPtr = (SQLCHAR *)_password.c_str();
                passwordLen = SQL_NTS;
            }
            ret = SQLConnect(_hdlDbc, (SQLCHAR *)_dbname.c_str(), SQL_NTS, usernamePtr, usernameLen, passwordPtr, passwordLen);
            if (!SQL_SUCCEEDED(ret)) { return ret; }
        }
        return SQL_SUCCESS;
    }

    SQLRETURN _freeHandles()
    {
        if (_hdlDbc == SQL_NULL_HDBC || _hdlEnv == SQL_NULL_HENV) { return SQL_SUCCESS; }

        SQLRETURN ret;

        ret = SQLDisconnect(_hdlDbc);
        if (!SQL_SUCCEEDED(ret)) { return ret; }

        ret = SQLFreeHandle(SQL_HANDLE_DBC, _hdlDbc);
        if (!SQL_SUCCEEDED(ret)) { return ret; }

        ret = SQLFreeHandle(SQL_HANDLE_ENV, _hdlEnv);
        if (!SQL_SUCCEEDED(ret)) { return ret; }

        _hdlDbc = SQL_NULL_HDBC;
        _hdlEnv = SQL_NULL_HENV;

        return SQL_SUCCESS;
    }
};
/** @} */
} // namespace interface1
using interface1::ODBCDataSource;

}
}
#endif
