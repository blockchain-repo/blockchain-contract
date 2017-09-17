package rethinkdb

import (
	"encoding/json"
	"errors"
	"fmt"
	r "gopkg.in/gorethink/gorethink.v3"
	"strconv"
	"time"
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
)

func Get(db string, name string, id string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Run(session)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	return res
}

func Insert(db string, name string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Insert(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	return res
}

func Update(db string, name string, id string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Update(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	return res
}

func Delete(db string, name string, id string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Delete().RunWrite(session)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	return res
}

/*----------------------------unicontract ops-------------------------------------*/

/*----------------------------- contracts start---------------------------------------*/
// contract serialize contract string
func InsertContract(contract string) bool {
	if contract == "" {
		return false
	}
	res := Insert(DBNAME, TABLE_CONTRACTS, contract)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func DeleteContract(id string) bool {
	if len(id) == 0 {
		return false
	}

	res := Delete(DBNAME, TABLE_CONTRACTS, id)
	if res.Deleted >= 1 {
		return true
	}
	return false
}

func GetContractById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACTS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取具有相同contractId的合约
func GetContractsByContractId(contractId string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Filter(r.Row.Field("ContractBody").Field("ContractId").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}

	return common.Serialize(blo), nil
}

//根据传入条件查询 contract content 合约 仅取出一条 , ContractState = Contract_Create or Contract_Signature
func GetContractContentByCondition(contractProductId string, owner string) (string, error) {
	if contractProductId == "" {
		return "", errors.New("contractProductId blank")
	}
	// company owner
	if owner == "" {
		return "", errors.New("owner blank")
	}
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
		Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq("Contract_Signature").
			Or(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq("Contract_Create"))).
		Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
		Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
		Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
		Max(r.Row.Field("transaction").Field("timestamp")).
		Ungroup().Field("reduction").
		OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
		Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}
	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

//根据传入条件查询 publish contract 合约 仅取出一条 , ContractState = Contract_Create
func GetPublishContractByCondition(contractProductId string, owner string, contractState string) (string, error) {
	if contractProductId == "" {
		return "", errors.New("contractProductId blank")
	}
	// company owner
	if owner == "" {
		return "", errors.New("owner blank")
	}
	contractState = "Contract_Create"
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
		Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
		Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
		Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
		Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
		Max(r.Row.Field("transaction").Field("timestamp")).
		Ungroup().Field("reduction").
		OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
		Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}
	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

//根据传入条件查询 execute 合约 仅取出一条 , ContractState = Contract_Signature
func GetOneContractByCondition(contractProductId string, owner string, contractState string) (string, error) {
	if contractProductId == "" {
		return "", errors.New("contractProductId blank")
	}

	if owner == "" {
		return "", errors.New("owner blank")
	}

	if contractState == "" {
		contractState = "Contract_Signature"
	}
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	if owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	}
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}
	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	uniledgerlog.Warn("query len is %+v", blo)
	return common.Serialize(blo), nil
}

func GetContractsPaginationByCondition(contractProductId string, owner string, contractState string, pageNumStart int32, pageNumEnd int32) (totalRecords int32, result string, err error) {
	//if owner == "" {
	//	return "", errors.New("owner blank")
	//}
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	if contractProductId == "" && contractState == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else if contractProductId == "" && contractState == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else if contractProductId == "" && contractState != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else if contractProductId != "" && contractState == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractProductId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else if contractProductId == "" && contractState != "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else if contractProductId != "" && contractState == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else if contractProductId != "" && contractState != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			Count().
			Run(session)
	}

	if err != nil {
		return 0, "", err
	}
	if res.IsNil() {
		return 0, "", nil
	}
	// totalRecords
	err = res.One(&totalRecords)
	if err != nil {
		return 0, "", err
	}

	if contractProductId == "" && contractState == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).
			Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else if contractProductId == "" && contractState == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else if contractProductId == "" && contractState != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else if contractProductId != "" && contractState == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else if contractProductId == "" && contractState != "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else if contractProductId != "" && contractState == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else if contractProductId != "" && contractState != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	}

	if err != nil {
		return 0, "", err
	}
	if res.IsNil() {
		return 0, "", nil
	}
	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return 0, "", err
	}
	return totalRecords, common.Serialize(blo), nil
}

