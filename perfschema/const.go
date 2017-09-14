// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package perfschema

// Performance Schema Name.
const (
	Name = "PERFORMANCE_SCHEMA"
)

// Definition order same as MySQL's reference manual, so don't bother to
// adjust according to alphabetical order.
const (
	TableGlobalStatus           = "GLOBAL_STATUS"
	TableSessionStatus          = "SESSION_STATUS"
	TableSetupActors            = "SETUP_ACTORS"
	TableSetupObjects           = "SETUP_OBJECTS"
	TableSetupInstruments       = "SETUP_INSTRUMENTS"
	TableSetupConsumers         = "SETUP_CONSUMERS"
	TableSetupTimers            = "SETUP_TIMERS"
	TableStmtsCurrent           = "EVENTS_STATEMENTS_CURRENT"
	TableStmtsHistory           = "EVENTS_STATEMENTS_HISTORY"
	TableStmtsHistoryLong       = "EVENTS_STATEMENTS_HISTORY_LONG"
	TablePreparedStmtsInstances = "PREPARED_STATEMENTS_INSTANCES"
	TableTransCurrent           = "EVENTS_TRANSACTIONS_CURRENT"
	TableTransHistory           = "EVENTS_TRANSACTIONS_HISTORY"
	TableTransHistoryLong       = "EVENTS_TRANSACTIONS_HISTORY_LONG"
	TableStagesCurrent          = "EVENTS_STAGES_CURRENT"
	TableStagesHistory          = "EVENTS_STAGES_HISTORY"
	TableStagesHistoryLong      = "EVENTS_STAGES_HISTORY_LONG"
	TableSessionStatus          = "SESSION_STATUS"
	TableGlobalStatus           = "GLOBAL_STATUS"
)

// PerfSchemaTables is a shortcut to involve all table names.
var PerfSchemaTables = []string{
	TableGlobalStatus,
	TableSessionStatus,
	TableSetupActors,
	TableSetupObjects,
	TableSetupInstruments,
	TableSetupConsumers,
	TableSetupTimers,
	TableStmtsCurrent,
	TableStmtsHistory,
	TableStmtsHistoryLong,
	TablePreparedStmtsInstances,
	TableTransCurrent,
	TableTransHistory,
	TableTransHistoryLong,
	TableStagesCurrent,
	TableStagesHistory,
	TableStagesHistoryLong,
	TableSessionStatus,
	TableGlobalStatus,
}

// ColumnGlobalStatus contains the column name definitions for table global_status, same as MySQL.
//
// CREATE TABLE performance_schema.global_status(
//     VARIABLE_NAME VARCHAR(64) not null,
//     VARIABLE_VALUE VARCHAR(1024));
var ColumnGlobalStatus = []string{"VARIABLE_NAME", "VARIABLE_VALUE"}

// ColumnSessionStatus contains the column name definitions for table session_status, same as MySQL.
//
// CREATE TABLE performance_schema.session_status(
//     VARIABLE_NAME VARCHAR(64) not null,
//     VARIABLE_VALUE VARCHAR(1024));
var ColumnSessionStatus = []string{"VARIABLE_NAME", "VARIABLE_VALUE"}

// ColumnSetupActors contains the column name definitions for table setup_actors, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.setup_actors (
// 		HOST			CHAR(60) NOT NULL  DEFAULT '%',
// 		USER			CHAR(32) NOT NULL  DEFAULT '%',
// 		ROLE			CHAR(16) NOT NULL  DEFAULT '%',
// 		ENABLED			ENUM('YES','NO') NOT NULL  DEFAULT 'YES',
// 		HISTORY			ENUM('YES','NO') NOT NULL  DEFAULT 'YES');
var ColumnSetupActors = []string{"HOST", "USER", "ROLE", "ENABLED", "HISTORY"}

// ColumnSetupObjects contains the column name definitions for table setup_objects, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.setup_objects (
// 		OBJECT_TYPE		ENUM('EVENT','FUNCTION','TABLE') NOT NULL  DEFAULT 'TABLE',
// 		OBJECT_SCHEMA	VARCHAR(64)  DEFAULT '%',
// 		OBJECT_NAME		VARCHAR(64) NOT NULL  DEFAULT '%',
// 		ENABLED			ENUM('YES','NO') NOT NULL  DEFAULT 'YES',
// 		TIMED			ENUM('YES','NO') NOT NULL  DEFAULT 'YES');
var ColumnSetupObjects = []string{"OBJECT_TYPE", "OBJECT_SCHEMA", "OBJECT_NAME", "ENABLED", "TIMED"}

