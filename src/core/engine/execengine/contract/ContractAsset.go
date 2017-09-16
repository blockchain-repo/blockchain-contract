package contract

type ContractAsset struct {
	AssetId     string            `json:"AssetId"`
	Name        string            `json:"Name"`
	Caption     string            `json:"Caption"`
	Description string            `json:"Description"`
	Unit        string            `json:"Unit"`
	Amount      interface{}       `json:"Amount"`
	MetaData    map[string]string `json:"MetaData"`
}

func NewContractAsset() *ContractAsset {
	ca := &ContractAsset{}
	return ca
}

/*
func (ca *ContractAsset) GetItem(p_propertyname string)interface{}{
	var r_result interface{}
	//Get Value By reflect
	v_refl_object := reflect.ValueOf(ca).Elem()
	v_refl_field := v_refl_object.FieldByName(p_propertyname)
	switch v_refl_field.Kind(){
	case reflect.Map:
		r_result = v_refl_field.MapIndex(reflect.ValueOf(p_propertyname))
	default:
		r_result = v_refl_field.Interface()
	}
	return r_result
}
*/

func (ca *ContractAsset) GetAssetId() string {
	return ca.AssetId
}

func (ca *ContractAsset) GetName() string {
	return ca.Name
}
func (ca *ContractAsset) GetCaption() string {
	return ca.Caption
}
func (ca *ContractAsset) GetDescription() string {
	return ca.Description
}
func (ca *ContractAsset) GetUnit() string {
	return ca.Unit
}
func (ca *ContractAsset) GetAmount() interface{} {
	return ca.Amount
}
func (ca *ContractAsset) GetMetaData() map[string]string {
	return ca.MetaData
}

func (ca *ContractAsset) SetAssetId(p_id string) {
	ca.AssetId = p_id
}

func (ca *ContractAsset) SetName(p_name string) {
	ca.Name = p_name
}
func (ca *ContractAsset) SetCaption(p_caption string) {
	ca.Caption = p_caption
}
func (ca *ContractAsset) SetDescription(p_description string) {
	ca.Description = p_description
}
func (ca *ContractAsset) SetUnit(p_unit string) {
	ca.Unit = p_unit
}
func (ca *ContractAsset) SetAmount(p_amount interface{}) {
	ca.Amount = p_amount
}
func (ca *ContractAsset) SetMetaData(p_metaData map[string]string) {
	ca.MetaData = p_metaData
}