// 根据传入条件查询 execute 合约, ContractState = Contract_Signature
func GetContractsByCondition(contractId string, owner string, contractState string) (string, error) {
	if owner == "" {
		return "", errors.New("owner blank")
	}
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error

	if contractId == "" && contractState == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId == "" && contractState == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId == "" && contractState != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId != "" && contractState == "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId == "" && contractState != "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId != "" && contractState == "" && owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else if contractId != "" && contractState != "" && owner == "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Group(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId")).
			Max(r.Row.Field("transaction").Field("timestamp")).
			Ungroup().Field("reduction").
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).Field("transaction").Field("Contract").
			Run(session)
	}

	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}
	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据传入条件查询合约
func GetContractsLogByCondition(contractId string, owner string, contractState string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}
	contractState = "Contract_In_Process"
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	if owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).
			Run(session)

	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractId").Eq(contractId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).
			Run(session)
	}
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}
	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// GetContractsLogPaginationByCondition 分页查询执行日志，暂时作为demo
func GetContractsLogPaginationByCondition(contractProductId string, owner string, contractState string, pageNumStart int32, pageNumEnd int32) (totalRecords int32, result string, err error) {
	if contractProductId == "" {
		return 0, "", errors.New("contractProductId blank")
	}
	contractState = "Contract_In_Process"
	session := ConnectDB(DBNAME)
	var res *r.Cursor

	if owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Count().Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Count().Run(session)
	}
	if err != nil {
		return 0, "", err
	}
	if err != nil {
		return 0, "", err
	}
	err = res.One(&totalRecords)
	if err != nil {
		return 0, "", err
	}

	if owner != "" {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractOwners").Contains(owner)).
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	} else {
		res, err = r.Table(TABLE_CONTRACT_OUTPUTS).
			//Filter(r.Row.Field("ContractBody").Field("ContractOwners").Contains(owner)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractProductId").Eq(contractProductId)).
			Filter(r.Row.Field("transaction").Field("Contract").Field("ContractBody").Field("ContractState").Eq(contractState)).
			OrderBy(r.Asc(r.Row.Field("transaction").Field("timestamp"))).
			Slice(pageNumStart, pageNumEnd).
			Run(session)
	}
	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return 0, "", err
	}

	return totalRecords, common.Serialize(blo), nil
}

//根据 contract.id 获取合约处理主节点
func GetContractMainPubkeyByContract(id string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Get(id).Count().Default(0).Run(session)
	if err != nil {
		return "", err
	}

	var blo int
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	if blo == 0 {
		return "", nil
	}

	// continue ...
	res, err = r.Table(TABLE_CONTRACTS).Get(id).Field("ContractHead").Field("MainPubkey").Run(session)
	if err != nil {
		return "", err
	}

	var blo2 string
	err = res.One(&blo2)
	if err != nil {
		return "", err
	}
	return blo2, nil
}