// ColumnSetupInstruments contains the column name definitions for table setup_instruments, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.setup_instruments (
// 		NAME			VARCHAR(128) NOT NULL,
// 		ENABLED			ENUM('YES','NO') NOT NULL,
// 		TIMED			ENUM('YES','NO') NOT NULL);
var ColumnSetupInstruments = []string{"NAMED", "ENABLED", "TIMED"}

// ColumnSetupConsumers contains the column name definitions for table setup_consumers, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.setup_consumers (
// 		NAME			VARCHAR(64) NOT NULL,
// 		ENABLED			ENUM('YES','NO') NOT NULL);
var ColumnSetupConsumers = []string{"NAMED", "ENABLED"}

// ColumnSetupTimers contains the column name definitions for table setup_timers, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.setup_timers (
// 		NAME			VARCHAR(64) NOT NULL,
// 		TIMER_NAME		ENUM('NANOSECOND','MICROSECOND','MILLISECOND') NOT NULL);
var ColumnSetupTimers = []string{"NAME", "TIMER_NAME"}

// ColumnStmtsCurrent contains the column name definitions for table events_statements_current, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_statements_current (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		LOCK_TIME		BIGINT(20) UNSIGNED NOT NULL,
// 		SQL_TEXT		LONGTEXT,
// 		DIGEST			VARCHAR(32),
// 		DIGEST_TEXT		LONGTEXT,
// 		CURRENT_SCHEMA	VARCHAR(64),
// 		OBJECT_TYPE		VARCHAR(64),
// 		OBJECT_SCHEMA	VARCHAR(64),
// 		OBJECT_NAME		VARCHAR(64),
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED,
// 		MYSQL_ERRNO		INT(11),
// 		RETURNED_SQLSTATE		VARCHAR(5),
// 		MESSAGE_TEXT	VARCHAR(128),
// 		ERRORS			BIGINT(20) UNSIGNED NOT NULL,
// 		WARNINGS		BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_AFFECTED	BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_SENT		BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_EXAMINED	BIGINT(20) UNSIGNED NOT NULL,
// 		CREATED_TMP_DISK_TABLES	BIGINT(20) UNSIGNED NOT NULL,
// 		CREATED_TMP_TABLES		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_FULL_JOIN		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_FULL_RANGE_JOIN	BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_RANGE	BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_RANGE_CHECK		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_SCAN		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_MERGE_PASSES		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_RANGE		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_ROWS		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_SCAN		BIGINT(20) UNSIGNED NOT NULL,
// 		NO_INDEX_USED	BIGINT(20) UNSIGNED NOT NULL,
// 		NO_GOOD_INDEX_USED		BIGINT(20) UNSIGNED NOT NULL,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'),
// 		NESTING_EVENT_LEVEL		INT(11));
var ColumnStmtsCurrent = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"LOCK_TIME",
	"SQL_TEXT",
	"DIGEST",
	"DIGEST_TEXT",
	"CURRENT_SCHEMA",
	"OBJECT_TYPE",
	"OBJECT_SCHEMA",
	"OBJECT_NAME",
	"OBJECT_INSTANCE_BEGIN",
	"MYSQL_ERRNO",
	"RETURNED_SQLSTATE",
	"MESSAGE_TEXT",
	"ERRORS",
	"WARNINGS",
	"ROWS_AFFECTED",
	"ROWS_SENT",
	"ROWS_EXAMINED",
	"CREATED_TMP_DISK_TABLES",
	"CREATED_TMP_TABLES",
	"SELECT_FULL_JOIN",
	"SELECT_FULL_RANGE_JOIN",
	"SELECT_RANGE",
	"SELECT_RANGE_CHECK",
	"SELECT_SCAN",
	"SORT_MERGE_PASSES",
	"SORT_RANGE",
	"SORT_ROWS",
	"SORT_SCAN",
	"NO_INDEX_USED",
	"NO_GOOD_INDEX_USED",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
	"NESTING_EVENT_LEVEL",
}

