package snapshot

import "io"

type SectionName string

const (
	SectionNameChainSnapshotHeader         SectionName = "amax::chain::chain_snapshot_header"
	SectionNameBlockState                  SectionName = "amax::chain::block_state"
	SectionNameAccountObject               SectionName = "amax::chain::account_object"
	SectionNameAccountMetadataObject       SectionName = "amax::chain::account_metadata_object"
	SectionNameAccountRamCorrectionObject  SectionName = "amax::chain::account_ram_correction_object"
	SectionNameGlobalPropertyObject        SectionName = "amax::chain::global_property_object"
	SectionNameProtocolStateObject         SectionName = "amax::chain::protocol_state_object"
	SectionNameDynamicGlobalPropertyObject SectionName = "amax::chain::dynamic_global_property_object"
	SectionNameBlockSummaryObject          SectionName = "amax::chain::block_summary_object"
	SectionNameTransactionObject           SectionName = "amax::chain::transaction_object"
	SectionNameGeneratedTransactionObject  SectionName = "amax::chain::generated_transaction_object"
	SectionNameCodeObject                  SectionName = "amax::chain::code_object"
	SectionNameContractTables              SectionName = "contract_tables"
	SectionNamePermissionObject            SectionName = "amax::chain::permission_object"
	SectionNamePermissionLinkObject        SectionName = "amax::chain::permission_link_object"
	SectionNameResourceLimitsObject        SectionName = "amax::chain::resource_limits::resource_limits_object"
	SectionNameResourceUsageObject         SectionName = "amax::chain::resource_limits::resource_usage_object"
	SectionNameResourceLimitsStateObject   SectionName = "amax::chain::resource_limits::resource_limits_state_object"
	SectionNameResourceLimitsConfigObject  SectionName = "amax::chain::resource_limits::resource_limits_config_object"
	SectionNameGenesisState                SectionName = "amax::chain::genesis_state"

	// Ultra Specific
	SectionAccountFreeActionsObject SectionName = "amax::chain::account_free_actions_object"
)

type Section struct {
	Name       SectionName
	Offset     uint64
	Size       uint64 // This includes the section name and row count
	BufferSize uint64 // This represents the bytes that are following the section header
	RowCount   uint64 // This is a count of rows packed in `Buffer`
	Buffer     io.Reader
}

type sectionHandlerFunc func(s *Section, f sectionCallbackFunc) error
type sectionCallbackFunc func(obj interface{}) error