func GetNoConsensusContracts(time string, state int) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).
		Filter(r.Row.Field("ContractHead").Field("ConsensusResult").Eq(state)).
		Filter(r.Row.Field("ContractHead").Field("AssignTime").Le(time)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func SetContractConsensusResultById(id string, time string, ConsensusResult int) error {
	strJSON := fmt.Sprintf("{\"ContractHead\":{\"AssignTime\":\"%s\",\"ConsensusResult\":%d}}",
		time, ConsensusResult)

	res := Update(DBNAME, TABLE_CONTRACTS, id, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

// 获取合约个数
func GetContractsCount() (string, error) {

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Count().Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

// 获取合约不同状态个数
func GetContractStatsCount(stat string) (string, error) {

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).
		Filter(r.Row.Field("ContractBody").Field("ContractState").Eq(stat)).
		Count().Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return blo, nil
}

//demo使用---------------------------------------------------------------------------------------------------------------
func QueryOutput(contractID string) (string, error) {
	if len(contractID) == 0 {
		return "", fmt.Errorf("contractID is null")
	}
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).
		Filter(r.Row.Field("transaction").Field("Relation").Field("ContractId").Eq(contractID)).
		Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func QueryOutputNum(contractID string) (int, error) {
	var count int
	if len(contractID) == 0 {
		return count, fmt.Errorf("contractID is null")
	}
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).
		Filter(r.Row.Field("transaction").Field("Relation").Field("ContractId").Eq(contractID)).
		Count().Run(session)
	if err != nil {
		return count, err
	}
	if res.IsNil() {
		return count, nil
	}

	var blo string
	err = res.One(&blo)
	if err != nil {
		return count, err
	}

	count, err = strconv.Atoi(blo)
	if err != nil {
		return count, err
	}

	return count, nil
}

func QueryContractStartTime(contractID string) (string, error) {
	if len(contractID) == 0 {
		return "", fmt.Errorf("contractID is null")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).
		Filter(r.Row.Field("ContractBody").Field("ContractId").Eq(contractID)).
		Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
	}

	body, ok := blo["ContractBody"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("error")
	}

	startTime, ok := body["StartTime"].(string)
	if !ok {
		return "", fmt.Errorf("error")
	}
	return startTime, nil
}

//demo使用---------------------------------------------------------------------------------------------------------------

/*----------------------------- contracts end---------------------------------------*/

/*----------------------------- votes start---------------------------------------*/
// vote serialize vote string
func InsertVote(vote string) bool {
	if vote == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_VOTES, vote)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetVoteById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}
	res := Get(DBNAME, TABLE_VOTES, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取所有 vote
func GetVotesByContractId(contractId string) (string, error) {

	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).Filter(r.Row.Field("Vote").Field("VoteFor").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 删除一系列id的vote
func DeleteVotes(slID []interface{}) (int, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).
		GetAll(slID...).Delete().RunWrite(session)
	return res.Deleted, err
}

/*----------------------------- votes end---------------------------------------*/

/*----------------------------- contractOutputs start---------------------------------------*/
// contractOutput serialize contractOutput string
func InsertContractOutput(contractOutput string) bool {
	if contractOutput == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_OUTPUTS, contractOutput)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func UpdateContractOutVote(id string, vote map[string]interface{}, index int) bool {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).Get(id).Update(map[string]interface{}{"transaction": map[string]interface{}{"Relation": map[string]interface{}{"Votes": r.Row.Field("transaction").Field("Relation").Field("Votes").ChangeAt(index, vote)}}}).RunWrite(session)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	if res.Replaced >= 1 {
		return true
	}
	return false
}