// ColumnStmtsHistory contains the column name definitions for table events_statements_history, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_statements_history (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		LOCK_TIME		BIGINT(20) UNSIGNED NOT NULL,
// 		SQL_TEXT		LONGTEXT,
// 		DIGEST			VARCHAR(32),
// 		DIGEST_TEXT		LONGTEXT,
// 		CURRENT_SCHEMA	VARCHAR(64),
// 		OBJECT_TYPE		VARCHAR(64),
// 		OBJECT_SCHEMA	VARCHAR(64),
// 		OBJECT_NAME		VARCHAR(64),
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED,
// 		MYSQL_ERRNO		INT(11),
// 		RETURNED_SQLSTATE		VARCHAR(5),
// 		MESSAGE_TEXT	VARCHAR(128),
// 		ERRORS			BIGINT(20) UNSIGNED NOT NULL,
// 		WARNINGS		BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_AFFECTED	BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_SENT		BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_EXAMINED	BIGINT(20) UNSIGNED NOT NULL,
// 		CREATED_TMP_DISK_TABLES	BIGINT(20) UNSIGNED NOT NULL,
// 		CREATED_TMP_TABLES		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_FULL_JOIN		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_FULL_RANGE_JOIN	BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_RANGE	BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_RANGE_CHECK		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_SCAN		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_MERGE_PASSES		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_RANGE		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_ROWS		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_SCAN		BIGINT(20) UNSIGNED NOT NULL,
// 		NO_INDEX_USED	BIGINT(20) UNSIGNED NOT NULL,
// 		NO_GOOD_INDEX_USED		BIGINT(20) UNSIGNED NOT NULL,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'),
// 		NESTING_EVENT_LEVEL		INT(11));
var ColumnStmtsHistory = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"LOCK_TIME",
	"SQL_TEXT",
	"DIGEST",
	"DIGEST_TEXT",
	"CURRENT_SCHEMA",
	"OBJECT_TYPE",
	"OBJECT_SCHEMA",
	"OBJECT_NAME",
	"OBJECT_INSTANCE_BEGIN",
	"MYSQL_ERRNO",
	"RETURNED_SQLSTATE",
	"MESSAGE_TEXT",
	"ERRORS",
	"WARNINGS",
	"ROWS_AFFECTED",
	"ROWS_SENT",
	"ROWS_EXAMINED",
	"CREATED_TMP_DISK_TABLES",
	"CREATED_TMP_TABLES",
	"SELECT_FULL_JOIN",
	"SELECT_FULL_RANGE_JOIN",
	"SELECT_RANGE",
	"SELECT_RANGE_CHECK",
	"SELECT_SCAN",
	"SORT_MERGE_PASSES",
	"SORT_RANGE",
	"SORT_ROWS",
	"SORT_SCAN",
	"NO_INDEX_USED",
	"NO_GOOD_INDEX_USED",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
	"NESTING_EVENT_LEVEL",
}

