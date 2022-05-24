// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	context "context"
	sql "database/sql"
	sync "sync"

	db "gitee.com/zhaochuninhefei/fabric-ca-gm/lib/server/db"
	sqlx "github.com/jmoiron/sqlx"
)

type FabricCADB struct {
	BeginTxStub        func() db.FabricCATx
	beginTxMutex       sync.RWMutex
	beginTxArgsForCall []struct {
	}
	beginTxReturns struct {
		result1 db.FabricCATx
	}
	beginTxReturnsOnCall map[int]struct {
		result1 db.FabricCATx
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DriverNameStub        func() string
	driverNameMutex       sync.RWMutex
	driverNameArgsForCall []struct {
	}
	driverNameReturns struct {
		result1 string
	}
	driverNameReturnsOnCall map[int]struct {
		result1 string
	}
	ExecStub        func(string, string, ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	GetStub        func(string, interface{}, string, ...interface{}) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 interface{}
		arg3 string
		arg4 []interface{}
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	IsInitializedStub        func() bool
	isInitializedMutex       sync.RWMutex
	isInitializedArgsForCall []struct {
	}
	isInitializedReturns struct {
		result1 bool
	}
	isInitializedReturnsOnCall map[int]struct {
		result1 bool
	}
	MustBeginStub        func() *sqlx.Tx
	mustBeginMutex       sync.RWMutex
	mustBeginArgsForCall []struct {
	}
	mustBeginReturns struct {
		result1 *sqlx.Tx
	}
	mustBeginReturnsOnCall map[int]struct {
		result1 *sqlx.Tx
	}
	NamedExecStub        func(string, string, interface{}) (sql.Result, error)
	namedExecMutex       sync.RWMutex
	namedExecArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}
	namedExecReturns struct {
		result1 sql.Result
		result2 error
	}
	namedExecReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	PingContextStub        func(context.Context) error
	pingContextMutex       sync.RWMutex
	pingContextArgsForCall []struct {
		arg1 context.Context
	}
	pingContextReturns struct {
		result1 error
	}
	pingContextReturnsOnCall map[int]struct {
		result1 error
	}
	QueryxStub        func(string, string, ...interface{}) (*sqlx.Rows, error)
	queryxMutex       sync.RWMutex
	queryxArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}
	queryxReturns struct {
		result1 *sqlx.Rows
		result2 error
	}
	queryxReturnsOnCall map[int]struct {
		result1 *sqlx.Rows
		result2 error
	}
	RebindStub        func(string) string
	rebindMutex       sync.RWMutex
	rebindArgsForCall []struct {
		arg1 string
	}
	rebindReturns struct {
		result1 string
	}
	rebindReturnsOnCall map[int]struct {
		result1 string
	}
	SelectStub        func(string, interface{}, string, ...interface{}) error
	selectMutex       sync.RWMutex
	selectArgsForCall []struct {
		arg1 string
		arg2 interface{}
		arg3 string
		arg4 []interface{}
	}
	selectReturns struct {
		result1 error
	}
	selectReturnsOnCall map[int]struct {
		result1 error
	}
	SetDBInitializedStub        func(bool)
	setDBInitializedMutex       sync.RWMutex
	setDBInitializedArgsForCall []struct {
		arg1 bool
	}
	SetMaxOpenConnsStub        func(int)
	setMaxOpenConnsMutex       sync.RWMutex
	setMaxOpenConnsArgsForCall []struct {
		arg1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FabricCADB) BeginTx() db.FabricCATx {
	fake.beginTxMutex.Lock()
	ret, specificReturn := fake.beginTxReturnsOnCall[len(fake.beginTxArgsForCall)]
	fake.beginTxArgsForCall = append(fake.beginTxArgsForCall, struct {
	}{})
	fake.recordInvocation("BeginTx", []interface{}{})
	fake.beginTxMutex.Unlock()
	if fake.BeginTxStub != nil {
		return fake.BeginTxStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.beginTxReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) BeginTxCallCount() int {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	return len(fake.beginTxArgsForCall)
}

func (fake *FabricCADB) BeginTxCalls(stub func() db.FabricCATx) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = stub
}

func (fake *FabricCADB) BeginTxReturns(result1 db.FabricCATx) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	fake.beginTxReturns = struct {
		result1 db.FabricCATx
	}{result1}
}