func GetContractOutputById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACT_OUTPUTS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约 整体的[id] transaction.contract.id 或者 relation.contractId 获取所有 contractOutput
func GetContractOutputByContractPrimaryId(id string) (string, error) {

	if id == "" {
		return "", errors.New("id blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).Filter(r.Row.Field("transaction").Field("Contract").Field("id").Eq(id)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- contractOutputs end---------------------------------------*/

/*----------------------------- contractTask start---------------------------------------*/
// contractTask serialize contractTask string
func InsertContractTask(contractTask string) bool {
	if contractTask == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_TASKS, contractTask)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetContractTaskById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACT_TASKS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetContractTasksByContractId(contractId string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_TASKS).Filter(r.Row.Field("ContractId").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- contractTask end---------------------------------------*/

/*----------------------------- consensusFailures start---------------------------------------*/
// consensusFailures serialize consensusFailures string
func InsertConsensusFailure(consensusFailures string) bool {
	if consensusFailures == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONSENSUS_FAILURES, consensusFailures)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetConsensusFailureById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONSENSUS_FAILURES, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetConsensusFailuresByConsensusId(consensusId string) (string, error) {
	if consensusId == "" {
		return "", errors.New("consensusId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONSENSUS_FAILURES).Filter(r.Row.Field("ConsensusId").Eq(consensusId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetConsensusFailuresCount() (int, error) {
	session := ConnectDB(DBNAME)
	count, err := r.Table(TABLE_CONSENSUS_FAILURES).Count().Run(session)
	if err != nil {
		return -1, err
	}
	if count.IsNil() {
		return -1, nil
	}

	var blo int
	err = count.One(&blo)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return -1, err
	}
	return blo, nil
}

/*----------------------------- consensusFailures end---------------------------------------*/

/*----------------------------- SendFailingRecords start---------------------------------------*/
func GetAllRecords(db string, name string) ([]string, error) {
	session := ConnectDB(db)
	ids, err := r.Table(name).Field("id").Run(session)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	var idlist []string
	err = ids.All(&idlist)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return nil, errors.New(err.Error())
	}
	return idlist, nil
}

func GetSendFailingRecordsCount() (int, error) {
	session := ConnectDB(DBNAME)
	count, err := r.Table(TABLE_SEND_FAILING_RECORDS).Count().Run(session)
	if err != nil {
		return -1, err
	}
	if count.IsNil() {
		return -1, nil
	}

	var blo int
	err = count.One(&blo)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return -1, err
	}
	return blo, nil
}

/*----------------------------- SendFailingRecords end---------------------------------------*/

/*智能微网demo start---------------------------------------------------------*/
func _Insert(strDBName, strTableName, strJson string) error {
	if len(strJson) == 0 {
		return fmt.Errorf("param is null")
	}
	session := ConnectDB(strDBName)
	res, err := r.Table(strTableName).Insert(r.JSON(strJson)).RunWrite(session)
	if err != nil {
		return err
	}
	if res.Inserted >= 1 {
		return nil
	}
	return fmt.Errorf("insert %s error", strTableName)
}

//---------------------------------------------------------------------------
// 插入 EnergyTradingDemoRole 表
func InsertEnergyTradingDemoRole(strJson string) error {
	return _Insert(DBNAME, TABLE_ENERGYTRADINGDEMO_ROLE, strJson)
}

//---------------------------------------------------------------------------
// 插入 EnergyTradingDemoEnergy 表
func InsertEnergyTradingDemoEnergy(strJson string) error {
	return _Insert(DBNAME, TABLE_ENERGYTRADINGDEMO_ENERGY, strJson)
}

//---------------------------------------------------------------------------
// 插入 EnergyTradingDemoTransaction 表
func InsertEnergyTradingDemoTransaction(strJson string) error {
	return _Insert(DBNAME, TABLE_ENERGYTRADINGDEMO_TRANSACTION, strJson)
}

//---------------------------------------------------------------------------
// 插入 EnergyTradingDemoBill 表
func InsertEnergyTradingDemoBill(strJson string) error {
	return _Insert(DBNAME, TABLE_ENERGYTRADINGDEMO_BILL, strJson)
}

//---------------------------------------------------------------------------
// 插入 EnergyTradingDemoMsgNotice 表
func InsertEnergyTradingDemoMsgNotice(strJson string) error {
	return _Insert(DBNAME, TABLE_ENERGYTRADINGDEMO_MSGNOTICE, strJson)
}

//---------------------------------------------------------------------------
// 插入 EnergyTradingDemoPrice 表
func InsertEnergyTradingDemoPrice(strJson string) error {
	return _Insert(DBNAME, TABLE_ENERGYTRADINGDEMO_PRICE, strJson)
}

//---------------------------------------------------------------------------
// 获取电表余额
func GetMoneyFromEnergy(strPublicKey string) (float64, error) {
	if len(strPublicKey) == 0 {
		return 0, fmt.Errorf("param is null")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_ENERGY).
		Filter(r.Row.Field("PublicKey").Eq(strPublicKey)).
		OrderBy(r.Desc("Timestamp")).
		Run(session)
	if err != nil {
		return 0, err
	}

	if res.IsNil() {
		return 0, nil
	}

	var items []map[string]interface{}
	err = res.All(&items)
	if err != nil {
		return 0, err
	}

	if len(items) == 0 {
		return 0, nil
	}

	money, ok := items[0]["Money"].(float64)
	if !ok {
		return 0, fmt.Errorf("items[0][\"Money\"].(float64) is error")
	}

	return money, nil
}

//---------------------------------------------------------------------------
// 获取用户账户余额
func GetUserMoneyFromTransaction(strPublicKey string) (float64, error) {
	if len(strPublicKey) == 0 {
		return 0, fmt.Errorf("param is null")
	}

	// 先查充值的金额
	in, err := _GetMoney("ToPublicKey", strPublicKey)
	if err != nil {
		return 0, err
	}

	// 再查购电的金额
	out, err := _GetMoney("FromPublicKey", strPublicKey)
	if err != nil {
		return 0, err
	}

	return (in - out), nil
}

func _GetMoney(strField, strPublicKey string) (float64, error) {
	var money float64
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_TRANSACTION).
		Filter(r.Row.Field(strField).Eq(strPublicKey)).
		Run(session)
	if err != nil {
		return 0, err
	}

	if res.IsNil() {
		return 0, nil
	}

	var items []map[string]interface{}
	err = res.All(&items)
	if err != nil {
		return 0, err
	}

	if len(items) == 0 {
		return 0, nil
	}

	for _, v := range items {
		value, ok := v["Money"].(float64)
		if !ok {
			return money, fmt.Errorf("assert error")
		}
		money += value
	}

	return money, nil
}

//---------------------------------------------------------------------------
// 通过用户key查询电表key
func GetMeterKeyByUserKey(strUserKey string) (string, error) {
	if len(strUserKey) == 0 {
		return "", fmt.Errorf("param is null")
	}

	information := fmt.Sprintf("{\"ownerPublicKey\":\"%s\"}", strUserKey)
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_ROLE).
		Filter(r.Row.Field("Infermation").Eq(information)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", fmt.Errorf("this user has no meter")
	}

	var item map[string]interface{}
	err = res.One(&item)
	if err != nil {
		return "", err
	}

	meterKey, ok := item["PublicKey"].(string)
	if !ok {
		return "", fmt.Errorf("item[\"PublicKey\"].(string) is error")
	}

	return meterKey, nil
}

//---------------------------------------------------------------------------
// 获取阶梯电价
func GetPrice() ([]map[string]interface{}, error) {
	var items []map[string]interface{}
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_PRICE).
		Run(session)
	if err != nil {
		return items, err
	}

	if res.IsNil() {
		return items, fmt.Errorf("no price")
	}

	err = res.All(&items)
	if err != nil {
		return items, err
	}

	return items, nil
}