// ColumnStmtsHistoryLong contains the column name definitions for table events_statements_history_long, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_statements_history_long (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		LOCK_TIME		BIGINT(20) UNSIGNED NOT NULL,
// 		SQL_TEXT		LONGTEXT,
// 		DIGEST			VARCHAR(32),
// 		DIGEST_TEXT		LONGTEXT,
// 		CURRENT_SCHEMA	VARCHAR(64),
// 		OBJECT_TYPE		VARCHAR(64),
// 		OBJECT_SCHEMA	VARCHAR(64),
// 		OBJECT_NAME		VARCHAR(64),
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED,
// 		MYSQL_ERRNO		INT(11),
// 		RETURNED_SQLSTATE		VARCHAR(5),
// 		MESSAGE_TEXT	VARCHAR(128),
// 		ERRORS			BIGINT(20) UNSIGNED NOT NULL,
// 		WARNINGS		BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_AFFECTED	BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_SENT		BIGINT(20) UNSIGNED NOT NULL,
// 		ROWS_EXAMINED	BIGINT(20) UNSIGNED NOT NULL,
// 		CREATED_TMP_DISK_TABLES	BIGINT(20) UNSIGNED NOT NULL,
// 		CREATED_TMP_TABLES		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_FULL_JOIN		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_FULL_RANGE_JOIN	BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_RANGE	BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_RANGE_CHECK		BIGINT(20) UNSIGNED NOT NULL,
// 		SELECT_SCAN		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_MERGE_PASSES		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_RANGE		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_ROWS		BIGINT(20) UNSIGNED NOT NULL,
// 		SORT_SCAN		BIGINT(20) UNSIGNED NOT NULL,
// 		NO_INDEX_USED	BIGINT(20) UNSIGNED NOT NULL,
// 		NO_GOOD_INDEX_USED		BIGINT(20) UNSIGNED NOT NULL,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'),
// 		NESTING_EVENT_LEVEL		INT(11));
var ColumnStmtsHistoryLong = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"LOCK_TIME",
	"SQL_TEXT",
	"DIGEST",
	"DIGEST_TEXT",
	"CURRENT_SCHEMA",
	"OBJECT_TYPE",
	"OBJECT_SCHEMA",
	"OBJECT_NAME",
	"OBJECT_INSTANCE_BEGIN",
	"MYSQL_ERRNO",
	"RETURNED_SQLSTATE",
	"MESSAGE_TEXT",
	"ERRORS",
	"WARNINGS",
	"ROWS_AFFECTED",
	"ROWS_SENT",
	"ROWS_EXAMINED",
	"CREATED_TMP_DISK_TABLES",
	"CREATED_TMP_TABLES",
	"SELECT_FULL_JOIN",
	"SELECT_FULL_RANGE_JOIN",
	"SELECT_RANGE",
	"SELECT_RANGE_CHECK",
	"SELECT_SCAN",
	"SORT_MERGE_PASSES",
	"SORT_RANGE",
	"SORT_ROWS",
	"SORT_SCAN",
	"NO_INDEX_USED",
	"NO_GOOD_INDEX_USED",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
	"NESTING_EVENT_LEVEL",
}

// ColumnPreparedStmtsInstances contains the column name definitions for table prepared_statements_instances, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.prepared_statements_instances (
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED NOT NULL,
// 		STATEMENT_ID	BIGINT(20) UNSIGNED NOT NULL,
// 		STATEMENT_NAME	VARCHAR(64),
// 		SQL_TEXT		LONGTEXT NOT NULL,
// 		OWNER_THREAD_ID	BIGINT(20) UNSIGNED NOT NULL,
// 		OWNER_EVENT_ID	BIGINT(20) UNSIGNED NOT NULL,
// 		OWNER_OBJECT_TYPE		ENUM('EVENT','FUNCTION','TABLE'),
// 		OWNER_OBJECT_SCHEMA		VARCHAR(64),
// 		OWNER_OBJECT_NAME		VARCHAR(64),
// 		TIMER_PREPARE	BIGINT(20) UNSIGNED NOT NULL,
// 		COUNT_REPREPARE	BIGINT(20) UNSIGNED NOT NULL,
// 		COUNT_EXECUTE	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_TIMER_EXECUTE		BIGINT(20) UNSIGNED NOT NULL,
// 		MIN_TIMER_EXECUTE		BIGINT(20) UNSIGNED NOT NULL,
// 		AVG_TIMER_EXECUTE		BIGINT(20) UNSIGNED NOT NULL,
// 		MAX_TIMER_EXECUTE		BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_LOCK_TIME	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_ERRORS		BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_WARNINGS	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_ROWS_AFFECTED		BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_ROWS_SENT	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_ROWS_EXAMINED		BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_CREATED_TMP_DISK_TABLES	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_CREATED_TMP_TABLES	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SELECT_FULL_JOIN	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SELECT_FULL_RANGE_JOIN	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SELECT_RANGE		BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SELECT_RANGE_CHECK	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SELECT_SCAN	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SORT_MERGE_PASSES	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SORT_RANGE	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SORT_ROWS	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_SORT_SCAN	BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_NO_INDEX_USED		BIGINT(20) UNSIGNED NOT NULL,
// 		SUM_NO_GOOD_INDEX_USED	BIGINT(20) UNSIGNED NOT NULL);
var ColumnPreparedStmtsInstances = []string{
	"OBJECT_INSTANCE_BEGIN",
	"STATEMENT_ID",
	"STATEMENT_NAME",
	"SQL_TEXT",
	"OWNER_THREAD_ID",
	"OWNER_EVENT_ID",
	"OWNER_OBJECT_TYPE",
	"OWNER_OBJECT_SCHEMA",
	"OWNER_OBJECT_NAME",
	"TIMER_PREPARE",
	"COUNT_REPREPARE",
	"COUNT_EXECUTE",
	"SUM_TIMER_EXECUTE",
	"MIN_TIMER_EXECUTE",
	"AVG_TIMER_EXECUTE",
	"MAX_TIMER_EXECUTE",
	"SUM_LOCK_TIME",
	"SUM_ERRORS",
	"SUM_WARNINGS",
	"SUM_ROWS_AFFECTED",
	"SUM_ROWS_SENT",
	"SUM_ROWS_EXAMINED",
	"SUM_CREATED_TMP_DISK_TABLES",
	"SUM_CREATED_TMP_TABLES",
	"SUM_SELECT_FULL_JOIN",
	"SUM_SELECT_FULL_RANGE_JOIN",
	"SUM_SELECT_RANGE",
	"SUM_SELECT_RANGE_CHECK",
	"SUM_SELECT_SCAN",
	"SUM_SORT_MERGE_PASSES",
	"SUM_SORT_RANGE",
	"SUM_SORT_ROWS",
	"SUM_SORT_SCAN",
	"SUM_NO_INDEX_USED",
	"SUM_NO_GOOD_INDEX_USED",
}