func (fake *FabricCADB) BeginTxReturnsOnCall(i int, result1 db.FabricCATx) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	if fake.beginTxReturnsOnCall == nil {
		fake.beginTxReturnsOnCall = make(map[int]struct {
			result1 db.FabricCATx
		})
	}
	fake.beginTxReturnsOnCall[i] = struct {
		result1 db.FabricCATx
	}{result1}
}

func (fake *FabricCADB) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FabricCADB) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FabricCADB) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) DriverName() string {
	fake.driverNameMutex.Lock()
	ret, specificReturn := fake.driverNameReturnsOnCall[len(fake.driverNameArgsForCall)]
	fake.driverNameArgsForCall = append(fake.driverNameArgsForCall, struct {
	}{})
	fake.recordInvocation("DriverName", []interface{}{})
	fake.driverNameMutex.Unlock()
	if fake.DriverNameStub != nil {
		return fake.DriverNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.driverNameReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) DriverNameCallCount() int {
	fake.driverNameMutex.RLock()
	defer fake.driverNameMutex.RUnlock()
	return len(fake.driverNameArgsForCall)
}

func (fake *FabricCADB) DriverNameCalls(stub func() string) {
	fake.driverNameMutex.Lock()
	defer fake.driverNameMutex.Unlock()
	fake.DriverNameStub = stub
}

func (fake *FabricCADB) DriverNameReturns(result1 string) {
	fake.driverNameMutex.Lock()
	defer fake.driverNameMutex.Unlock()
	fake.DriverNameStub = nil
	fake.driverNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FabricCADB) DriverNameReturnsOnCall(i int, result1 string) {
	fake.driverNameMutex.Lock()
	defer fake.driverNameMutex.Unlock()
	fake.DriverNameStub = nil
	if fake.driverNameReturnsOnCall == nil {
		fake.driverNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.driverNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FabricCADB) Exec(arg1 string, arg2 string, arg3 ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2, arg3})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FabricCADB) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FabricCADB) ExecCalls(stub func(string, string, ...interface{}) (sql.Result, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FabricCADB) ExecArgsForCall(i int) (string, string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FabricCADB) ExecReturns(result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FabricCADB) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FabricCADB) Get(arg1 string, arg2 interface{}, arg3 string, arg4 ...interface{}) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 interface{}
		arg3 string
		arg4 []interface{}
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FabricCADB) GetCalls(stub func(string, interface{}, string, ...interface{}) error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FabricCADB) GetArgsForCall(i int) (string, interface{}, string, []interface{}) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FabricCADB) GetReturns(result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) GetReturnsOnCall(i int, result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) IsInitialized() bool {
	fake.isInitializedMutex.Lock()
	ret, specificReturn := fake.isInitializedReturnsOnCall[len(fake.isInitializedArgsForCall)]
	fake.isInitializedArgsForCall = append(fake.isInitializedArgsForCall, struct {
	}{})
	fake.recordInvocation("IsInitialized", []interface{}{})
	fake.isInitializedMutex.Unlock()
	if fake.IsInitializedStub != nil {
		return fake.IsInitializedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isInitializedReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) IsInitializedCallCount() int {
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	return len(fake.isInitializedArgsForCall)
}

func (fake *FabricCADB) IsInitializedCalls(stub func() bool) {
	fake.isInitializedMutex.Lock()
	defer fake.isInitializedMutex.Unlock()
	fake.IsInitializedStub = stub
}

func (fake *FabricCADB) IsInitializedReturns(result1 bool) {
	fake.isInitializedMutex.Lock()
	defer fake.isInitializedMutex.Unlock()
	fake.IsInitializedStub = nil
	fake.isInitializedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FabricCADB) IsInitializedReturnsOnCall(i int, result1 bool) {
	fake.isInitializedMutex.Lock()
	defer fake.isInitializedMutex.Unlock()
	fake.IsInitializedStub = nil
	if fake.isInitializedReturnsOnCall == nil {
		fake.isInitializedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isInitializedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FabricCADB) MustBegin() *sqlx.Tx {
	fake.mustBeginMutex.Lock()
	ret, specificReturn := fake.mustBeginReturnsOnCall[len(fake.mustBeginArgsForCall)]
	fake.mustBeginArgsForCall = append(fake.mustBeginArgsForCall, struct {
	}{})
	fake.recordInvocation("MustBegin", []interface{}{})
	fake.mustBeginMutex.Unlock()
	if fake.MustBeginStub != nil {
		return fake.MustBeginStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mustBeginReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) MustBeginCallCount() int {
	fake.mustBeginMutex.RLock()
	defer fake.mustBeginMutex.RUnlock()
	return len(fake.mustBeginArgsForCall)
}

func (fake *FabricCADB) MustBeginCalls(stub func() *sqlx.Tx) {
	fake.mustBeginMutex.Lock()
	defer fake.mustBeginMutex.Unlock()
	fake.MustBeginStub = stub
}

func (fake *FabricCADB) MustBeginReturns(result1 *sqlx.Tx) {
	fake.mustBeginMutex.Lock()
	defer fake.mustBeginMutex.Unlock()
	fake.MustBeginStub = nil
	fake.mustBeginReturns = struct {
		result1 *sqlx.Tx
	}{result1}
}

func (fake *FabricCADB) MustBeginReturnsOnCall(i int, result1 *sqlx.Tx) {
	fake.mustBeginMutex.Lock()
	defer fake.mustBeginMutex.Unlock()
	fake.MustBeginStub = nil
	if fake.mustBeginReturnsOnCall == nil {
		fake.mustBeginReturnsOnCall = make(map[int]struct {
			result1 *sqlx.Tx
		})
	}
	fake.mustBeginReturnsOnCall[i] = struct {
		result1 *sqlx.Tx
	}{result1}
}

func (fake *FabricCADB) NamedExec(arg1 string, arg2 string, arg3 interface{}) (sql.Result, error) {
	fake.namedExecMutex.Lock()
	ret, specificReturn := fake.namedExecReturnsOnCall[len(fake.namedExecArgsForCall)]
	fake.namedExecArgsForCall = append(fake.namedExecArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("NamedExec", []interface{}{arg1, arg2, arg3})
	fake.namedExecMutex.Unlock()
	if fake.NamedExecStub != nil {
		return fake.NamedExecStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.namedExecReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FabricCADB) NamedExecCallCount() int {
	fake.namedExecMutex.RLock()
	defer fake.namedExecMutex.RUnlock()
	return len(fake.namedExecArgsForCall)
}

func (fake *FabricCADB) NamedExecCalls(stub func(string, string, interface{}) (sql.Result, error)) {
	fake.namedExecMutex.Lock()
	defer fake.namedExecMutex.Unlock()
	fake.NamedExecStub = stub
}

func (fake *FabricCADB) NamedExecArgsForCall(i int) (string, string, interface{}) {
	fake.namedExecMutex.RLock()
	defer fake.namedExecMutex.RUnlock()
	argsForCall := fake.namedExecArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FabricCADB) NamedExecReturns(result1 sql.Result, result2 error) {
	fake.namedExecMutex.Lock()
	defer fake.namedExecMutex.Unlock()
	fake.NamedExecStub = nil
	fake.namedExecReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FabricCADB) NamedExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.namedExecMutex.Lock()
	defer fake.namedExecMutex.Unlock()
	fake.NamedExecStub = nil
	if fake.namedExecReturnsOnCall == nil {
		fake.namedExecReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.namedExecReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FabricCADB) PingContext(arg1 context.Context) error {
	fake.pingContextMutex.Lock()
	ret, specificReturn := fake.pingContextReturnsOnCall[len(fake.pingContextArgsForCall)]
	fake.pingContextArgsForCall = append(fake.pingContextArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("PingContext", []interface{}{arg1})
	fake.pingContextMutex.Unlock()
	if fake.PingContextStub != nil {
		return fake.PingContextStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pingContextReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) PingContextCallCount() int {
	fake.pingContextMutex.RLock()
	defer fake.pingContextMutex.RUnlock()
	return len(fake.pingContextArgsForCall)
}

func (fake *FabricCADB) PingContextCalls(stub func(context.Context) error) {
	fake.pingContextMutex.Lock()
	defer fake.pingContextMutex.Unlock()
	fake.PingContextStub = stub
}

func (fake *FabricCADB) PingContextArgsForCall(i int) context.Context {
	fake.pingContextMutex.RLock()
	defer fake.pingContextMutex.RUnlock()
	argsForCall := fake.pingContextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FabricCADB) PingContextReturns(result1 error) {
	fake.pingContextMutex.Lock()
	defer fake.pingContextMutex.Unlock()
	fake.PingContextStub = nil
	fake.pingContextReturns = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) PingContextReturnsOnCall(i int, result1 error) {
	fake.pingContextMutex.Lock()
	defer fake.pingContextMutex.Unlock()
	fake.PingContextStub = nil
	if fake.pingContextReturnsOnCall == nil {
		fake.pingContextReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingContextReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) Queryx(arg1 string, arg2 string, arg3 ...interface{}) (*sqlx.Rows, error) {
	fake.queryxMutex.Lock()
	ret, specificReturn := fake.queryxReturnsOnCall[len(fake.queryxArgsForCall)]
	fake.queryxArgsForCall = append(fake.queryxArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("Queryx", []interface{}{arg1, arg2, arg3})
	fake.queryxMutex.Unlock()
	if fake.QueryxStub != nil {
		return fake.QueryxStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryxReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FabricCADB) QueryxCallCount() int {
	fake.queryxMutex.RLock()
	defer fake.queryxMutex.RUnlock()
	return len(fake.queryxArgsForCall)
}

func (fake *FabricCADB) QueryxCalls(stub func(string, string, ...interface{}) (*sqlx.Rows, error)) {
	fake.queryxMutex.Lock()
	defer fake.queryxMutex.Unlock()
	fake.QueryxStub = stub
}

func (fake *FabricCADB) QueryxArgsForCall(i int) (string, string, []interface{}) {
	fake.queryxMutex.RLock()
	defer fake.queryxMutex.RUnlock()
	argsForCall := fake.queryxArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FabricCADB) QueryxReturns(result1 *sqlx.Rows, result2 error) {
	fake.queryxMutex.Lock()
	defer fake.queryxMutex.Unlock()
	fake.QueryxStub = nil
	fake.queryxReturns = struct {
		result1 *sqlx.Rows
		result2 error
	}{result1, result2}
}

func (fake *FabricCADB) QueryxReturnsOnCall(i int, result1 *sqlx.Rows, result2 error) {
	fake.queryxMutex.Lock()
	defer fake.queryxMutex.Unlock()
	fake.QueryxStub = nil
	if fake.queryxReturnsOnCall == nil {
		fake.queryxReturnsOnCall = make(map[int]struct {
			result1 *sqlx.Rows
			result2 error
		})
	}
	fake.queryxReturnsOnCall[i] = struct {
		result1 *sqlx.Rows
		result2 error
	}{result1, result2}
}

func (fake *FabricCADB) Rebind(arg1 string) string {
	fake.rebindMutex.Lock()
	ret, specificReturn := fake.rebindReturnsOnCall[len(fake.rebindArgsForCall)]
	fake.rebindArgsForCall = append(fake.rebindArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Rebind", []interface{}{arg1})
	fake.rebindMutex.Unlock()
	if fake.RebindStub != nil {
		return fake.RebindStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.rebindReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) RebindCallCount() int {
	fake.rebindMutex.RLock()
	defer fake.rebindMutex.RUnlock()
	return len(fake.rebindArgsForCall)
}

func (fake *FabricCADB) RebindCalls(stub func(string) string) {
	fake.rebindMutex.Lock()
	defer fake.rebindMutex.Unlock()
	fake.RebindStub = stub
}

func (fake *FabricCADB) RebindArgsForCall(i int) string {
	fake.rebindMutex.RLock()
	defer fake.rebindMutex.RUnlock()
	argsForCall := fake.rebindArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FabricCADB) RebindReturns(result1 string) {
	fake.rebindMutex.Lock()
	defer fake.rebindMutex.Unlock()
	fake.RebindStub = nil
	fake.rebindReturns = struct {
		result1 string
	}{result1}
}

func (fake *FabricCADB) RebindReturnsOnCall(i int, result1 string) {
	fake.rebindMutex.Lock()
	defer fake.rebindMutex.Unlock()
	fake.RebindStub = nil
	if fake.rebindReturnsOnCall == nil {
		fake.rebindReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.rebindReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FabricCADB) Select(arg1 string, arg2 interface{}, arg3 string, arg4 ...interface{}) error {
	fake.selectMutex.Lock()
	ret, specificReturn := fake.selectReturnsOnCall[len(fake.selectArgsForCall)]
	fake.selectArgsForCall = append(fake.selectArgsForCall, struct {
		arg1 string
		arg2 interface{}
		arg3 string
		arg4 []interface{}
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Select", []interface{}{arg1, arg2, arg3, arg4})
	fake.selectMutex.Unlock()
	if fake.SelectStub != nil {
		return fake.SelectStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.selectReturns
	return fakeReturns.result1
}

func (fake *FabricCADB) SelectCallCount() int {
	fake.selectMutex.RLock()
	defer fake.selectMutex.RUnlock()
	return len(fake.selectArgsForCall)
}

func (fake *FabricCADB) SelectCalls(stub func(string, interface{}, string, ...interface{}) error) {
	fake.selectMutex.Lock()
	defer fake.selectMutex.Unlock()
	fake.SelectStub = stub
}

func (fake *FabricCADB) SelectArgsForCall(i int) (string, interface{}, string, []interface{}) {
	fake.selectMutex.RLock()
	defer fake.selectMutex.RUnlock()
	argsForCall := fake.selectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FabricCADB) SelectReturns(result1 error) {
	fake.selectMutex.Lock()
	defer fake.selectMutex.Unlock()
	fake.SelectStub = nil
	fake.selectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) SelectReturnsOnCall(i int, result1 error) {
	fake.selectMutex.Lock()
	defer fake.selectMutex.Unlock()
	fake.SelectStub = nil
	if fake.selectReturnsOnCall == nil {
		fake.selectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.selectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FabricCADB) SetDBInitialized(arg1 bool) {
	fake.setDBInitializedMutex.Lock()
	fake.setDBInitializedArgsForCall = append(fake.setDBInitializedArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetDBInitialized", []interface{}{arg1})
	fake.setDBInitializedMutex.Unlock()
	if fake.SetDBInitializedStub != nil {
		fake.SetDBInitializedStub(arg1)
	}
}

func (fake *FabricCADB) SetDBInitializedCallCount() int {
	fake.setDBInitializedMutex.RLock()
	defer fake.setDBInitializedMutex.RUnlock()
	return len(fake.setDBInitializedArgsForCall)
}

func (fake *FabricCADB) SetDBInitializedCalls(stub func(bool)) {
	fake.setDBInitializedMutex.Lock()
	defer fake.setDBInitializedMutex.Unlock()
	fake.SetDBInitializedStub = stub
}

func (fake *FabricCADB) SetDBInitializedArgsForCall(i int) bool {
	fake.setDBInitializedMutex.RLock()
	defer fake.setDBInitializedMutex.RUnlock()
	argsForCall := fake.setDBInitializedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FabricCADB) SetMaxOpenConns(arg1 int) {
	fake.setMaxOpenConnsMutex.Lock()
	fake.setMaxOpenConnsArgsForCall = append(fake.setMaxOpenConnsArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetMaxOpenConns", []interface{}{arg1})
	fake.setMaxOpenConnsMutex.Unlock()
	if fake.SetMaxOpenConnsStub != nil {
		fake.SetMaxOpenConnsStub(arg1)
	}
}

func (fake *FabricCADB) SetMaxOpenConnsCallCount() int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return len(fake.setMaxOpenConnsArgsForCall)
}

func (fake *FabricCADB) SetMaxOpenConnsCalls(stub func(int)) {
	fake.setMaxOpenConnsMutex.Lock()
	defer fake.setMaxOpenConnsMutex.Unlock()
	fake.SetMaxOpenConnsStub = stub
}

func (fake *FabricCADB) SetMaxOpenConnsArgsForCall(i int) int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	argsForCall := fake.setMaxOpenConnsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FabricCADB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.driverNameMutex.RLock()
	defer fake.driverNameMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	fake.mustBeginMutex.RLock()
	defer fake.mustBeginMutex.RUnlock()
	fake.namedExecMutex.RLock()
	defer fake.namedExecMutex.RUnlock()
	fake.pingContextMutex.RLock()
	defer fake.pingContextMutex.RUnlock()
	fake.queryxMutex.RLock()
	defer fake.queryxMutex.RUnlock()
	fake.rebindMutex.RLock()
	defer fake.rebindMutex.RUnlock()
	fake.selectMutex.RLock()
	defer fake.selectMutex.RUnlock()
	fake.setDBInitializedMutex.RLock()
	defer fake.setDBInitializedMutex.RUnlock()
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FabricCADB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.FabricCADB = new(FabricCADB)