//---------------------------------------------------------------------------
// 查询电表最后查询时间点
func GetMeterQueryLastTime(strPublicKey string) (string, error) {
	if len(strPublicKey) == 0 {
		return "", fmt.Errorf("param is null")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_ROLE).
		Filter(r.Row.Field("PublicKey").Eq(strPublicKey)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", fmt.Errorf("no meter")
	}

	var item map[string]interface{}
	err = res.One(&item)
	if err != nil {
		return "", err
	}

	LastTimestamp, ok := item["LastTimestamp"].(string)
	if !ok {
		return "", fmt.Errorf("item[\"LastTimestamp\"].(string) is error")
	}

	return LastTimestamp, nil
}

//---------------------------------------------------------------------------
// 更新电表最后查询时间点
func UpdateMeterQueryLastTime(strPublicKey, strTimestamp string) error {
	if len(strPublicKey) == 0 ||
		len(strTimestamp) == 0 {
		return fmt.Errorf("param is null")
	}

	json := fmt.Sprintf("{\"LastTimestamp\":\"%s\"}", strTimestamp)
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_ROLE).
		Filter(r.Row.Field("PublicKey").Eq(strPublicKey)).
		Update(r.JSON(json)).
		RunWrite(session)
	if err != nil {
		return err
	}
	if res.Replaced|res.Unchanged >= 1 {

	} else {
		return fmt.Errorf("update failed")
	}
	return nil
}