// ColumnTransCurrent contains the column name definitions for table events_transactions_current, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_transactions_current (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		STATE			ENUM('ACTIVE','COMMITTED',"ROLLED BACK"),
// 		TRX_ID			BIGINT(20) UNSIGNED,
// 		GTID			VARCHAR(64),
// 		XID_FORMAT_ID	INT(11),
// 		XID_GTRID		VARCHAR(130),
// 		XID_BQUAL		VARCHAR(130),
// 		XA_STATE		VARCHAR(64),
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		ACCESS_MODE		ENUM('READ ONLY','READ WRITE'),
// 		ISOLATION_LEVEL	VARCHAR(64),
// 		AUTOCOMMIT		ENUM('YES','NO') NOT NULL,
// 		NUMBER_OF_SAVEPOINTS	BIGINT(20) UNSIGNED,
// 		NUMBER_OF_ROLLBACK_TO_SAVEPOINT	BIGINT(20) UNSIGNED,
// 		NUMBER_OF_RELEASE_SAVEPOINT		BIGINT(20) UNSIGNED,
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'));
var ColumnTransCurrent = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"STATE",
	"TRX_ID",
	"GTID",
	"XID_FORMAT_ID",
	"XID_GTRID",
	"XID_BQUAL",
	"XA_STATE",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"ACCESS_MODE",
	"ISOLATION_LEVEL",
	"AUTOCOMMIT",
	"NUMBER_OF_SAVEPOINTS",
	"NUMBER_OF_ROLLBACK_TO_SAVEPOINT",
	"NUMBER_OF_RELEASE_SAVEPOINT",
	"OBJECT_INSTANCE_BEGIN",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
}

// ColumnTransHistory contains the column name definitions for table events_transactions_history, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_transactions_history (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		STATE			ENUM('ACTIVE','COMMITTED',"ROLLED BACK"),
// 		TRX_ID			BIGINT(20) UNSIGNED,
// 		GTID			VARCHAR(64),
// 		XID_FORMAT_ID	INT(11),
// 		XID_GTRID		VARCHAR(130),
// 		XID_BQUAL		VARCHAR(130),
// 		XA_STATE		VARCHAR(64),
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		ACCESS_MODE		ENUM('READ ONLY','READ WRITE'),
// 		ISOLATION_LEVEL	VARCHAR(64),
// 		AUTOCOMMIT		ENUM('YES','NO') NOT NULL,
// 		NUMBER_OF_SAVEPOINTS	BIGINT(20) UNSIGNED,
// 		NUMBER_OF_ROLLBACK_TO_SAVEPOINT	BIGINT(20) UNSIGNED,
// 		NUMBER_OF_RELEASE_SAVEPOINT		BIGINT(20) UNSIGNED,
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'));
var ColumnTransHistory = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"STATE",
	"TRX_ID",
	"GTID",
	"XID_FORMAT_ID",
	"XID_GTRID",
	"XID_BQUAL",
	"XA_STATE",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"ACCESS_MODE",
	"ISOLATION_LEVEL",
	"AUTOCOMMIT",
	"NUMBER_OF_SAVEPOINTS",
	"NUMBER_OF_ROLLBACK_TO_SAVEPOINT",
	"NUMBER_OF_RELEASE_SAVEPOINT",
	"OBJECT_INSTANCE_BEGIN",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
}

// ColumnTransHistoryLong contains the column name definitions for table events_transactions_history_long, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_transactions_history_long (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		STATE			ENUM('ACTIVE','COMMITTED',"ROLLED BACK"),
// 		TRX_ID			BIGINT(20) UNSIGNED,
// 		GTID			VARCHAR(64),
// 		XID_FORMAT_ID	INT(11),
// 		XID_GTRID		VARCHAR(130),
// 		XID_BQUAL		VARCHAR(130),
// 		XA_STATE		VARCHAR(64),
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		ACCESS_MODE		ENUM('READ ONLY','READ WRITE'),
// 		ISOLATION_LEVEL	VARCHAR(64),
// 		AUTOCOMMIT		ENUM('YES','NO') NOT NULL,
// 		NUMBER_OF_SAVEPOINTS	BIGINT(20) UNSIGNED,
// 		NUMBER_OF_ROLLBACK_TO_SAVEPOINT	BIGINT(20) UNSIGNED,
// 		NUMBER_OF_RELEASE_SAVEPOINT		BIGINT(20) UNSIGNED,
// 		OBJECT_INSTANCE_BEGIN	BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'));
var ColumnTransHistoryLong = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"STATE",
	"TRX_ID",
	"GTID",
	"XID_FORMAT_ID",
	"XID_GTRID",
	"XID_BQUAL",
	"XA_STATE",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"ACCESS_MODE",
	"ISOLATION_LEVEL",
	"AUTOCOMMIT",
	"NUMBER_OF_SAVEPOINTS",
	"NUMBER_OF_ROLLBACK_TO_SAVEPOINT",
	"NUMBER_OF_RELEASE_SAVEPOINT",
	"OBJECT_INSTANCE_BEGIN",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
}

// ColumnStagesCurrent contains the column name definitions for table events_stages_current, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_stages_current (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		WORK_COMPLETED	BIGINT(20) UNSIGNED,
// 		WORK_ESTIMATED	BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'));
var ColumnStagesCurrent = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"WORK_COMPLETED",
	"WORK_ESTIMATED",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
}

// ColumnStagesHistory contains the column name definitions for table events_stages_history, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_stages_history (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		WORK_COMPLETED	BIGINT(20) UNSIGNED,
// 		WORK_ESTIMATED	BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'));
var ColumnStagesHistory = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"WORK_COMPLETED",
	"WORK_ESTIMATED",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
}

// ColumnStagesHistoryLong contains the column name definitions for table events_stages_history_long, same as MySQL.
//
// CREATE TABLE if not exists performance_schema.events_stages_history_long (
// 		THREAD_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		EVENT_ID		BIGINT(20) UNSIGNED NOT NULL,
// 		END_EVENT_ID	BIGINT(20) UNSIGNED,
// 		EVENT_NAME		VARCHAR(128) NOT NULL,
// 		SOURCE			VARCHAR(64),
// 		TIMER_START		BIGINT(20) UNSIGNED,
// 		TIMER_END		BIGINT(20) UNSIGNED,
// 		TIMER_WAIT		BIGINT(20) UNSIGNED,
// 		WORK_COMPLETED	BIGINT(20) UNSIGNED,
// 		WORK_ESTIMATED	BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_ID		BIGINT(20) UNSIGNED,
// 		NESTING_EVENT_TYPE		ENUM('TRANSACTION','STATEMENT','STAGE'));
var ColumnStagesHistoryLong = []string{
	"THREAD_ID",
	"EVENT_ID",
	"END_EVENT_ID",
	"EVENT_NAME",
	"SOURCE",
	"TIMER_START",
	"TIMER_END",
	"TIMER_WAIT",
	"WORK_COMPLETED",
	"WORK_ESTIMATED",
	"NESTING_EVENT_ID",
	"NESTING_EVENT_TYPE",
}