//---------------------------------------------------------------------------
// 获得某个时间点的电表信息from EnergyTradingDemoEnergy
func GetMeterinforFromEnergy(strPublicKey, strTimestamp string, desc bool) (map[string]interface{}, error) {
	var item map[string]interface{}
	if len(strPublicKey) == 0 {
		return item, fmt.Errorf("param is null")
	}

	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	if desc {
		res, err = r.Table(TABLE_ENERGYTRADINGDEMO_ENERGY).
			Filter(r.Row.Field("PublicKey").Eq(strPublicKey)).
			Filter(r.Row.Field("Timestamp").Le(strTimestamp)).
			OrderBy(r.Desc("Timestamp")).
			Run(session)
	} else {
		res, err = r.Table(TABLE_ENERGYTRADINGDEMO_ENERGY).
			Filter(r.Row.Field("PublicKey").Eq(strPublicKey)).
			Filter(r.Row.Field("Timestamp").Ge(strTimestamp)).
			OrderBy("Timestamp").
			Run(session)
	}

	if err != nil {
		return item, err
	}

	if res.IsNil() {
		return item, fmt.Errorf("is null")
	}

	err = res.One(&item)
	if err != nil {
		return item, err
	}

	return item, nil
}

//---------------------------------------------------------------------------
// 获得电表某两个时间段的耗电量，以及此时的电表余额和当月耗电量
func GetMeterInformation(strPublicKey, startTime, endTime string) (float64, float64, float64, error) {
	if len(strPublicKey) == 0 ||
		len(startTime) == 0 ||
		len(endTime) == 0 {
		return 0, 0, 0, fmt.Errorf("param is null")
	}

	// 获得电表key
	meterKey, err := GetMeterKeyByUserKey(strPublicKey)
	if err != nil {
		return 0, 0, 0, err
	}

	// 获得上个时间点的电表信息
	meter1, err := GetMeterinforFromEnergy(meterKey, startTime, false)
	if err != nil {
		return 0, 0, 0, err
	}

	// 获得当前时间点的电表信息
	meter2, err := GetMeterinforFromEnergy(meterKey, endTime, true)
	if err != nil {
		return 0, 0, 0, err
	}

	// 获得耗电量、当前余额、当月耗电量
	electricity1, ok := meter1["Electricity"].(float64)
	if !ok {
		return 0, 0, 0, fmt.Errorf("meter1[\"Electricity\"].(float64) is error")
	}

	electricity2, ok := meter2["Electricity"].(float64)
	if !ok {
		return 0, 0, 0, fmt.Errorf("meter2[\"Electricity\"].(float64) is error")
	}

	money, ok := meter2["Money"].(float64)
	if !ok {
		return 0, 0, 0, fmt.Errorf("meter2[\"Money\"].(float64) is error")
	}

	totalElectricity, ok := meter2["TotalElectricity"].(float64)
	if !ok {
		return 0, 0, 0, fmt.Errorf("meter2[\"TotalElectricity\"].(float64) is error")
	}

	return (electricity2 - electricity1), money, totalElectricity, nil
}

//---------------------------------------------------------------------------
// 获得某时间段内各个发电厂的发电量
func GetPowerPlantEnergy(slPubkey []string, startTime, endTime string) (map[string]float64, error) {
	energys := make(map[string]float64)
	var err error

	for index, value := range slPubkey {
		energys[slPubkey[index]], err = _GetEnergy(value, startTime, endTime)
		if err != nil {
			return energys, err
		}
	}

	return energys, nil
}

func _GetEnergy(strPublickey string, startTime, endTime string) (float64, error) {
	var energy float64
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_ENERGY).
		Filter(r.Row.Field("PublicKey").Eq(strPublickey)).
		Filter(r.Row.Field("Timestamp").Ge(startTime)).
		Filter(r.Row.Field("Timestamp").Le(endTime)).
		Run(session)
	if err != nil {
		return energy, err
	}

	if res.IsNil() {
		return energy, fmt.Errorf("no power plant")
	}

	var items []map[string]interface{}
	err = res.All(&items)
	if err != nil {
		return energy, err
	}

	for _, v := range items {
		e, ok := v["Electricity"].(float64)
		if !ok {
			return energy, fmt.Errorf("[\"Electricity\"].(float64) is error")
		}
		energy += e
	}

	return energy, nil
}

//---------------------------------------------------------------------------
// 获得相应类型role的publickey
func GetRolePublicKey(type_ int) ([]string, error) {
	var slKeys []string
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_ROLE).
		Filter(r.Row.Field("Type").Eq(type_)).
		Run(session)
	if err != nil {
		return slKeys, err
	}

	if res.IsNil() {
		return slKeys, fmt.Errorf("no role")
	}

	var items []map[string]interface{}
	err = res.All(&items)
	if err != nil {
		return slKeys, err
	}

	for _, v := range items {
		e, ok := v["PublicKey"].(string)
		if !ok {
			return slKeys, fmt.Errorf("[\"PublicKey\"].(string) is error")
		}

		slKeys = append(slKeys, e)
	}
	return slKeys, nil
}

//---------------------------------------------------------------------------
// 获得转账记录
func GetTransactionRecords() (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_ENERGYTRADINGDEMO_TRANSACTION).
		Filter(r.Row.Field("Type").Eq(1)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", fmt.Errorf("no records")
	}

	var items []map[string]interface{}
	err = res.All(&items)
	if err != nil {
		return "", err
	}

	fotmat := "2006-01-02 15:04:05"
	var slmap []map[string]interface{}
	for _, value := range items {
		mapRecords := make(map[string]interface{})
		mapRecords["Id"], _ = value["id"].(string)
		mapRecords["From"] = "个人"
		mapRecords["To"] = "运营商"
		mapRecords["Money"], _ = value["Money"].(float64)
		strtimeStamp, _ := value["Timestamp"].(string)
		timeStamp_, _ := strconv.Atoi(strtimeStamp)
		tm := time.Unix(int64(timeStamp_)/1000, 0)
		//time1, _ := time.Parse(fotmat, tm.Format(fotmat))
		mapRecords["Timestamp"] = tm.Format(fotmat)
		mapRecords["BillId"], _ = value["BillId"].(string)

		slmap = append(slmap, mapRecords)
	}

	slData, _ := json.Marshal(slmap)

	return string(slData), nil
}

/*智能微网demo end---------------------------------------------------------*/

/* tianan */
func GetInfoByUser(pubkey string) (map[string]interface{}, error) {
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	//var resYield *r.Cursor
	var err error
	res, _ = r.Table(TABLE_EARNINGS).Filter(r.Row.Field("pubkey").Eq(pubkey)).Max(r.Row.Field("timestamp")).Run(session)

	//uniledgerlog.Info(res)
	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return blo, err
	}
	return blo, nil
}

func GetLastInterest(pubkey string) ([]map[string]interface{}, error) {
	session := ConnectDB(DBNAME)
	var res *r.Cursor
	var err error
	res, _ = r.Table(TABLE_EARNINGS).Filter(r.Row.Field("pubkey").Eq(pubkey)).Run(session)

	//uniledgerlog.Info(res)
	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return blo, err
	}
	return blo, nil
}

func InsertInterestCount(str string) bool {
	if str == "" {
		return false
	}
	res := Insert(DBNAME, TABLE_EARNINGS, str)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

/* tianan end*/
