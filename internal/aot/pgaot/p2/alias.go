package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	_ "unsafe"
)
//go:linkname F_brin_build_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brin_build_desc
func F_brin_build_desc(m *base.Module, l0 int32) int32
//go:linkname F_union_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_union_tuples
func F_union_tuples(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_terminate_brin_buildstate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_terminate_brin_buildstate
func F_terminate_brin_buildstate(m *base.Module, l0 int32)
//go:linkname F_add_values_to_range github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_values_to_range
func F_add_values_to_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_inclusion_get_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inclusion_get_procinfo
func F_inclusion_get_procinfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_brin_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_brin_range_deserialize
func F_brin_range_deserialize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_range_deduplicate_values github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_deduplicate_values
func F_range_deduplicate_values(m *base.Module, l0 int32)
//go:linkname F_build_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_build_expanded_ranges
func F_build_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_minmax_multi_get_procinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_minmax_multi_get_procinfo
func F_minmax_multi_get_procinfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_reduce_expanded_ranges github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_reduce_expanded_ranges
func F_reduce_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_brin_doupdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brin_doupdate
func F_brin_doupdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_brin_getinsertbuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brin_getinsertbuffer
func F_brin_getinsertbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brinRevmapInitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_brinRevmapInitialize
func F_brinRevmapInitialize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_brinRevmapTerminate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_brinRevmapTerminate
func F_brinRevmapTerminate(m *base.Module, l0 int32)
//go:linkname F_brinGetTupleForHeapBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brinGetTupleForHeapBlock
func F_brinGetTupleForHeapBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_brin_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_brin_form_tuple
func F_brin_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_brin_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_brin_deform_tuple
func F_brin_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_free_attrmap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_attrmap
func F_free_attrmap(m *base.Module, l0 int32)
//go:linkname F_build_attrmap_by_name_if_req github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_attrmap_by_name_if_req
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_toast_raw_datum_size github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_toast_raw_datum_size
func F_toast_raw_datum_size(m *base.Module, l0 int32) int32
//go:linkname F_getmissingattr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getmissingattr
func F_getmissingattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_attisnull github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_attisnull
func F_heap_attisnull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nocachegetattr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nocachegetattr
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_getsysattr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_getsysattr
func F_heap_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_copytuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copytuple
func F_heap_copytuple(m *base.Module, l0 int32) int32
//go:linkname F_heap_copy_tuple_as_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_copy_tuple_as_datum
func F_heap_copy_tuple_as_datum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_heap_form_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_form_tuple
func F_heap_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_modify_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_modify_tuple
func F_heap_modify_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_deform_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_deform_tuple
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_form_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_form_minimal_tuple
func F_heap_form_minimal_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_heap_copy_minimal_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_copy_minimal_tuple
func F_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CopyIndexTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CopyIndexTuple
func F_CopyIndexTuple(m *base.Module, l0 int32) int32
//go:linkname F_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_open
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_try_relation_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_try_relation_open
func F_try_relation_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_relation_openrv
func F_relation_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_relation_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_relation_close
func F_relation_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_local_real_reloption github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_add_local_real_reloption
func F_add_local_real_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64, l5 float64, l6 int32)
//go:linkname F_transformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformRelOptions
func F_transformRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_untransformRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_untransformRelOptions
func F_untransformRelOptions(m *base.Module, l0 int32) int32
//go:linkname F_heap_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_reloptions
func F_heap_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initialize_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initialize_reloptions
func F_initialize_reloptions(m *base.Module)
//go:linkname F_parseRelOptionsInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parseRelOptionsInternal
func F_parseRelOptionsInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fillRelOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fillRelOptions
func F_fillRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_ScanKeyEntryInitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ScanKeyEntryInitialize
func F_ScanKeyEntryInitialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ScanKeyInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ScanKeyInit
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ss_get_location github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ss_get_location
func F_ss_get_location(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ss_report_location github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ss_report_location
func F_ss_report_location(m *base.Module, l0 int32, l1 int32)
//go:linkname F_TidStoreCreateLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TidStoreCreateLocal
func F_TidStoreCreateLocal(m *base.Module, l0 int32) int32
//go:linkname F_TidStoreCreateShared github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TidStoreCreateShared
func F_TidStoreCreateShared(m *base.Module, l0 int32) int32
//go:linkname F_TidStoreDestroy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TidStoreDestroy
func F_TidStoreDestroy(m *base.Module, l0 int32)
//go:linkname F_TidStoreSetBlockOffsets github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TidStoreSetBlockOffsets
func F_TidStoreSetBlockOffsets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_toast_open_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_open_indexes
func F_toast_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_toast_delete_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_toast_delete_datum
func F_toast_delete_datum(m *base.Module, l0 int32, l1 int32)
//go:linkname F_convert_tuples_by_position github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_convert_tuples_by_position
func F_convert_tuples_by_position(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_convert_tuples_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_convert_tuples_by_name
func F_convert_tuples_by_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_execute_attr_map_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_execute_attr_map_tuple
func F_execute_attr_map_tuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_execute_attr_map_slot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_execute_attr_map_slot
func F_execute_attr_map_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_free_conversion_map github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_free_conversion_map
func F_free_conversion_map(m *base.Module, l0 int32)
//go:linkname F_populate_compact_attribute github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_populate_compact_attribute
func F_populate_compact_attribute(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateTemplateTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateTemplateTupleDesc
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32
//go:linkname F_CreateTupleDescCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTupleDescCopy
func F_CreateTupleDescCopy(m *base.Module, l0 int32) int32
//go:linkname F_FreeTupleDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTupleDesc
func F_FreeTupleDesc(m *base.Module, l0 int32)
//go:linkname F_IncrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IncrTupleDescRefCount
func F_IncrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_DecrTupleDescRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecrTupleDescRefCount
func F_DecrTupleDescRefCount(m *base.Module, l0 int32)
//go:linkname F_TupleDescInitEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TupleDescInitEntry
func F_TupleDescInitEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_TupleDescInitBuiltinEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TupleDescInitBuiltinEntry
func F_TupleDescInitBuiltinEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ginFinishSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginFinishSplit
func F_ginFinishSplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ginPlaceToPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginPlaceToPage
func F_ginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ginInitBA github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ginInitBA
func F_ginInitBA(m *base.Module, l0 int32)
//go:linkname F_GinBufferInit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GinBufferInit
func F_GinBufferInit(m *base.Module, l0 int32) int32
//go:linkname F_GinBufferStoreTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GinBufferStoreTuple
func F_GinBufferStoreTuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ginFlushBuildState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ginFlushBuildState
func F_ginFlushBuildState(m *base.Module, l0 int32, l1 int32)
//go:linkname F__gin_build_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__gin_build_tuple
func F__gin_build_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ginFreeScanKeys github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ginFreeScanKeys
func F_ginFreeScanKeys(m *base.Module, l0 int32)
//go:linkname F_initGinState github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initGinState
func F_initGinState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistFindCorrectParent github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gistFindCorrectParent
func F_gistFindCorrectParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistSplit
func F_gistSplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_gistinserttuples github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistinserttuples
func F_gistinserttuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_gistProcessItup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gistProcessItup
func F_gistProcessItup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistLoadNodeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistLoadNodeBuffer
func F_gistLoadNodeBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_index_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_getattr_2
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistfillbuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistfillbuffer
func F_gistfillbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_gistextractpage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_gistextractpage
func F_gistextractpage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gistjoinvector github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistjoinvector
func F_gistjoinvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_gistNewBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistNewBuffer
func F_gistNewBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_gistGetFakeLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_gistGetFakeLSN
func F_gistGetFakeLSN(m *base.Module, l0 int32) int64
//go:linkname F_gistvacuumscan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistvacuumscan
func F_gistvacuumscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_gistRedoClearFollowRight github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_gistRedoClearFollowRight
func F_gistRedoClearFollowRight(m *base.Module, l0 int32, l1 int32)
//go:linkname F_gistXLogUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gistXLogUpdate
func F_gistXLogUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64
//go:linkname F_hashbucketcleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hashbucketcleanup
func F_hashbucketcleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F__hash_addovflpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_addovflpage
func F__hash_addovflpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_getbuf
func F__hash_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_getnewbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_getnewbuf
func F__hash_getnewbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__hash_splitbucket github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__hash_splitbucket
func F__hash_splitbucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F__hash_getbucketbuf_from_hashkey github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__hash_getbucketbuf_from_hashkey
func F__hash_getbucketbuf_from_hashkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__hash_readpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_readpage
func F__hash_readpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__hash_readnext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__hash_readnext
func F__hash_readnext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__hash_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__hash_checkpage
func F__hash_checkpage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_HeapCheckForSerializableConflictOut github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapCheckForSerializableConflictOut
func F_HeapCheckForSerializableConflictOut(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_heap_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getnext
func F_heap_getnext(m *base.Module, l0 int32) int32
//go:linkname F_heapgettup_pagemode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heapgettup_pagemode
func F_heapgettup_pagemode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heapgettup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heapgettup
func F_heapgettup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_heap_getattr_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_getattr_1
func F_heap_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_log_heap_new_cid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_heap_new_cid
func F_log_heap_new_cid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DoesMultiXactIdConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DoesMultiXactIdConflict
func F_DoesMultiXactIdConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_Do_MultiXactIdWait github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_Do_MultiXactIdWait
func F_Do_MultiXactIdWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExtractReplicaIdentity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExtractReplicaIdentity
func F_ExtractReplicaIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_compute_new_xmax_infomask github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_compute_new_xmax_infomask
func F_compute_new_xmax_infomask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_heap_prepare_freeze_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_heap_prepare_freeze_tuple
func F_heap_prepare_freeze_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_heap_tuple_should_freeze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_tuple_should_freeze
func F_heap_tuple_should_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_HeapTupleSetHintBits github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_HeapTupleSetHintBits
func F_HeapTupleSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_HeapTupleSatisfiesUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_HeapTupleSatisfiesUpdate
func F_HeapTupleSatisfiesUpdate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_HeapTupleSatisfiesVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleSatisfiesVacuum
func F_HeapTupleSatisfiesVacuum(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_HeapTupleHeaderIsOnlyLocked github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleHeaderIsOnlyLocked
func F_HeapTupleHeaderIsOnlyLocked(m *base.Module, l0 int32) int32
//go:linkname F_HeapTupleSatisfiesVisibility github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_HeapTupleSatisfiesVisibility
func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_heap_toast_insert_or_update github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_heap_toast_insert_or_update
func F_heap_toast_insert_or_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RelationPutHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationPutHeapTuple
func F_RelationPutHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RelationGetBufferForTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetBufferForTuple
func F_RelationGetBufferForTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_heap_page_prune_and_freeze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_heap_page_prune_and_freeze
func F_heap_page_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_heap_get_root_tuples github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_heap_get_root_tuples
func F_heap_get_root_tuples(m *base.Module, l0 int32, l1 int32)
//go:linkname F_raw_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_raw_heap_insert
func F_raw_heap_insert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logical_heap_rewrite_flush_mappings github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_logical_heap_rewrite_flush_mappings
func F_logical_heap_rewrite_flush_mappings(m *base.Module, l0 int32)
//go:linkname F_lazy_check_wraparound_failsafe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lazy_check_wraparound_failsafe
func F_lazy_check_wraparound_failsafe(m *base.Module, l0 int32) int32
//go:linkname F_lazy_vacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lazy_vacuum
func F_lazy_vacuum(m *base.Module, l0 int32)
//go:linkname F_visibilitymap_clear github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_clear
func F_visibilitymap_clear(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_pin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_visibilitymap_pin
func F_visibilitymap_pin(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_vm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vm_readbuf
func F_vm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_visibilitymap_set github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_visibilitymap_set
func F_visibilitymap_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_visibilitymap_get_status github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_visibilitymap_get_status
func F_visibilitymap_get_status(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetIndexAmRoutine github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetIndexAmRoutine
func F_GetIndexAmRoutine(m *base.Module, l0 int32) int32
//go:linkname F_IndexAmTranslateStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IndexAmTranslateStrategy
func F_IndexAmTranslateStrategy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IndexAmTranslateCompareType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IndexAmTranslateCompareType
func F_IndexAmTranslateCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_identify_opfamily_groups github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_identify_opfamily_groups
func F_identify_opfamily_groups(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_amop_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_amop_signature
func F_check_amop_signature(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_opfamily_can_sort_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_opfamily_can_sort_type
func F_opfamily_can_sort_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetIndexScan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationGetIndexScan
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_systable_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_beginscan
func F_systable_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_systable_getnext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_getnext
func F_systable_getnext(m *base.Module, l0 int32) int32
//go:linkname F_systable_recheck_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_recheck_tuple
func F_systable_recheck_tuple(m *base.Module, l0 int32) int32
//go:linkname F_systable_endscan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_systable_endscan
func F_systable_endscan(m *base.Module, l0 int32)
//go:linkname F_systable_inplace_update_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_inplace_update_begin
func F_systable_inplace_update_begin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_systable_inplace_update_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_systable_inplace_update_finish
func F_systable_inplace_update_finish(m *base.Module, l0 int32, l1 int32)
//go:linkname F_systable_inplace_update_cancel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_systable_inplace_update_cancel
func F_systable_inplace_update_cancel(m *base.Module, l0 int32)
//go:linkname F_index_open github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_open
func F_index_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_index_insert
func F_index_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_index_beginscan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_beginscan
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_index_rescan
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_index_parallelscan_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_parallelscan_estimate
func F_index_parallelscan_estimate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_index_vacuum_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_vacuum_cleanup
func F_index_vacuum_cleanup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_index_getprocinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_index_getprocinfo
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_btoidvectorcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_btoidvectorcmp
func F_btoidvectorcmp(m *base.Module, l0 int32) int32
//go:linkname F__bt_update_posting github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_update_posting
func F__bt_update_posting(m *base.Module, l0 int32)
//go:linkname F__bt_getstackbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_getstackbuf
func F__bt_getstackbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_checkpage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_checkpage
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_getbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_getbuf
func F__bt_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_relbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_relbuf
func F__bt_relbuf(m *base.Module, l0 int32)
//go:linkname F__bt_unlockbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_unlockbuf
func F__bt_unlockbuf(m *base.Module, l0 int32)
//go:linkname F__bt_getroot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_getroot
func F__bt_getroot(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_relandgetbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_relandgetbuf
func F__bt_relandgetbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_metaversion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_metaversion
func F__bt_metaversion(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__bt_leftsib_splitflag github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_leftsib_splitflag
func F__bt_leftsib_splitflag(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_parallel_done github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_parallel_done
func F__bt_parallel_done(m *base.Module, l0 int32)
//go:linkname F__bt_parallel_release github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_parallel_release
func F__bt_parallel_release(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__bt_moveright github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_moveright
func F__bt_moveright(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F__bt_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_compare
func F__bt_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_readpage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_readpage
func F__bt_readpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F__bt_end_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__bt_end_parallel
func F__bt_end_parallel(m *base.Module, l0 int32)
//go:linkname F__bt_parallel_scan_and_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_parallel_scan_and_sort
func F__bt_parallel_scan_and_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F__bt_buildadd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_buildadd
func F__bt_buildadd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__bt_sort_dedup_finish_pending github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_sort_dedup_finish_pending
func F__bt_sort_dedup_finish_pending(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__bt_start_array_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__bt_start_array_keys
func F__bt_start_array_keys(m *base.Module, l0 int32, l1 int32)
//go:linkname F__bt_check_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_check_compare
func F__bt_check_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F__bt_tuple_before_array_skeys github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__bt_tuple_before_array_skeys
func F__bt_tuple_before_array_skeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F__bt_keep_natts_fast github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__bt_keep_natts_fast
func F__bt_keep_natts_fast(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__bt_allequalimage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__bt_allequalimage
func F__bt_allequalimage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_delvacuum_desc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_delvacuum_desc
func F_delvacuum_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_standby_desc_invalidations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_standby_desc_invalidations
func F_standby_desc_invalidations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_spgPageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_spgPageIndexMultiDelete
func F_spgPageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_saveNodeLink github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_saveNodeLink
func F_saveNodeLink(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_box_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_box_copy
func F_box_copy(m *base.Module, l0 int32) int32
//go:linkname F_spgGetCache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_spgGetCache
func F_spgGetCache(m *base.Module, l0 int32) int32
//go:linkname F_initSpGistState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_initSpGistState
func F_initSpGistState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SpGistUpdateMetaPage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SpGistUpdateMetaPage
func F_SpGistUpdateMetaPage(m *base.Module, l0 int32)
//go:linkname F_SpGistGetBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SpGistGetBuffer
func F_SpGistGetBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SpGistGetLeafTupleSize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SpGistGetLeafTupleSize
func F_SpGistGetLeafTupleSize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_spgFormLeafTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_spgFormLeafTuple
func F_spgFormLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_spgFormInnerTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_spgFormInnerTuple
func F_spgFormInnerTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_spgDeformLeafTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_spgDeformLeafTuple
func F_spgDeformLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_spgExtractNodeLabels github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_spgExtractNodeLabels
func F_spgExtractNodeLabels(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SpGistPageAddNewItem github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SpGistPageAddNewItem
func F_SpGistPageAddNewItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_vacuumLeafPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vacuumLeafPage
func F_vacuumLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_vacuumRedirectAndPlaceholder github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vacuumRedirectAndPlaceholder
func F_vacuumRedirectAndPlaceholder(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sequence_close github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sequence_close
func F_sequence_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_open
func F_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_try_table_open github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_try_table_open
func F_try_table_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_openrv github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_table_openrv
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_slot_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_slot_create
func F_table_slot_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_beginscan_catalog github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_table_beginscan_catalog
func F_table_beginscan_catalog(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_table_parallelscan_estimate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_table_parallelscan_estimate
func F_table_parallelscan_estimate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_table_parallelscan_initialize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_table_parallelscan_initialize
func F_table_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_table_beginscan_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_table_beginscan_parallel
func F_table_beginscan_parallel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TransactionIdGetCommitTsData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdGetCommitTsData
func F_TransactionIdGetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MultiXactIdCreateFromMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MultiXactIdCreateFromMembers
func F_MultiXactIdCreateFromMembers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RecordNewMultiXact github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordNewMultiXact
func F_RecordNewMultiXact(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetMultiXactIdMembers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetMultiXactIdMembers
func F_GetMultiXactIdMembers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MultiXactIdSetOldestMember github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MultiXactIdSetOldestMember
func F_MultiXactIdSetOldestMember(m *base.Module)
//go:linkname F_ReadNextMultiXactId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReadNextMultiXactId
func F_ReadNextMultiXactId(m *base.Module) int32
//go:linkname F_SetMultiXactIdLimit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetMultiXactIdLimit
func F_SetMultiXactIdLimit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateParallelContext
func F_CreateParallelContext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InitializeParallelDSM github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InitializeParallelDSM
func F_InitializeParallelDSM(m *base.Module, l0 int32)
//go:linkname F_LaunchParallelWorkers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LaunchParallelWorkers
func F_LaunchParallelWorkers(m *base.Module, l0 int32)
//go:linkname F_DestroyParallelContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DestroyParallelContext
func F_DestroyParallelContext(m *base.Module, l0 int32)
//go:linkname F_SimpleLruZeroPage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruZeroPage
func F_SimpleLruZeroPage(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_SimpleLruWaitIO github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SimpleLruWaitIO
func F_SimpleLruWaitIO(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SlruInternalWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SlruInternalWritePage
func F_SlruInternalWritePage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SimpleLruReadPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SimpleLruReadPage
func F_SimpleLruReadPage(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SimpleLruWritePage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SimpleLruWritePage
func F_SimpleLruWritePage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SimpleLruTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SimpleLruTruncate
func F_SimpleLruTruncate(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SlruInternalDeleteSegment github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SlruInternalDeleteSegment
func F_SlruInternalDeleteSegment(m *base.Module, l0 int32, l1 int64)
//go:linkname F_SubTransGetTopmostTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SubTransGetTopmostTransaction
func F_SubTransGetTopmostTransaction(m *base.Module, l0 int32) int32
//go:linkname F_TruncateSUBTRANS github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TruncateSUBTRANS
func F_TruncateSUBTRANS(m *base.Module, l0 int32)
//go:linkname F_readTimeLineHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_readTimeLineHistory
func F_readTimeLineHistory(m *base.Module, l0 int32) int32
//go:linkname F_findNewestTimeLine github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_findNewestTimeLine
func F_findNewestTimeLine(m *base.Module, l0 int32) int32
//go:linkname F_tliOfPointInHistory github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tliOfPointInHistory
func F_tliOfPointInHistory(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_tliSwitchPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tliSwitchPoint
func F_tliSwitchPoint(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_TransactionIdDidCommit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransactionIdDidCommit
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32
//go:linkname F_TransactionIdDidAbort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TransactionIdDidAbort
func F_TransactionIdDidAbort(m *base.Module, l0 int32) int32
//go:linkname F_StandbyTransactionIdIsPrepared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StandbyTransactionIdIsPrepared
func F_StandbyTransactionIdIsPrepared(m *base.Module, l0 int32) int32
//go:linkname F_ProcessTwoPhaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessTwoPhaseBuffer
func F_ProcessTwoPhaseBuffer(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ReadNextFullTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadNextFullTransactionId
func F_ReadNextFullTransactionId(m *base.Module) int64
//go:linkname F_AdvanceNextFullTransactionIdPastXid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AdvanceNextFullTransactionIdPastXid
func F_AdvanceNextFullTransactionIdPastXid(m *base.Module, l0 int32)
//go:linkname F_AdvanceOldestClogXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AdvanceOldestClogXid
func F_AdvanceOldestClogXid(m *base.Module, l0 int32)
//go:linkname F_GetNewObjectId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetNewObjectId
func F_GetNewObjectId(m *base.Module) int32
//go:linkname F_GetTopTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTopTransactionId
func F_GetTopTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentTransactionId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetCurrentTransactionId
func F_GetCurrentTransactionId(m *base.Module) int32
//go:linkname F_GetCurrentCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetCurrentCommandId
func F_GetCurrentCommandId(m *base.Module, l0 int32) int32
//go:linkname F_CommandCounterIncrement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommandCounterIncrement
func F_CommandCounterIncrement(m *base.Module)
//go:linkname F_StartTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_StartTransactionCommand
func F_StartTransactionCommand(m *base.Module)
//go:linkname F_StartTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_StartTransaction
func F_StartTransaction(m *base.Module)
//go:linkname F_CommitTransactionCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommitTransactionCommand
func F_CommitTransactionCommand(m *base.Module)
//go:linkname F_AbortCurrentTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_AbortCurrentTransaction
func F_AbortCurrentTransaction(m *base.Module)
//go:linkname F_PreventInTransactionBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PreventInTransactionBlock
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PushTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PushTransaction
func F_PushTransaction(m *base.Module)
//go:linkname F_RollbackAndReleaseCurrentSubTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RollbackAndReleaseCurrentSubTransaction
func F_RollbackAndReleaseCurrentSubTransaction(m *base.Module)
//go:linkname F_AbortOutOfAnyTransaction github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AbortOutOfAnyTransaction
func F_AbortOutOfAnyTransaction(m *base.Module)
//go:linkname F_WALInsertLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WALInsertLockRelease
func F_WALInsertLockRelease(m *base.Module)
//go:linkname F_WALInsertLockAcquireExclusive github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WALInsertLockAcquireExclusive
func F_WALInsertLockAcquireExclusive(m *base.Module)
//go:linkname F_WaitXLogInsertionsToFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitXLogInsertionsToFinish
func F_WaitXLogInsertionsToFinish(m *base.Module, l0 int64) int64
//go:linkname F_XLogWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogWrite
func F_XLogWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RecoveryInProgress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RecoveryInProgress
func F_RecoveryInProgress(m *base.Module) int32
//go:linkname F_XLogFileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogFileClose
func F_XLogFileClose(m *base.Module)
//go:linkname F_XLogFileName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogFileName
func F_XLogFileName(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_AdvanceXLInsertBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AdvanceXLInsertBuffer
func F_AdvanceXLInsertBuffer(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_XLogFileInitInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogFileInitInternal
func F_XLogFileInitInternal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CheckXLogRemoved github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckXLogRemoved
func F_CheckXLogRemoved(m *base.Module, l0 int64, l1 int32)
//go:linkname F_XLogGetOldestSegno github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogGetOldestSegno
func F_XLogGetOldestSegno(m *base.Module, l0 int32) int64
//go:linkname F_RemoveXlogFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RemoveXlogFile
func F_RemoveXlogFile(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)
//go:linkname F_DataChecksumsEnabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DataChecksumsEnabled
func F_DataChecksumsEnabled(m *base.Module) int32
//go:linkname F_GetRedoRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetRedoRecPtr
func F_GetRedoRecPtr(m *base.Module) int64
//go:linkname F_GetFullPageWriteInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetFullPageWriteInfo
func F_GetFullPageWriteInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetLastImportantRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetLastImportantRecPtr
func F_GetLastImportantRecPtr(m *base.Module) int64
//go:linkname F_LogCheckpointStart github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogCheckpointStart
func F_LogCheckpointStart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CheckPointGuts github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckPointGuts
func F_CheckPointGuts(m *base.Module, l0 int64, l1 int32)
//go:linkname F_KeepLogSeg github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_KeepLogSeg
func F_KeepLogSeg(m *base.Module, l0 int64, l1 int32)
//go:linkname F_LogCheckpointEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LogCheckpointEnd
func F_LogCheckpointEnd(m *base.Module, l0 int32)
//go:linkname F_do_pg_backup_start github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_do_pg_backup_start
func F_do_pg_backup_start(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_do_pg_abort_backup github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_do_pg_abort_backup
func F_do_pg_abort_backup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_do_pg_backup_stop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_do_pg_backup_stop
func F_do_pg_backup_stop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetOldestRestartPoint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetOldestRestartPoint
func F_GetOldestRestartPoint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_KeepFileRestoredFromArchive github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_KeepFileRestoredFromArchive
func F_KeepFileRestoredFromArchive(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogArchiveCheckDone github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_XLogArchiveCheckDone
func F_XLogArchiveCheckDone(m *base.Module, l0 int32) int32
//go:linkname F_build_backup_content github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_backup_content
func F_build_backup_content(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_XLogBeginInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_XLogBeginInsert
func F_XLogBeginInsert(m *base.Module)
//go:linkname F_XLogEnsureRecordSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogEnsureRecordSpace
func F_XLogEnsureRecordSpace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogRegisterBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogRegisterBuffer
func F_XLogRegisterBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_XLogRegisterData github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRegisterData
func F_XLogRegisterData(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogInsert
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_log_newpage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_newpage
func F_log_newpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_log_newpage_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_log_newpage_buffer
func F_log_newpage_buffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_XLogReaderValidatePageHeader github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReaderValidatePageHeader
func F_XLogReaderValidatePageHeader(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_XLogRecGetBlockTag github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_XLogRecGetBlockTag
func F_XLogRecGetBlockTag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_SetRecoveryPause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetRecoveryPause
func F_SetRecoveryPause(m *base.Module, l0 int32)
//go:linkname F_PromoteIsTriggered github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PromoteIsTriggered
func F_PromoteIsTriggered(m *base.Module) int32
//go:linkname F_GetXLogReplayRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetXLogReplayRecPtr
func F_GetXLogReplayRecPtr(m *base.Module, l0 int32) int64
//go:linkname F_GetCurrentReplayRecPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentReplayRecPtr
func F_GetCurrentReplayRecPtr(m *base.Module, l0 int32) int64
//go:linkname F_XLogReadBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_XLogReadBufferForRedo
func F_XLogReadBufferForRedo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_XLogReadBufferExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogReadBufferExtended
func F_XLogReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_XLogInitBufferForRedo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_XLogInitBufferForRedo
func F_XLogInitBufferForRedo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AppendStringToManifest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendStringToManifest
func F_AppendStringToManifest(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sendDir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sendDir
func F_sendDir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64
//go:linkname F_sendTablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sendTablespace
func F_sendTablespace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64
//go:linkname F_sendFileWithContent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sendFileWithContent
func F_sendFileWithContent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_sendFile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sendFile
func F_sendFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F__tarWriteHeader github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__tarWriteHeader
func F__tarWriteHeader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_SendXlogRecPtrResult github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SendXlogRecPtrResult
func F_SendXlogRecPtrResult(m *base.Module, l0 int64, l1 int32)
//go:linkname F_bbsink_forward_begin_backup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bbsink_forward_begin_backup
func F_bbsink_forward_begin_backup(m *base.Module, l0 int32)
//go:linkname F_bbsink_forward_manifest_contents github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bbsink_forward_manifest_contents
func F_bbsink_forward_manifest_contents(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetWalSummaries github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetWalSummaries
func F_GetWalSummaries(m *base.Module, l0 int32, l1 int64, l2 int64) int32
//go:linkname F_ExecuteGrantStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecuteGrantStmt
func F_ExecuteGrantStmt(m *base.Module, l0 int32)
//go:linkname F_string_to_privilege github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_string_to_privilege
func F_string_to_privilege(m *base.Module, l0 int32) int64
//go:linkname F_ExecGrantStmt_oids github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecGrantStmt_oids
func F_ExecGrantStmt_oids(m *base.Module, l0 int32)
//go:linkname F_privilege_to_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_privilege_to_string
func F_privilege_to_string(m *base.Module, l0 int64) int32
//go:linkname F_heap_getattr_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_heap_getattr_2
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SetDefaultACL github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetDefaultACL
func F_SetDefaultACL(m *base.Module, l0 int32)
//go:linkname F_aclcheck_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclcheck_error
func F_aclcheck_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_aclcheck_error_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_aclcheck_error_type
func F_aclcheck_error_type(m *base.Module, l0 int32, l1 int32)
//go:linkname F_object_aclmask_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_object_aclmask_ext
func F_object_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64
//go:linkname F_pg_attribute_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_attribute_aclcheck
func F_pg_attribute_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_pg_attribute_aclcheck_all github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_attribute_aclcheck_all
func F_pg_attribute_aclcheck_all(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_class_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_class_aclcheck
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pg_class_aclcheck_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_class_aclcheck_ext
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pg_parameter_aclcheck github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_parameter_aclcheck
func F_pg_parameter_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_object_ownercheck github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_object_ownercheck
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsCatalogRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsCatalogRelation
func F_IsCatalogRelation(m *base.Module, l0 int32) int32
//go:linkname F_GetNewRelFileNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetNewRelFileNumber
func F_GetNewRelFileNumber(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_performDeletion github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_performDeletion
func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_findDependentObjects github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_findDependentObjects
func F_findDependentObjects(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_reportDependentObjects github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_reportDependentObjects
func F_reportDependentObjects(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AcquireDeletionLock github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AcquireDeletionLock
func F_AcquireDeletionLock(m *base.Module, l0 int32)
//go:linkname F_new_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_new_object_addresses
func F_new_object_addresses(m *base.Module) int32
//go:linkname F_doDeletion github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_doDeletion
func F_doDeletion(m *base.Module, l0 int32, l1 int32)
//go:linkname F_free_object_addresses github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_object_addresses
func F_free_object_addresses(m *base.Module, l0 int32)
//go:linkname F_ReleaseDeletionLock github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseDeletionLock
func F_ReleaseDeletionLock(m *base.Module, l0 int32)
//go:linkname F_recordDependencyOnExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recordDependencyOnExpr
func F_recordDependencyOnExpr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordDependencyOnSingleRelExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_recordDependencyOnSingleRelExpr
func F_recordDependencyOnSingleRelExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_add_exact_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_exact_object_address
func F_add_exact_object_address(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AddRelationNewConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AddRelationNewConstraints
func F_AddRelationNewConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_cookDefault github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cookDefault
func F_cookDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_index_build github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_index_build
func F_index_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_FormIndexDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FormIndexDatum
func F_FormIndexDatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_BuildIndexInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BuildIndexInfo
func F_BuildIndexInfo(m *base.Module, l0 int32) int32
//go:linkname F_reindex_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_reindex_relation
func F_reindex_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CatalogOpenIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CatalogOpenIndexes
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32
//go:linkname F_CatalogCloseIndexes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCloseIndexes
func F_CatalogCloseIndexes(m *base.Module, l0 int32)
//go:linkname F_CatalogTupleInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleInsert
func F_CatalogTupleInsert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CatalogTupleUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdate
func F_CatalogTupleUpdate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CatalogTupleUpdateWithInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CatalogTupleUpdateWithInfo
func F_CatalogTupleUpdateWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CatalogTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CatalogTupleDelete
func F_CatalogTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RangeVarGetRelidExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetRelidExtended
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_recomputeNamespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recomputeNamespacePath
func F_recomputeNamespacePath(m *base.Module)
//go:linkname F_get_namespace_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_oid
func F_get_namespace_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RangeVarGetAndCheckCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RangeVarGetAndCheckCreationNamespace
func F_RangeVarGetAndCheckCreationNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RangeVarAdjustRelationPersistence github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RangeVarAdjustRelationPersistence
func F_RangeVarAdjustRelationPersistence(m *base.Module, l0 int32, l1 int32)
//go:linkname F_isAnyTempNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isAnyTempNamespace
func F_isAnyTempNamespace(m *base.Module, l0 int32) int32
//go:linkname F_RelationIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationIsVisibleExt
func F_RelationIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TypeIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TypeIsVisibleExt
func F_TypeIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DeconstructQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeconstructQualifiedName
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_NameListToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NameListToString
func F_NameListToString(m *base.Module, l0 int32) int32
//go:linkname F_OpernameGetOprid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OpernameGetOprid
func F_OpernameGetOprid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OpernameGetCandidates github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OpernameGetCandidates
func F_OpernameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_statistics_object_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_statistics_object_oid
func F_get_statistics_object_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_parser_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_ts_parser_oid
func F_get_ts_parser_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_dict_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_ts_dict_oid
func F_get_ts_dict_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TSDictionaryIsVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TSDictionaryIsVisible
func F_TSDictionaryIsVisible(m *base.Module, l0 int32) int32
//go:linkname F_TSDictionaryIsVisibleExt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TSDictionaryIsVisibleExt
func F_TSDictionaryIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_template_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_template_oid
func F_get_ts_template_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_ts_config_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_ts_config_oid
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LookupNamespaceNoError github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupNamespaceNoError
func F_LookupNamespaceNoError(m *base.Module, l0 int32) int32
//go:linkname F_LookupCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupCreationNamespace
func F_LookupCreationNamespace(m *base.Module, l0 int32) int32
//go:linkname F_QualifiedNameGetCreationNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QualifiedNameGetCreationNamespace
func F_QualifiedNameGetCreationNamespace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeRangeVarFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeRangeVarFromNameList
func F_makeRangeVarFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_isTempToastNamespace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_isTempToastNamespace
func F_isTempToastNamespace(m *base.Module, l0 int32) int32
//go:linkname F_SearchPathMatchesCurrentEnvironment github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchPathMatchesCurrentEnvironment
func F_SearchPathMatchesCurrentEnvironment(m *base.Module, l0 int32) int32
//go:linkname F_get_collation_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_collation_oid
func F_get_collation_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fetch_search_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_search_path
func F_fetch_search_path(m *base.Module, l0 int32) int32
//go:linkname F_RunObjectPostCreateHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostCreateHook
func F_RunObjectPostCreateHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectDropHook github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RunObjectDropHook
func F_RunObjectDropHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_RunObjectPostAlterHook github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RunObjectPostAlterHook
func F_RunObjectPostAlterHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_RunFunctionExecuteHook github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RunFunctionExecuteHook
func F_RunFunctionExecuteHook(m *base.Module, l0 int32)
//go:linkname F_get_object_address github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_address
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_getObjectDescription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getObjectDescription
func F_getObjectDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_object_ownership github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_object_ownership
func F_check_object_ownership(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_get_object_class_descr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_object_class_descr
func F_get_object_class_descr(m *base.Module, l0 int32) int32
//go:linkname F_get_object_catcache_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_catcache_oid
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_attnum_oid
func F_get_object_attnum_oid(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_object_attnum_name
func F_get_object_attnum_name(m *base.Module, l0 int32) int32
//go:linkname F_get_object_attnum_acl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_object_attnum_acl
func F_get_object_attnum_acl(m *base.Module, l0 int32) int32
//go:linkname F_get_object_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_object_type
func F_get_object_type(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_catalog_object_by_oid_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_catalog_object_by_oid_extended
func F_get_catalog_object_by_oid_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_strlist_to_textarray github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_strlist_to_textarray
func F_strlist_to_textarray(m *base.Module, l0 int32) int32
//go:linkname F_get_partition_parent github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_partition_parent
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_partition_ancestors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_partition_ancestors
func F_get_partition_ancestors(m *base.Module, l0 int32) int32
//go:linkname F_index_get_partition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_index_get_partition
func F_index_get_partition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_map_partition_varattnos github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_map_partition_varattnos
func F_map_partition_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lookup_agg_function github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_agg_function
func F_lookup_agg_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_CastCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CastCreate
func F_CastCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_errdetail_relkind_not_supported github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errdetail_relkind_not_supported
func F_errdetail_relkind_not_supported(m *base.Module, l0 int32)
//go:linkname F_CollationCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CollationCreate
func F_CollationCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32
//go:linkname F_CreateConstraintEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateConstraintEntry
func F_CreateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) int32
//go:linkname F_ConstraintNameIsUsed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConstraintNameIsUsed
func F_ConstraintNameIsUsed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ChooseConstraintName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ChooseConstraintName
func F_ChooseConstraintName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_findNotNullConstraintAttnum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_findNotNullConstraintAttnum
func F_findNotNullConstraintAttnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_extractNotNullColumn github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_extractNotNullColumn
func F_extractNotNullColumn(m *base.Module, l0 int32) int32
//go:linkname F_get_relation_constraint_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_relation_constraint_oid
func F_get_relation_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recordDependencyOn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOn
func F_recordDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recordMultipleDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recordMultipleDependencies
func F_recordMultipleDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_recordDependencyOnCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_recordDependencyOnCurrentExtension
func F_recordDependencyOnCurrentExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_checkMembershipInCurrentExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkMembershipInCurrentExtension
func F_checkMembershipInCurrentExtension(m *base.Module, l0 int32)
//go:linkname F_deleteDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_deleteDependencyRecordsFor
func F_deleteDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deleteDependencyRecordsForClass github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deleteDependencyRecordsForClass
func F_deleteDependencyRecordsForClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_changeDependencyFor github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_changeDependencyFor
func F_changeDependencyFor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getIdentitySequence github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getIdentitySequence
func F_getIdentitySequence(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_index_constraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_index_constraint
func F_get_index_constraint(m *base.Module, l0 int32) int32
//go:linkname F_find_inheritance_children github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_inheritance_children
func F_find_inheritance_children(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_all_inheritors github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_all_inheritors
func F_find_all_inheritors(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_has_superclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_has_superclass
func F_has_superclass(m *base.Module, l0 int32) int32
//go:linkname F_OperatorValidateParams github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OperatorValidateParams
func F_OperatorValidateParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_get_other_operator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_other_operator
func F_get_other_operator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeOperatorDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeOperatorDependencies
func F_makeOperatorDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_OperatorUpd github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OperatorUpd
func F_OperatorUpd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ProcedureCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcedureCreate
func F_ProcedureCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 float32, l28 float32)
//go:linkname F_function_parse_error_transpose github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_function_parse_error_transpose
func F_function_parse_error_transpose(m *base.Module, l0 int32) int32
//go:linkname F_is_schema_publication github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_is_schema_publication
func F_is_schema_publication(m *base.Module, l0 int32) int32
//go:linkname F_GetPubPartitionOptionRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetPubPartitionOptionRelations
func F_GetPubPartitionOptionRelations(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetTopMostAncestorInPublication github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTopMostAncestorInPublication
func F_GetTopMostAncestorInPublication(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetRelationPublications github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetRelationPublications
func F_GetRelationPublications(m *base.Module, l0 int32) int32
//go:linkname F_GetSchemaPublications github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetSchemaPublications
func F_GetSchemaPublications(m *base.Module, l0 int32) int32
//go:linkname F_pub_collist_validate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pub_collist_validate
func F_pub_collist_validate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_publication_add_schema github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_publication_add_schema
func F_publication_add_schema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetPublicationRelations
func F_GetPublicationRelations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPublicationSchemas github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetPublicationSchemas
func F_GetPublicationSchemas(m *base.Module, l0 int32) int32
//go:linkname F_GetAllSchemaPublicationRelations github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetAllSchemaPublicationRelations
func F_GetAllSchemaPublicationRelations(m *base.Module, l0 int32) int32
//go:linkname F_shdepLockAndCheckObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shdepLockAndCheckObject
func F_shdepLockAndCheckObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_changeDependencyOnOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_changeDependencyOnOwner
func F_changeDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_updateInitAclDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_updateInitAclDependencies
func F_updateInitAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_deleteSharedDependencyRecordsFor github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_deleteSharedDependencyRecordsFor
func F_deleteSharedDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetPublicationsStr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPublicationsStr
func F_GetPublicationsStr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetSubscriptionRelState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetSubscriptionRelState
func F_GetSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GenerateTypeDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GenerateTypeDependencies
func F_GenerateTypeDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_TypeCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_TypeCreate
func F_TypeCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32)
//go:linkname F_moveArrayTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_moveArrayTypeName
func F_moveArrayTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeArrayTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayTypeName
func F_makeArrayTypeName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationCreateStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationCreateStorage
func F_RelationCreateStorage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_log_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_log_smgrcreate
func F_log_smgrcreate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationDropStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationDropStorage
func F_RelationDropStorage(m *base.Module, l0 int32)
//go:linkname F_RelationTruncate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationTruncate
func F_RelationTruncate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationCopyStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationCopyStorage
func F_RelationCopyStorage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_NewRelationCreateToastTable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_NewRelationCreateToastTable
func F_NewRelationCreateToastTable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_parse_analyze_fixedparams github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_analyze_fixedparams
func F_parse_analyze_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformStmt
func F_transformStmt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeAConst github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeAConst
func F_makeAConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeStringConstCast github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeStringConstCast
func F_makeStringConstCast(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeRangeVarFromAnyName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeRangeVarFromAnyName
func F_makeRangeVarFromAnyName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SplitColQualList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SplitColQualList
func F_SplitColQualList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_check_func_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_func_name
func F_check_func_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_preprocess_pubobj_list github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_preprocess_pubobj_list
func F_preprocess_pubobj_list(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeRangeVarFromQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeRangeVarFromQualifiedName
func F_makeRangeVarFromQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeRecursiveViewSelect github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeRecursiveViewSelect
func F_makeRecursiveViewSelect(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_indirection github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_indirection
func F_check_indirection(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_insertSelectOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_insertSelectOptions
func F_insertSelectOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_doNegate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_doNegate
func F_doNegate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SystemFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SystemFuncName
func F_SystemFuncName(m *base.Module, l0 int32) int32
//go:linkname F_makeNotExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeNotExpr
func F_makeNotExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeColumnRef github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeColumnRef
func F_makeColumnRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_transformWhereClause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_transformWhereClause
func F_transformWhereClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_findTargetlistEntrySQL92 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_findTargetlistEntrySQL92
func F_findTargetlistEntrySQL92(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_addTargetToGroupList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addTargetToGroupList
func F_addTargetToGroupList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_coerce_to_target_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_target_type
func F_coerce_to_target_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_can_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_can_coerce_type
func F_can_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_coerce_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_type
func F_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_coerce_to_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_coerce_to_domain
func F_coerce_to_domain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_is_complex_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_is_complex_array
func F_is_complex_array(m *base.Module, l0 int32) int32
//go:linkname F_coerce_to_boolean github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_coerce_to_boolean
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_valid_polymorphic_signature github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_valid_polymorphic_signature
func F_check_valid_polymorphic_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsBinaryCoercible github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IsBinaryCoercible
func F_IsBinaryCoercible(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_expr_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_assign_expr_collations
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformExpr
func F_transformExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_func_get_detail github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_get_detail
func F_func_get_detail(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32
//go:linkname F_func_signature_string github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_func_signature_string
func F_func_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_func_select_candidate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_select_candidate
func F_func_select_candidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_func_match_argtypes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_func_match_argtypes
func F_func_match_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LookupFuncName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupFuncName
func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LookupFuncWithArgs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupFuncWithArgs
func F_LookupFuncWithArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_parsestate
func F_make_parsestate(m *base.Module, l0 int32) int32
//go:linkname F_free_parsestate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_free_parsestate
func F_free_parsestate(m *base.Module, l0 int32)
//go:linkname F_parser_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parser_errposition
func F_parser_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LookupOperName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LookupOperName
func F_LookupOperName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LookupOperWithArgs github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LookupOperWithArgs
func F_LookupOperWithArgs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sort_group_operators github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sort_group_operators
func F_get_sort_group_operators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_oper github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_oper
func F_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_make_oper_cache_key github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_oper_cache_key
func F_make_oper_cache_key(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_check_lateral_ref_ok github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_lateral_ref_ok
func F_check_lateral_ref_ok(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_markNullableIfNeeded github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_markNullableIfNeeded
func F_markNullableIfNeeded(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addRTEPermissionInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRTEPermissionInfo
func F_addRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addRangeTableEntryForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addRangeTableEntryForRelation
func F_addRangeTableEntryForRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_addNSItemToQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_addNSItemToQuery
func F_addNSItemToQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_expandRTE github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expandRTE
func F_expandRTE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_attnumAttName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_attnumAttName
func F_attnumAttName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_attnumTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_attnumTypeId
func F_attnumTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_isQueryUsingTempRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_isQueryUsingTempRelation
func F_isQueryUsingTempRelation(m *base.Module, l0 int32) int32
//go:linkname F_LookupTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupTypeName
func F_LookupTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LookupTypeNameExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LookupTypeNameExtended
func F_LookupTypeNameExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_TypeNameToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TypeNameToString
func F_TypeNameToString(m *base.Module, l0 int32) int32
//go:linkname F_LookupTypeNameOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LookupTypeNameOid
func F_LookupTypeNameOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_typenameType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typenameType
func F_typenameType(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typenameTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_typenameTypeId
func F_typenameTypeId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_typenameTypeIdAndMod github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_typenameTypeIdAndMod
func F_typenameTypeIdAndMod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_TypeNameListToString github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TypeNameListToString
func F_TypeNameListToString(m *base.Module, l0 int32) int32
//go:linkname F_GetColumnDefCollation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetColumnDefCollation
func F_GetColumnDefCollation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_typeTypeId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typeTypeId
func F_typeTypeId(m *base.Module, l0 int32) int32
//go:linkname F_typeOrDomainTypeRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_typeOrDomainTypeRelid
func F_typeOrDomainTypeRelid(m *base.Module, l0 int32) int32
//go:linkname F_typeStringToTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_typeStringToTypeName
func F_typeStringToTypeName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generateSerialExtraStmts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generateSerialExtraStmts
func F_generateSerialExtraStmts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_transformTableConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformTableConstraint
func F_transformTableConstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformColumnDefinition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_transformColumnDefinition
func F_transformColumnDefinition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_transformIndexConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformIndexConstraints
func F_transformIndexConstraints(m *base.Module, l0 int32)
//go:linkname F_transformIndexStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformIndexStmt
func F_transformIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformStatsStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformStatsStmt
func F_transformStatsStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_transformAlterTableStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_transformAlterTableStmt
func F_transformAlterTableStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_transformPartitionBoundValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_transformPartitionBoundValue
func F_transformPartitionBoundValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_str_udeescape github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_str_udeescape
func F_str_udeescape(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_core_yylex github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_core_yylex
func F_core_yylex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scanner_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_yyerror
func F_scanner_yyerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_scanner_errposition
func F_scanner_errposition(m *base.Module, l0 int32, l1 int32)
//go:linkname F_yy_fatal_error_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_yy_fatal_error_2
func F_yy_fatal_error_2(m *base.Module, l0 int32)
//go:linkname F_scanner_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_scanner_init
func F_scanner_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_scanner_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_scanner_finish
func F_scanner_finish(m *base.Module, l0 int32)
//go:linkname F_downcase_truncate_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_downcase_truncate_identifier
func F_downcase_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_downcase_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_downcase_identifier
func F_downcase_identifier(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_truncate_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_truncate_identifier
func F_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_extractModify github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_extractModify
func F_extractModify(m *base.Module, l0 int32) int32
//go:linkname F_ExecRenameStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecRenameStmt
func F_ExecRenameStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecAlterObjectDependsStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecAlterObjectDependsStmt
func F_ExecAlterObjectDependsStmt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecAlterObjectSchemaStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecAlterObjectSchemaStmt
func F_ExecAlterObjectSchemaStmt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecAlterOwnerStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAlterOwnerStmt
func F_ExecAlterOwnerStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_index_am_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_index_am_oid
func F_get_index_am_oid(m *base.Module, l0 int32) int32
//go:linkname F_examine_attribute github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_examine_attribute
func F_examine_attribute(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_update_attstats github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_update_attstats
func F_update_attstats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_std_typanalyze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_std_typanalyze
func F_std_typanalyze(m *base.Module, l0 int32) int32
//go:linkname F_AsyncExistsPendingNotify github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AsyncExistsPendingNotify
func F_AsyncExistsPendingNotify(m *base.Module, l0 int32) int32
//go:linkname F_AddEventToPendingNotifies github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AddEventToPendingNotifies
func F_AddEventToPendingNotifies(m *base.Module, l0 int32)
//go:linkname F_queue_listen github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_queue_listen
func F_queue_listen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_asyncQueueUnregister github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_asyncQueueUnregister
func F_asyncQueueUnregister(m *base.Module)
//go:linkname F_CommentObject github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CommentObject
func F_CommentObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GetComment github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetComment
func F_GetComment(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CopyGetAttnums github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyGetAttnums
func F_CopyGetAttnums(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessCopyOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ProcessCopyOptions
func F_ProcessCopyOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CopyFrom github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CopyFrom
func F_CopyFrom(m *base.Module, l0 int32) int64
//go:linkname F_CopyLoadRawBuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CopyLoadRawBuf
func F_CopyLoadRawBuf(m *base.Module, l0 int32)
//go:linkname F_CreateTableAsRelExists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateTableAsRelExists
func F_CreateTableAsRelExists(m *base.Module, l0 int32) int32
//go:linkname F_create_ctas_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_ctas_internal
func F_create_ctas_internal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_db_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_db_info
func F_get_db_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32) int32
//go:linkname F_check_encoding_locale_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_encoding_locale_matches
func F_check_encoding_locale_matches(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_database_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_database_oid
func F_get_database_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errdetail_busy_db github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdetail_busy_db
func F_errdetail_busy_db(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_db_file_conflict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_db_file_conflict
func F_check_db_file_conflict(m *base.Module, l0 int32) int32
//go:linkname F_createdb_failure_callback github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_createdb_failure_callback
func F_createdb_failure_callback(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateDirAndVersionFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateDirAndVersionFile
func F_CreateDirAndVersionFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_defGetString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetString
func F_defGetString(m *base.Module, l0 int32) int32
//go:linkname F_defGetNumeric github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetNumeric
func F_defGetNumeric(m *base.Module, l0 int32) float64
//go:linkname F_defGetBoolean github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetBoolean
func F_defGetBoolean(m *base.Module, l0 int32) int32
//go:linkname F_defGetInt32 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_defGetInt32
func F_defGetInt32(m *base.Module, l0 int32) int32
//go:linkname F_defGetQualifiedName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_defGetQualifiedName
func F_defGetQualifiedName(m *base.Module, l0 int32) int32
//go:linkname F_errorConflictingDefElem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errorConflictingDefElem
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32)
//go:linkname F_owningrel_does_not_exist_skipping github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_owningrel_does_not_exist_skipping
func F_owningrel_does_not_exist_skipping(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_EventTriggerInvoke github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EventTriggerInvoke
func F_EventTriggerInvoke(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EventTriggerSQLDrop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EventTriggerSQLDrop
func F_EventTriggerSQLDrop(m *base.Module, l0 int32)
//go:linkname F_EventTriggerSQLDropAddObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EventTriggerSQLDropAddObject
func F_EventTriggerSQLDropAddObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EventTriggerCollectSimpleCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EventTriggerCollectSimpleCommand
func F_EventTriggerCollectSimpleCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_EventTriggerAlterTableStart github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_EventTriggerAlterTableStart
func F_EventTriggerAlterTableStart(m *base.Module, l0 int32)
//go:linkname F_EventTriggerAlterTableEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EventTriggerAlterTableEnd
func F_EventTriggerAlterTableEnd(m *base.Module)
//go:linkname F_EventTriggerCollectAlterOpFam github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EventTriggerCollectAlterOpFam
func F_EventTriggerCollectAlterOpFam(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_EventTriggerCollectAlterTSConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EventTriggerCollectAlterTSConfig
func F_EventTriggerCollectAlterTSConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExplainIndentText github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExplainIndentText
func F_ExplainIndentText(m *base.Module, l0 int32)
//go:linkname F_ExplainPropertyText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExplainPropertyText
func F_ExplainPropertyText(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExplainProperty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExplainProperty
func F_ExplainProperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExplainPropertyUInteger github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExplainPropertyUInteger
func F_ExplainPropertyUInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_check_valid_extension_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_valid_extension_name
func F_check_valid_extension_name(m *base.Module, l0 int32)
//go:linkname F_parse_extension_control_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_extension_control_file
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_ext_ver_list github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_ext_ver_list
func F_get_ext_ver_list(m *base.Module, l0 int32) int32
//go:linkname F_find_update_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_find_update_path
func F_find_update_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_InsertExtensionTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InsertExtensionTuple
func F_InsertExtensionTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_execute_extension_script github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_execute_extension_script
func F_execute_extension_script(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ApplyExtensionUpdates github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ApplyExtensionUpdates
func F_ApplyExtensionUpdates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_extension_file_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_extension_file_exists
func F_extension_file_exists(m *base.Module, l0 int32) int32
//go:linkname F_ExecAlterExtensionContentsRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecAlterExtensionContentsRecurse
func F_ExecAlterExtensionContentsRecurse(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_transformGenericOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_transformGenericOptions
func F_transformGenericOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_parse_func_options github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_func_options
func F_parse_func_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_interpret_function_parameter_list github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_interpret_function_parameter_list
func F_interpret_function_parameter_list(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32)
//go:linkname F_interpret_func_support github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_interpret_func_support
func F_interpret_func_support(m *base.Module, l0 int32) int32
//go:linkname F_interpret_func_parallel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_interpret_func_parallel
func F_interpret_func_parallel(m *base.Module, l0 int32) int32
//go:linkname F_get_transform_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_transform_oid
func F_get_transform_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_transform_function github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_transform_function
func F_check_transform_function(m *base.Module, l0 int32)
//go:linkname F_DefineIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DefineIndex
func F_DefineIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_GetDefaultOpClass github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDefaultOpClass
func F_GetDefaultOpClass(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReindexRelationConcurrently github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReindexRelationConcurrently
func F_ReindexRelationConcurrently(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ReindexMultipleInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReindexMultipleInternal
func F_ReindexMultipleInternal(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RefreshMatViewByOid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RefreshMatViewByOid
func F_RefreshMatViewByOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_get_opfamily_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_opfamily_oid
func F_get_opfamily_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_opclass_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opclass_oid
func F_get_opclass_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CreateOpFamily github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateOpFamily
func F_CreateOpFamily(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_assignOperTypes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_assignOperTypes
func F_assignOperTypes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addFamilyMember github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addFamilyMember
func F_addFamilyMember(m *base.Module, l0 int32, l1 int32)
//go:linkname F_processTypesSpec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_processTypesSpec
func F_processTypesSpec(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_assignProcTypes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_assignProcTypes
func F_assignProcTypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_storeProcedures github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_storeProcedures
func F_storeProcedures(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ValidateRestrictionEstimator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ValidateRestrictionEstimator
func F_ValidateRestrictionEstimator(m *base.Module, l0 int32) int32
//go:linkname F_ValidateJoinEstimator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ValidateJoinEstimator
func F_ValidateJoinEstimator(m *base.Module, l0 int32) int32
//go:linkname F_ValidateOperatorReference github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ValidateOperatorReference
func F_ValidateOperatorReference(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_policy_role_list_to_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_policy_role_list_to_array
func F_policy_role_list_to_array(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecuteQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecuteQuery
func F_ExecuteQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_FetchPreparedStatement github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FetchPreparedStatement
func F_FetchPreparedStatement(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_language_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_language_oid
func F_get_language_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_publication_options github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_publication_options
func F_parse_publication_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_ObjectsInPublicationToOids github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ObjectsInPublicationToOids
func F_ObjectsInPublicationToOids(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_OpenTableList github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpenTableList
func F_OpenTableList(m *base.Module, l0 int32) int32
//go:linkname F_TransformPubWhereClauses github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_TransformPubWhereClauses
func F_TransformPubWhereClauses(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckPubRelationColumnList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckPubRelationColumnList
func F_CheckPubRelationColumnList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_InvalidatePublicationRels github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InvalidatePublicationRels
func F_InvalidatePublicationRels(m *base.Module, l0 int32)
//go:linkname F_PublicationDropTables github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PublicationDropTables
func F_PublicationDropTables(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateSchemaCommand github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateSchemaCommand
func F_CreateSchemaCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecSecLabelStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecSecLabelStmt
func F_ExecSecLabelStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DeleteSharedSecurityLabel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DeleteSharedSecurityLabel
func F_DeleteSharedSecurityLabel(m *base.Module, l0 int32, l1 int32)
//go:linkname F_init_params github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_init_params
func F_init_params(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_fill_seq_with_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fill_seq_with_data
func F_fill_seq_with_data(m *base.Module, l0 int32, l1 int32)
//go:linkname F_process_owned_by github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_process_owned_by
func F_process_owned_by(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_seq_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_seq_tuple
func F_read_seq_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_nextval_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nextval_internal
func F_nextval_internal(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_CreateStatistics github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateStatistics
func F_CreateStatistics(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateSubscription github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateSubscription
func F_CreateSubscription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_AlterSubscription github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AlterSubscription
func F_AlterSubscription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ReplicationSlotDropAtPubNode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotDropAtPubNode
func F_ReplicationSlotDropAtPubNode(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DropSubscription github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DropSubscription
func F_DropSubscription(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DefineRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DefineRelation
func F_DefineRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_CheckTableNotInUse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CheckTableNotInUse
func F_CheckTableNotInUse(m *base.Module, l0 int32, l1 int32)
//go:linkname F_SetRelationHasSubclass github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SetRelationHasSubclass
func F_SetRelationHasSubclass(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_attnotnull github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_attnotnull
func F_set_attnotnull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetAttributeCompression github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetAttributeCompression
func F_GetAttributeCompression(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetAttributeStorage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetAttributeStorage
func F_GetAttributeStorage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addFkConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addFkConstraint
func F_addFkConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32)
//go:linkname F_CheckRelationTableSpaceMove github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckRelationTableSpaceMove
func F_CheckRelationTableSpaceMove(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetRelationTableSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SetRelationTableSpace
func F_SetRelationTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATController github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATController
func F_ATController(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_AlterTableInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AlterTableInternal
func F_AlterTableInternal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AlterTableGetLockLevel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AlterTableGetLockLevel
func F_AlterTableGetLockLevel(m *base.Module, l0 int32) int32
//go:linkname F_check_of_type github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_check_of_type
func F_check_of_type(m *base.Module, l0 int32)
//go:linkname F_ATSimplePermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATSimplePermissions
func F_ATSimplePermissions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ATExecValidateConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ATExecValidateConstraint
func F_ATExecValidateConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_createForeignKeyActionTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createForeignKeyActionTriggers
func F_createForeignKeyActionTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_directory_is_empty github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_directory_is_empty
func F_directory_is_empty(m *base.Module, l0 int32) int32
//go:linkname F_PrepareTempTablespaces github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareTempTablespaces
func F_PrepareTempTablespaces(m *base.Module)
//go:linkname F_get_tablespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_tablespace_name
func F_get_tablespace_name(m *base.Module, l0 int32) int32
//go:linkname F_CreateTrigger github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateTrigger
func F_CreateTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_FreeTriggerDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeTriggerDesc
func F_FreeTriggerDesc(m *base.Module, l0 int32)
//go:linkname F_before_stmt_triggers_fired github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_stmt_triggers_fired
func F_before_stmt_triggers_fired(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_TriggerEnabled github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TriggerEnabled
func F_TriggerEnabled(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecCallTriggerFunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCallTriggerFunc
func F_ExecCallTriggerFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_AfterTriggerSaveEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AfterTriggerSaveEvent
func F_AfterTriggerSaveEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)
//go:linkname F_afterTriggerAddEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_afterTriggerAddEvent
func F_afterTriggerAddEvent(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecBRInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRInsertTriggers
func F_ExecBRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecIRInsertTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecIRInsertTriggers
func F_ExecIRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecBRUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecBRUpdateTriggers
func F_ExecBRUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_ExecARUpdateTriggers github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecARUpdateTriggers
func F_ExecARUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_AfterTriggerEndQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AfterTriggerEndQuery
func F_AfterTriggerEndQuery(m *base.Module, l0 int32)
//go:linkname F_get_ts_parser_func github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_ts_parser_func
func F_get_ts_parser_func(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_verify_dictoptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_verify_dictoptions
func F_verify_dictoptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_serialize_deflist github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_serialize_deflist
func F_serialize_deflist(m *base.Module, l0 int32) int32
//go:linkname F_deserialize_deflist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_deserialize_deflist
func F_deserialize_deflist(m *base.Module, l0 int32) int32
//go:linkname F_get_ts_template_func github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_ts_template_func
func F_get_ts_template_func(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeConfigurationDependencies github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeConfigurationDependencies
func F_makeConfigurationDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_findTypeReceiveFunction github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_findTypeReceiveFunction
func F_findTypeReceiveFunction(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_findTypeSendFunction github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_findTypeSendFunction
func F_findTypeSendFunction(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_findTypeTypmodinFunction github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_findTypeTypmodinFunction
func F_findTypeTypmodinFunction(m *base.Module, l0 int32) int32
//go:linkname F_findTypeTypmodoutFunction github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_findTypeTypmodoutFunction
func F_findTypeTypmodoutFunction(m *base.Module, l0 int32) int32
//go:linkname F_findTypeAnalyzeFunction github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_findTypeAnalyzeFunction
func F_findTypeAnalyzeFunction(m *base.Module, l0 int32) int32
//go:linkname F_AssignTypeArrayOid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AssignTypeArrayOid
func F_AssignTypeArrayOid(m *base.Module) int32
//go:linkname F_domainAddNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_domainAddNotNullConstraint
func F_domainAddNotNullConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_checkDomainOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkDomainOwner
func F_checkDomainOwner(m *base.Module, l0 int32)
//go:linkname F_AlterDomainNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterDomainNotNull
func F_AlterDomainNotNull(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AlterDomainAddConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AlterDomainAddConstraint
func F_AlterDomainAddConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_validateDomainCheckConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_validateDomainCheckConstraint
func F_validateDomainCheckConstraint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AlterTypeRecurse github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AlterTypeRecurse
func F_AlterTypeRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_vacuum_get_cutoffs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_vacuum_get_cutoffs
func F_vacuum_get_cutoffs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vac_update_relstats github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_vac_update_relstats
func F_vac_update_relstats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32)
//go:linkname F_vac_open_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vac_open_indexes
func F_vac_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_vac_close_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_vac_close_indexes
func F_vac_close_indexes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_vacuum_delay_point github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vacuum_delay_point
func F_vacuum_delay_point(m *base.Module, l0 int32)
//go:linkname F_vac_cleanup_one_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vac_cleanup_one_index
func F_vac_cleanup_one_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parallel_vacuum_process_all_indexes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parallel_vacuum_process_all_indexes
func F_parallel_vacuum_process_all_indexes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecReScan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecReScan
func F_ExecReScan(m *base.Module, l0 int32)
//go:linkname F_ExecSupportsBackwardScan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecSupportsBackwardScan
func F_ExecSupportsBackwardScan(m *base.Module, l0 int32) int32
//go:linkname F_ExecInitExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExpr
func F_ExecInitExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPushExprSetupSteps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPushExprSetupSteps
func F_ExecPushExprSetupSteps(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInitExprRec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExprRec
func F_ExecInitExprRec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecInitExprWithParams github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecInitExprWithParams
func F_ExecInitExprWithParams(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExpr
func F_ExecPrepareExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareQual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareQual
func F_ExecPrepareQual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecPrepareExprList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecPrepareExprList
func F_ExecPrepareExprList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecReadyInterpretedExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecReadyInterpretedExpr
func F_ExecReadyInterpretedExpr(m *base.Module, l0 int32)
//go:linkname F_ExecEvalSysVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecEvalSysVar
func F_ExecEvalSysVar(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_cached_rowtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_cached_rowtype
func F_get_cached_rowtype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecAggCopyTransValue github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecAggCopyTransValue
func F_ExecAggCopyTransValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_GetJsonBehaviorValueString github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetJsonBehaviorValueString
func F_GetJsonBehaviorValueString(m *base.Module, l0 int32) int32
//go:linkname F_tuplehash_lookup_hash_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplehash_lookup_hash_internal
func F_tuplehash_lookup_hash_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_execTuplesMatchPrepare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_execTuplesMatchPrepare
func F_execTuplesMatchPrepare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_BuildTupleHashTable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BuildTupleHashTable
func F_BuildTupleHashTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32
//go:linkname F_LookupTupleHashEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LookupTupleHashEntry
func F_LookupTupleHashEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecOpenIndices github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecOpenIndices
func F_ExecOpenIndices(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecInsertIndexTuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecInsertIndexTuples
func F_ExecInsertIndexTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_ExecCheckIndexConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCheckIndexConstraints
func F_ExecCheckIndexConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecFilterJunk github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFilterJunk
func F_ExecFilterJunk(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecutorStart github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecutorStart
func F_ExecutorStart(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecutorRun github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecutorRun
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_ExecutorEnd github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorEnd
func F_ExecutorEnd(m *base.Module, l0 int32)
//go:linkname F_ExecutorRewind github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecutorRewind
func F_ExecutorRewind(m *base.Module, l0 int32)
//go:linkname F_InitResultRelInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InitResultRelInfo
func F_InitResultRelInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_ExecPartitionCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecPartitionCheck
func F_ExecPartitionCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ExecConstraints github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecConstraints
func F_ExecConstraints(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecWithCheckOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecWithCheckOptions
func F_ExecWithCheckOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecUpdateLockMode github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecUpdateLockMode
func F_ExecUpdateLockMode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_EvalPlanQualFetchRowMark github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_EvalPlanQualFetchRowMark
func F_EvalPlanQualFetchRowMark(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecInitParallelPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecInitParallelPlan
func F_ExecInitParallelPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32
//go:linkname F_ExecParallelReinitialize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecParallelReinitialize
func F_ExecParallelReinitialize(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecParallelFinish github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecParallelFinish
func F_ExecParallelFinish(m *base.Module, l0 int32)
//go:linkname F_ExecParallelCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelCleanup
func F_ExecParallelCleanup(m *base.Module, l0 int32)
//go:linkname F_ExecFindPartition github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFindPartition
func F_ExecFindPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_CheckAndReportConflict github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CheckAndReportConflict
func F_CheckAndReportConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ExecStoreMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecStoreMinimalTuple
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_MakeTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MakeTupleTableSlot
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MakeSingleTupleTableSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MakeSingleTupleTableSlot
func F_MakeSingleTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecStoreHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecStoreHeapTuple
func F_ExecStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecFetchSlotHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecFetchSlotHeapTuple
func F_ExecFetchSlotHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecFetchSlotMinimalTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecFetchSlotMinimalTuple
func F_ExecFetchSlotMinimalTuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_slot_getsomeattrs_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slot_getsomeattrs_int
func F_slot_getsomeattrs_int(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BuildTupleFromCStrings github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BuildTupleFromCStrings
func F_BuildTupleFromCStrings(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_HeapTupleHeaderGetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HeapTupleHeaderGetDatum
func F_HeapTupleHeaderGetDatum(m *base.Module, l0 int32) int32
//go:linkname F_do_tup_output github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_tup_output
func F_do_tup_output(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_end_tup_output github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_tup_output
func F_end_tup_output(m *base.Module, l0 int32)
//go:linkname F_CreateExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateExecutorState
func F_CreateExecutorState(m *base.Module) int32
//go:linkname F_FreeExecutorState github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeExecutorState
func F_FreeExecutorState(m *base.Module, l0 int32)
//go:linkname F_CreateStandaloneExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateStandaloneExprContext
func F_CreateStandaloneExprContext(m *base.Module) int32
//go:linkname F_MakePerTupleExprContext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_MakePerTupleExprContext
func F_MakePerTupleExprContext(m *base.Module, l0 int32) int32
//go:linkname F_ExecGetResultSlotOps github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetResultSlotOps
func F_ExecGetResultSlotOps(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetReturningSlot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecGetReturningSlot
func F_ExecGetReturningSlot(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetRootToChildMap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetRootToChildMap
func F_ExecGetRootToChildMap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecGetAllUpdatedCols github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecGetAllUpdatedCols
func F_ExecGetAllUpdatedCols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_prepare_sql_fn_parse_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_prepare_sql_fn_parse_info
func F_prepare_sql_fn_parse_info(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CreateCommandName github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateCommandName
func F_CreateCommandName(m *base.Module, l0 int32) int32
//go:linkname F_check_sql_fn_retval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_sql_fn_retval
func F_check_sql_fn_retval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_InstrStopNode github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InstrStopNode
func F_InstrStopNode(m *base.Module, l0 int32, l1 float64)
//go:linkname F_BufferUsageAccumDiff github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufferUsageAccumDiff
func F_BufferUsageAccumDiff(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initialize_hash_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initialize_hash_entry
func F_initialize_hash_entry(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashagg_spill_init github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hashagg_spill_init
func F_hashagg_spill_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64)
//go:linkname F_hashagg_spill_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hashagg_spill_tuple
func F_hashagg_spill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecBitmapIndexScanRetrieveInstrumentation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecBitmapIndexScanRetrieveInstrumentation
func F_ExecBitmapIndexScanRetrieveInstrumentation(m *base.Module, l0 int32)
//go:linkname F_gather_merge_readnext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gather_merge_readnext
func F_gather_merge_readnext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_load_tuple_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_load_tuple_array
func F_load_tuple_array(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ExecParallelHashIncreaseNumBatches github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelHashIncreaseNumBatches
func F_ExecParallelHashIncreaseNumBatches(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashIncreaseNumBuckets github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecParallelHashIncreaseNumBuckets
func F_ExecParallelHashIncreaseNumBuckets(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashEnsureBatchAccessors github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecParallelHashEnsureBatchAccessors
func F_ExecParallelHashEnsureBatchAccessors(m *base.Module, l0 int32)
//go:linkname F_ExecParallelHashMergeCounters github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashMergeCounters
func F_ExecParallelHashMergeCounters(m *base.Module, l0 int32)
//go:linkname F_dense_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dense_alloc
func F_dense_alloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ExecHashIncreaseNumBatches github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecHashIncreaseNumBatches
func F_ExecHashIncreaseNumBatches(m *base.Module, l0 int32)
//go:linkname F_ExecHashTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecHashTableInsert
func F_ExecHashTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ExecChooseHashTableSize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ExecChooseHashTableSize
func F_ExecChooseHashTableSize(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_ExecParallelHashTupleAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecParallelHashTupleAlloc
func F_ExecParallelHashTupleAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecHashTableDetachBatch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecHashTableDetachBatch
func F_ExecHashTableDetachBatch(m *base.Module, l0 int32)
//go:linkname F_ExecComputeStoredGenerated github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecComputeStoredGenerated
func F_ExecComputeStoredGenerated(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_ExecProcessReturning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecProcessReturning
func F_ExecProcessReturning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ExecUpdate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExecUpdate
func F_ExecUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ExecPendingInserts github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ExecPendingInserts
func F_ExecPendingInserts(m *base.Module, l0 int32)
//go:linkname F_ExecCheckTupleVisible github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ExecCheckTupleVisible
func F_ExecCheckTupleVisible(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SeqNext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SeqNext
func F_SeqNext(m *base.Module, l0 int32) int32
//go:linkname F_setop_load_group github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_setop_load_group
func F_setop_load_group(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setop_compare_slots github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setop_compare_slots
func F_setop_compare_slots(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findPartialMatch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_findPartialMatch
func F_findPartialMatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ExecSetParamPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ExecSetParamPlan
func F_ExecSetParamPlan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_update_frameheadpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_update_frameheadpos
func F_update_frameheadpos(m *base.Module, l0 int32)
//go:linkname F_window_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_window_gettupleslot
func F_window_gettupleslot(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_row_is_in_frame github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_row_is_in_frame
func F_row_is_in_frame(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_update_frametailpos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_update_frametailpos
func F_update_frametailpos(m *base.Module, l0 int32)
//go:linkname F_update_grouptailpos github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_update_grouptailpos
func F_update_grouptailpos(m *base.Module, l0 int32)
//go:linkname F_SPI_connect_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SPI_connect_ext
func F_SPI_connect_ext(m *base.Module, l0 int32)
//go:linkname F_SPI_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_execute
func F_SPI_execute(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SPI_freetuptable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_freetuptable
func F_SPI_freetuptable(m *base.Module, l0 int32)
//go:linkname F_SPI_execute_plan_with_paramlist github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SPI_execute_plan_with_paramlist
func F_SPI_execute_plan_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_copytuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SPI_copytuple
func F_SPI_copytuple(m *base.Module, l0 int32) int32
//go:linkname F_SPI_getbinval github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SPI_getbinval
func F_SPI_getbinval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_cursor_open_with_paramlist github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SPI_cursor_open_with_paramlist
func F_SPI_cursor_open_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SPI_plan_get_cached_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SPI_plan_get_cached_plan
func F_SPI_plan_get_cached_plan(m *base.Module, l0 int32) int32
//go:linkname F_TupleQueueReaderNext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TupleQueueReaderNext
func F_TupleQueueReaderNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GetForeignDataWrapper github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetForeignDataWrapper
func F_GetForeignDataWrapper(m *base.Module, l0 int32) int32
//go:linkname F_GetForeignDataWrapperByName github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetForeignDataWrapperByName
func F_GetForeignDataWrapperByName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_foreign_data_wrapper_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_foreign_data_wrapper_oid
func F_get_foreign_data_wrapper_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetForeignServerByName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetForeignServerByName
func F_GetForeignServerByName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_foreign_server_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_foreign_server_oid
func F_get_foreign_server_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetFdwRoutine github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetFdwRoutine
func F_GetFdwRoutine(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineByRelId github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetFdwRoutineByRelId
func F_GetFdwRoutineByRelId(m *base.Module, l0 int32) int32
//go:linkname F_GetFdwRoutineForRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetFdwRoutineForRelation
func F_GetFdwRoutineForRelation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dshash_find github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dshash_find
func F_dshash_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dshash_delete_key github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dshash_delete_key
func F_dshash_delete_key(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dshash_seq_next github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dshash_seq_next
func F_dshash_seq_next(m *base.Module, l0 int32) int32
//go:linkname F_secure_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_secure_read
func F_secure_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_free_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_free_auth_file
func F_free_auth_file(m *base.Module, l0 int32)
//go:linkname F_tokenize_auth_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tokenize_auth_file
func F_tokenize_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_parse_hba_line github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_hba_line
func F_parse_hba_line(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_startmsgread github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_startmsgread
func F_pq_startmsgread(m *base.Module)
//go:linkname F_pq_getmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_getmessage
func F_pq_getmessage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_internal_flush_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internal_flush_buffer
func F_internal_flush_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_beginmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_beginmessage
func F_pq_beginmessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_sendbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendbytes
func F_pq_sendbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_sendtext
func F_pq_sendtext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_sendstring github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_sendstring
func F_pq_sendstring(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pq_endmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_endmessage
func F_pq_endmessage(m *base.Module, l0 int32)
//go:linkname F_pq_begintypsend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pq_begintypsend
func F_pq_begintypsend(m *base.Module, l0 int32)
//go:linkname F_pq_puttextmessage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_puttextmessage
func F_pq_puttextmessage(m *base.Module)
//go:linkname F_pq_putemptymessage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_putemptymessage
func F_pq_putemptymessage(m *base.Module, l0 int32)
//go:linkname F_pq_getmsgbyte github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pq_getmsgbyte
func F_pq_getmsgbyte(m *base.Module, l0 int32) int32
//go:linkname F_pq_copymsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pq_copymsgbytes
func F_pq_copymsgbytes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pq_getmsgfloat8 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pq_getmsgfloat8
func F_pq_getmsgfloat8(m *base.Module, l0 int32) float64
//go:linkname F_pq_getmsgbytes github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgbytes
func F_pq_getmsgbytes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pq_getmsgtext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_getmsgtext
func F_pq_getmsgtext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pq_redirect_to_shm_mq github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pq_redirect_to_shm_mq
func F_pq_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bms_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_copy
func F_bms_copy(m *base.Module, l0 int32) int32
//go:linkname F_bms_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bms_equal
func F_bms_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_make_singleton github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_bms_make_singleton
func F_bms_make_singleton(m *base.Module, l0 int32) int32
//go:linkname F_bms_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_free
func F_bms_free(m *base.Module, l0 int32)
//go:linkname F_bms_is_subset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_subset
func F_bms_is_subset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_is_member github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_is_member
func F_bms_is_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_overlap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_overlap
func F_bms_overlap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_singleton_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_singleton_member
func F_bms_singleton_member(m *base.Module, l0 int32) int32
//go:linkname F_bms_add_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_add_member
func F_bms_add_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_del_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_del_member
func F_bms_del_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_members github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bms_add_members
func F_bms_add_members(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bms_add_range github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_bms_add_range
func F_bms_add_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bms_join github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bms_join
func F_bms_join(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_equal
func F_equal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetExtensibleNodeMethods github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetExtensibleNodeMethods
func F_GetExtensibleNodeMethods(m *base.Module, l0 int32) int32
//go:linkname F_list_make1_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_make1_impl
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_make3_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_make3_impl
func F_list_make3_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lappend github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lappend
func F_lappend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_int github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_lappend_int
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lappend_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lappend_oid
func F_lappend_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lcons github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons
func F_lcons(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_new_head_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_new_head_cell
func F_new_head_cell(m *base.Module, l0 int32)
//go:linkname F_lcons_int github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lcons_int
func F_lcons_int(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_concat
func F_list_concat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_copy
func F_list_copy(m *base.Module, l0 int32) int32
//go:linkname F_list_concat_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_concat_copy
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_member
func F_list_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_member_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_member_ptr
func F_list_member_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_nth_cell github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_delete_nth_cell
func F_list_delete_nth_cell(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_free
func F_list_free(m *base.Module, l0 int32)
//go:linkname F_list_delete_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_delete_ptr
func F_list_delete_ptr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_delete_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_list_delete_first
func F_list_delete_first(m *base.Module, l0 int32) int32
//go:linkname F_list_delete_first_n github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_delete_first_n
func F_list_delete_first_n(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_difference github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_difference
func F_list_difference(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_append_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_append_unique
func F_list_append_unique(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_append_unique_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_list_append_unique_oid
func F_list_append_unique_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_concat_unique_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_concat_unique_oid
func F_list_concat_unique_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_free_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_list_free_deep
func F_list_free_deep(m *base.Module, l0 int32)
//go:linkname F_list_copy_head github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_list_copy_head
func F_list_copy_head(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_copy_tail github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_copy_tail
func F_list_copy_tail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_list_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_list_sort
func F_list_sort(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makeA_Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeA_Expr
func F_makeA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeSimpleA_Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeSimpleA_Expr
func F_makeSimpleA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeVar
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_makeVarFromTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeVarFromTargetEntry
func F_makeVarFromTargetEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeTargetEntry
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_flatCopyTargetEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_flatCopyTargetEntry
func F_flatCopyTargetEntry(m *base.Module, l0 int32) int32
//go:linkname F_makeFromExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeFromExpr
func F_makeFromExpr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeConst github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeConst
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_makeBoolConst github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolConst
func F_makeBoolConst(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeBoolExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeBoolExpr
func F_makeBoolExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeRangeVar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeRangeVar
func F_makeRangeVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeNotNullConstraint github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeNotNullConstraint
func F_makeNotNullConstraint(m *base.Module, l0 int32) int32
//go:linkname F_makeTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeTypeName
func F_makeTypeName(m *base.Module, l0 int32) int32
//go:linkname F_makeTypeNameFromNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeTypeNameFromNameList
func F_makeTypeNameFromNameList(m *base.Module, l0 int32) int32
//go:linkname F_makeColumnDef github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeColumnDef
func F_makeColumnDef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeDefElemExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeDefElemExtended
func F_makeDefElemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeFuncCall
func F_makeFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_opclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_opclause
func F_make_opclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_make_andclause github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_andclause
func F_make_andclause(m *base.Module, l0 int32) int32
//go:linkname F_make_orclause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_orclause
func F_make_orclause(m *base.Module, l0 int32) int32
//go:linkname F_make_and_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_and_qual
func F_make_and_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_ands_explicit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_ands_explicit
func F_make_ands_explicit(m *base.Module, l0 int32) int32
//go:linkname F_make_ands_implicit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_ands_implicit
func F_make_ands_implicit(m *base.Module, l0 int32) int32
//go:linkname F_makeGroupingSet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeGroupingSet
func F_makeGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeVacuumRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeVacuumRelation
func F_makeVacuumRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonFormat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeJsonFormat
func F_makeJsonFormat(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonValueExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeJsonValueExpr
func F_makeJsonValueExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_makeJsonKeyValue github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeJsonKeyValue
func F_makeJsonKeyValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeJsonIsPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_makeJsonIsPredicate
func F_makeJsonIsPredicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeJsonTablePathSpec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeJsonTablePathSpec
func F_makeJsonTablePathSpec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_exprType github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_exprType
func F_exprType(m *base.Module, l0 int32) int32
//go:linkname F_expression_returns_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expression_returns_set
func F_expression_returns_set(m *base.Module, l0 int32) int32
//go:linkname F_fix_opfuncids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fix_opfuncids
func F_fix_opfuncids(m *base.Module, l0 int32)
//go:linkname F_set_opfuncid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_opfuncid
func F_set_opfuncid(m *base.Module, l0 int32)
//go:linkname F_check_functions_in_node github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_functions_in_node
func F_check_functions_in_node(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_tree_walker_impl
func F_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_walker_impl
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_range_table_entry_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_table_entry_walker_impl
func F_range_table_entry_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_expression_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expression_tree_mutator_impl
func F_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_query_tree_mutator_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_query_tree_mutator_impl
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_query_or_expression_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_query_or_expression_tree_walker_impl
func F_query_or_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_planstate_tree_walker_impl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_planstate_tree_walker_impl
func F_planstate_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_outToken github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_outToken
func F_outToken(m *base.Module, l0 int32, l1 int32)
//go:linkname F_outBitmapset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_outBitmapset
func F_outBitmapset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bmsToString github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bmsToString
func F_bmsToString(m *base.Module, l0 int32) int32
//go:linkname F_makeParamList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeParamList
func F_makeParamList(m *base.Module, l0 int32) int32
//go:linkname F_copyParamList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copyParamList
func F_copyParamList(m *base.Module, l0 int32) int32
//go:linkname F_elog_node_display github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_elog_node_display
func F_elog_node_display(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_JumbleQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_JumbleQuery
func F_JumbleQuery(m *base.Module, l0 int32) int32
//go:linkname F_AppendJumble32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AppendJumble32
func F_AppendJumble32(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleFuncExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleFuncExpr
func F__jumbleFuncExpr(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleBoolExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleBoolExpr
func F__jumbleBoolExpr(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleFieldStore github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleFieldStore
func F__jumbleFieldStore(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleRelabelType github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleRelabelType
func F__jumbleRelabelType(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleArrayCoerceExpr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleArrayCoerceExpr
func F__jumbleArrayCoerceExpr(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleJsonIsPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleJsonIsPredicate
func F__jumbleJsonIsPredicate(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleJsonTablePath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleJsonTablePath
func F__jumbleJsonTablePath(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleA_Expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleA_Expr
func F__jumbleA_Expr(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleRoleSpec github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleRoleSpec
func F__jumbleRoleSpec(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleA_Indices github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleA_Indices
func F__jumbleA_Indices(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleResTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleResTarget
func F__jumbleResTarget(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleRangeTableSample github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleRangeTableSample
func F__jumbleRangeTableSample(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleTableSampleClause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleTableSampleClause
func F__jumbleTableSampleClause(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleWithClause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleWithClause
func F__jumbleWithClause(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleJsonObjectConstructor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleJsonObjectConstructor
func F__jumbleJsonObjectConstructor(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleJsonArrayQueryConstructor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleJsonArrayQueryConstructor
func F__jumbleJsonArrayQueryConstructor(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleUpdateStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleUpdateStmt
func F__jumbleUpdateStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumblePLAssignStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumblePLAssignStmt
func F__jumblePLAssignStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleCreateSchemaStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleCreateSchemaStmt
func F__jumbleCreateSchemaStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleVariableShowStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleVariableShowStmt
func F__jumbleVariableShowStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleDropTableSpaceStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F__jumbleDropTableSpaceStmt
func F__jumbleDropTableSpaceStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleCreateExtensionStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleCreateExtensionStmt
func F__jumbleCreateExtensionStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleCreateUserMappingStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleCreateUserMappingStmt
func F__jumbleCreateUserMappingStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleAlterUserMappingStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F__jumbleAlterUserMappingStmt
func F__jumbleAlterUserMappingStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleCreateEventTrigStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__jumbleCreateEventTrigStmt
func F__jumbleCreateEventTrigStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleCreateRoleStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__jumbleCreateRoleStmt
func F__jumbleCreateRoleStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F__jumbleCreateSeqStmt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__jumbleCreateSeqStmt
func F__jumbleCreateSeqStmt(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble8 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble8
func F_AppendJumble8(m *base.Module, l0 int32, l1 int32)
//go:linkname F_AppendJumble github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AppendJumble
func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_AppendJumble16 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_AppendJumble16
func F_AppendJumble16(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IsSquashableConstant github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_IsSquashableConstant
func F_IsSquashableConstant(m *base.Module, l0 int32) int32
//go:linkname F_nodeRead github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_nodeRead
func F_nodeRead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strtok github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strtok
func F_pg_strtok(m *base.Module, l0 int32) int32
//go:linkname F_tbm_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tbm_create
func F_tbm_create(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pagetable_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pagetable_insert
func F_pagetable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tbm_union_page github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tbm_union_page
func F_tbm_union_page(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tbm_begin_private_iterate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tbm_begin_private_iterate
func F_tbm_begin_private_iterate(m *base.Module, l0 int32) int32
//go:linkname F_makeFloat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeFloat
func F_makeFloat(m *base.Module, l0 int32) int32
//go:linkname F_makeBoolean github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeBoolean
func F_makeBoolean(m *base.Module, l0 int32) int32
//go:linkname F_makeString github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_makeString
func F_makeString(m *base.Module, l0 int32) int32
//go:linkname F_add_paths_to_append_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_paths_to_append_rel
func F_add_paths_to_append_rel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_recurse_pushdown_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_recurse_pushdown_safe
func F_recurse_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clauselist_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clauselist_selectivity
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64
//go:linkname F_clamp_row_est github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clamp_row_est
func F_clamp_row_est(m *base.Module, l0 float64) float64
//go:linkname F_cost_qual_eval_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_qual_eval_walker
func F_cost_qual_eval_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cost_subqueryscan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cost_subqueryscan
func F_cost_subqueryscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_cost_incremental_sort github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cost_incremental_sort
func F_cost_incremental_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 float64, l7 float64, l8 int32, l9 int32, l10 float64)
//go:linkname F_cost_tuplesort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cost_tuplesort
func F_cost_tuplesort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 float64, l5 int32, l6 float64)
//go:linkname F_initial_cost_nestloop github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initial_cost_nestloop
func F_initial_cost_nestloop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_cost_subplan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cost_subplan
func F_cost_subplan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_set_joinrel_size_estimates github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_joinrel_size_estimates
func F_set_joinrel_size_estimates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_canonicalize_ec_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_ec_expression
func F_canonicalize_ec_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ec_add_clause_to_derives_hash github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ec_add_clause_to_derives_hash
func F_ec_add_clause_to_derives_hash(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ec_clear_derived_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ec_clear_derived_clauses
func F_ec_clear_derived_clauses(m *base.Module, l0 int32)
//go:linkname F_get_eclass_for_sort_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_eclass_for_sort_expr
func F_get_eclass_for_sort_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_find_ec_member_matching_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_ec_member_matching_expr
func F_find_ec_member_matching_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_generate_join_implied_equalities github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_generate_join_implied_equalities
func F_generate_join_implied_equalities(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_create_join_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_join_clause
func F_create_join_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_exprs_known_equal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_exprs_known_equal
func F_exprs_known_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_build_index_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_build_index_paths
func F_build_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_match_index_to_operand github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_match_index_to_operand
func F_match_index_to_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_match_boolean_index_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_match_boolean_index_clause
func F_match_boolean_index_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_join_index_paths github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_join_index_paths
func F_get_join_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_get_index_clause_from_support github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_clause_from_support
func F_get_index_clause_from_support(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_add_paths_to_joinrel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_add_paths_to_joinrel
func F_add_paths_to_joinrel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_join_is_legal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_join_is_legal
func F_join_is_legal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_make_canonical_pathkey github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_canonical_pathkey
func F_make_canonical_pathkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_compare_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_compare_pathkeys
func F_compare_pathkeys(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_pathkeys_for_sortclauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_pathkeys_for_sortclauses
func F_make_pathkeys_for_sortclauses(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_IsBinaryTidClause github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_IsBinaryTidClause
func F_IsBinaryTidClause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_plan
func F_create_plan(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_scan_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_scan_plan
func F_create_scan_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_nestloop_params_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replace_nestloop_params_mutator
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_order_qual_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_order_qual_clauses
func F_order_qual_clauses(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_switched_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_switched_clauses
func F_get_switched_clauses(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_prepare_sort_from_pathkeys github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_prepare_sort_from_pathkeys
func F_prepare_sort_from_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_inject_projection_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_inject_projection_plan
func F_inject_projection_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mark_async_capable_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mark_async_capable_plan
func F_mark_async_capable_plan(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_gating_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_gating_plan
func F_create_gating_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_materialize_finished_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_materialize_finished_plan
func F_materialize_finished_plan(m *base.Module, l0 int32) int32
//go:linkname F_fix_indexqual_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fix_indexqual_clause
func F_fix_indexqual_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_subquery_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subquery_planner
func F_subquery_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32) int32
//go:linkname F_expression_planner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expression_planner
func F_expression_planner(m *base.Module, l0 int32) int32
//go:linkname F_set_plan_references github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_plan_references
func F_set_plan_references(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_upper_expr_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fix_upper_expr_mutator
func F_fix_upper_expr_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fix_expr_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fix_expr_common
func F_fix_expr_common(m *base.Module, l0 int32, l1 int32)
//go:linkname F_simplify_EXISTS_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_simplify_EXISTS_query
func F_simplify_EXISTS_query(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_build_subplan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_build_subplan
func F_build_subplan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_SS_finalize_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SS_finalize_plan
func F_SS_finalize_plan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_finalize_primnode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_finalize_primnode
func F_finalize_primnode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_finalize_agg_primnode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_finalize_agg_primnode
func F_finalize_agg_primnode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_preprocess_aggrefs_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_preprocess_aggrefs_walker
func F_preprocess_aggrefs_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replace_empty_jointree github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replace_empty_jointree
func F_replace_empty_jointree(m *base.Module, l0 int32)
//go:linkname F_get_relids_in_jointree github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_relids_in_jointree
func F_get_relids_in_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_remove_result_refs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_remove_result_refs
func F_remove_result_refs(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_jointree_contains_lateral_outer_refs github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jointree_contains_lateral_outer_refs
func F_jointree_contains_lateral_outer_refs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_replace_vars_in_jointree github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replace_vars_in_jointree
func F_replace_vars_in_jointree(m *base.Module, l0 int32, l1 int32)
//go:linkname F_canonicalize_qual github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_canonicalize_qual
func F_canonicalize_qual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_adjust_appendrel_attrs_mutator github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_adjust_appendrel_attrs_mutator
func F_adjust_appendrel_attrs_mutator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_adjust_appendrel_attrs_multilevel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_adjust_appendrel_attrs_multilevel
func F_adjust_appendrel_attrs_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_appinfos_by_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_appinfos_by_relids
func F_find_appinfos_by_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_adjust_child_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_adjust_child_relids
func F_adjust_child_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_contain_subplans github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_contain_subplans
func F_contain_subplans(m *base.Module, l0 int32) int32
//go:linkname F_contain_mutable_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_mutable_functions
func F_contain_mutable_functions(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_contain_volatile_functions
func F_contain_volatile_functions(m *base.Module, l0 int32) int32
//go:linkname F_contain_volatile_functions_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_volatile_functions_walker
func F_contain_volatile_functions_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_max_parallel_hazard_walker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_max_parallel_hazard_walker
func F_max_parallel_hazard_walker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_eval_const_expressions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_eval_const_expressions
func F_eval_const_expressions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_SAOP_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_SAOP_expr
func F_make_SAOP_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_have_relevant_joinclause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_have_relevant_joinclause
func F_have_relevant_joinclause(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_assign_special_exec_param github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_assign_special_exec_param
func F_assign_special_exec_param(m *base.Module, l0 int32) int32
//go:linkname F_compare_fractional_path_costs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_compare_fractional_path_costs
func F_compare_fractional_path_costs(m *base.Module, l0 int32, l1 int32, l2 float64) int32
//go:linkname F_set_cheapest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_cheapest
func F_set_cheapest(m *base.Module, l0 int32)
//go:linkname F_add_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_path
func F_add_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_add_partial_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_partial_path
func F_add_partial_path(m *base.Module, l0 int32, l1 int32)
//go:linkname F_create_tidscan_path github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_create_tidscan_path
func F_create_tidscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_create_append_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_append_path
func F_create_append_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 float64) int32
//go:linkname F_create_unique_path github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_create_unique_path
func F_create_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_calc_non_nestloop_required_outer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_calc_non_nestloop_required_outer
func F_calc_non_nestloop_required_outer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_nestloop_path github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_nestloop_path
func F_create_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_create_hashjoin_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_create_hashjoin_path
func F_create_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32
//go:linkname F_reparameterize_path_by_child github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_reparameterize_path_by_child
func F_reparameterize_path_by_child(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_path_is_reparameterizable_by_child github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_path_is_reparameterizable_by_child
func F_path_is_reparameterizable_by_child(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_placeholder_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_placeholder_expr
func F_make_placeholder_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_placeholder_info github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_placeholder_info
func F_find_placeholder_info(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rel_data_width github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rel_data_width
func F_get_rel_data_width(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_predicate_implied_by github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_predicate_implied_by
func F_predicate_implied_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_base_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_base_rel
func F_find_base_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_find_join_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_join_rel
func F_find_join_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_match_expr_to_partition_keys github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_match_expr_to_partition_keys
func F_match_expr_to_partition_keys(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fetch_upper_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_upper_rel
func F_fetch_upper_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_find_childrel_parents github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_find_childrel_parents
func F_find_childrel_parents(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_restrictinfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_restrictinfo
func F_make_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_extract_actual_clauses github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_extract_actual_clauses
func F_extract_actual_clauses(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_join_clause_is_movable_to github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_join_clause_is_movable_to
func F_join_clause_is_movable_to(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tlist_member github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tlist_member
func F_tlist_member(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tlist_same_exprs github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tlist_same_exprs
func F_tlist_same_exprs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupref_tle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_sortgroupref_tle
func F_get_sortgroupref_tle(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_sortgroupclause_tle github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_sortgroupclause_tle
func F_get_sortgroupclause_tle(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_extract_grouping_ops github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_extract_grouping_ops
func F_extract_grouping_ops(m *base.Module, l0 int32) int32
//go:linkname F_extract_grouping_collations github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_extract_grouping_collations
func F_extract_grouping_collations(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_extract_grouping_cols github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_extract_grouping_cols
func F_extract_grouping_cols(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_create_empty_pathtarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_create_empty_pathtarget
func F_create_empty_pathtarget(m *base.Module) int32
//go:linkname F_apply_pathtarget_labeling_to_tlist github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_apply_pathtarget_labeling_to_tlist
func F_apply_pathtarget_labeling_to_tlist(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pull_varnos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pull_varnos
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_contain_var_clause
func F_contain_var_clause(m *base.Module, l0 int32) int32
//go:linkname F_locate_var_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_locate_var_of_level
func F_locate_var_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pull_var_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pull_var_clause
func F_pull_var_clause(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flatten_group_exprs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_flatten_group_exprs
func F_flatten_group_exprs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_qual_from_partbound github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_qual_from_partbound
func F_get_qual_from_partbound(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_partition_op_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_partition_op_expr
func F_make_partition_op_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_range_nulltest github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_range_nulltest
func F_get_range_nulltest(m *base.Module, l0 int32) int32
//go:linkname F_generate_matching_part_pairs github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_generate_matching_part_pairs
func F_generate_matching_part_pairs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_build_merged_partition_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_build_merged_partition_bounds
func F_build_merged_partition_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_RelationGetPartitionDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetPartitionDesc
func F_RelationGetPartitionDesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DestroyPartitionDirectory github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DestroyPartitionDirectory
func F_DestroyPartitionDirectory(m *base.Module, l0 int32)
//go:linkname F_gen_partprune_steps_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_gen_partprune_steps_internal
func F_gen_partprune_steps_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rebuild_database_list github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_rebuild_database_list
func F_rebuild_database_list(m *base.Module, l0 int32)
//go:linkname F_BackgroundWorkerUnblockSignals github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_BackgroundWorkerUnblockSignals
func F_BackgroundWorkerUnblockSignals(m *base.Module)
//go:linkname F_GetBackgroundWorkerPid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetBackgroundWorkerPid
func F_GetBackgroundWorkerPid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AbsorbSyncRequests github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AbsorbSyncRequests
func F_AbsorbSyncRequests(m *base.Module)
//go:linkname F_RequestCheckpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RequestCheckpoint
func F_RequestCheckpoint(m *base.Module, l0 int32)
//go:linkname F_GetLatestLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetLatestLSN
func F_GetLatestLSN(m *base.Module, l0 int32) int64
//go:linkname F_WaitForWalSummarization github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WaitForWalSummarization
func F_WaitForWalSummarization(m *base.Module, l0 int64)
//go:linkname F_pg_regcomp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_regcomp
func F_pg_regcomp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_subcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subcolor
func F_subcolor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_okcolors github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_okcolors
func F_okcolors(m *base.Module, l0 int32, l1 int32)
//go:linkname F_newstate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_newstate
func F_newstate(m *base.Module, l0 int32) int32
//go:linkname F_createarc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_createarc
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_newcolor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_newcolor
func F_newcolor(m *base.Module, l0 int32) int32
//go:linkname F_cparc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cparc
func F_cparc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_colorcomplement github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_colorcomplement
func F_colorcomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_sortouts github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_sortouts
func F_sortouts(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_ctype_get_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ctype_get_cache
func F_pg_ctype_get_cache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_newhicolorrow github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_newhicolorrow
func F_newhicolorrow(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_subcoloronerow github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_subcoloronerow
func F_subcoloronerow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_clonesuccessorstates github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_clonesuccessorstates
func F_clonesuccessorstates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_pg_regexec github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_regexec
func F_pg_regexec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_logicalrep_worker_wakeup_ptr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_worker_wakeup_ptr
func F_logicalrep_worker_wakeup_ptr(m *base.Module, l0 int32)
//go:linkname F_logicalrep_worker_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_logicalrep_worker_attach
func F_logicalrep_worker_attach(m *base.Module, l0 int32)
//go:linkname F_logicalrep_launcher_attach_dshmem github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_launcher_attach_dshmem
func F_logicalrep_launcher_attach_dshmem(m *base.Module)
//go:linkname F_CheckLogicalDecodingRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckLogicalDecodingRequirements
func F_CheckLogicalDecodingRequirements(m *base.Module)
//go:linkname F_StartupDecodingContext github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StartupDecodingContext
func F_StartupDecodingContext(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_LogicalConfirmReceivedLocation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LogicalConfirmReceivedLocation
func F_LogicalConfirmReceivedLocation(m *base.Module, l0 int64)
//go:linkname F_LogicalSlotAdvanceAndCheckSnapState github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalSlotAdvanceAndCheckSnapState
func F_LogicalSlotAdvanceAndCheckSnapState(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_pg_logical_slot_get_changes_guts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_logical_slot_get_changes_guts
func F_pg_logical_slot_get_changes_guts(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replorigin_by_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replorigin_by_name
func F_replorigin_by_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replorigin_create github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_create
func F_replorigin_create(m *base.Module, l0 int32) int32
//go:linkname F_replorigin_advance github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replorigin_advance
func F_replorigin_advance(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32)
//go:linkname F_replorigin_session_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_session_setup
func F_replorigin_session_setup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_replorigin_session_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_session_reset
func F_replorigin_session_reset(m *base.Module)
//go:linkname F_replorigin_session_get_progress github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replorigin_session_get_progress
func F_replorigin_session_get_progress(m *base.Module) int64
//go:linkname F_replorigin_check_prerequisites github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replorigin_check_prerequisites
func F_replorigin_check_prerequisites(m *base.Module)
//go:linkname F_logicalrep_relmap_update github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_logicalrep_relmap_update
func F_logicalrep_relmap_update(m *base.Module, l0 int32)
//go:linkname F_logicalrep_rel_open github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_logicalrep_rel_open
func F_logicalrep_rel_open(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_logicalrep_rel_close github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_logicalrep_rel_close
func F_logicalrep_rel_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferFreeChange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferFreeChange
func F_ReorderBufferFreeChange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferTXNByXid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferTXNByXid
func F_ReorderBufferTXNByXid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_ReorderBufferStreamTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferStreamTXN
func F_ReorderBufferStreamTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferTransferSnapToParent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReorderBufferTransferSnapToParent
func F_ReorderBufferTransferSnapToParent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferTruncateTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReorderBufferTruncateTXN
func F_ReorderBufferTruncateTXN(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReorderBufferCleanupTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReorderBufferCleanupTXN
func F_ReorderBufferCleanupTXN(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ReorderBufferProcessTXN github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReorderBufferProcessTXN
func F_ReorderBufferProcessTXN(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32)
//go:linkname F_finish_sync_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_finish_sync_worker
func F_finish_sync_worker(m *base.Module)
//go:linkname F_ReplicationOriginNameForLogicalRep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationOriginNameForLogicalRep
func F_ReplicationOriginNameForLogicalRep(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_maybe_reread_subscription github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_maybe_reread_subscription
func F_maybe_reread_subscription(m *base.Module)
//go:linkname F_apply_dispatch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_apply_dispatch
func F_apply_dispatch(m *base.Module, l0 int32)
//go:linkname F_stream_cleanup_files github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_stream_cleanup_files
func F_stream_cleanup_files(m *base.Module, l0 int32, l1 int32)
//go:linkname F_start_apply github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_start_apply
func F_start_apply(m *base.Module, l0 int64)
//go:linkname F_DisableSubscriptionAndExit github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DisableSubscriptionAndExit
func F_DisableSubscriptionAndExit(m *base.Module)
//go:linkname F_InitializeLogRepWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitializeLogRepWorker
func F_InitializeLogRepWorker(m *base.Module)
//go:linkname F_SetupApplyOrSyncWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SetupApplyOrSyncWorker
func F_SetupApplyOrSyncWorker(m *base.Module, l0 int32)
//go:linkname F_set_apply_error_context_origin github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_apply_error_context_origin
func F_set_apply_error_context_origin(m *base.Module, l0 int32)
//go:linkname F_replication_yyensure_buffer_stack github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replication_yyensure_buffer_stack
func F_replication_yyensure_buffer_stack(m *base.Module, l0 int32)
//go:linkname F_replication_yy_create_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_replication_yy_create_buffer
func F_replication_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replication_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_replication_yyerror
func F_replication_yyerror(m *base.Module, l0 int32)
//go:linkname F_yy_fatal_error_3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_yy_fatal_error_3
func F_yy_fatal_error_3(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ReplicationSlotRelease
func F_ReplicationSlotRelease(m *base.Module)
//go:linkname F_ReplicationSlotCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReplicationSlotCreate
func F_ReplicationSlotCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SaveSlotToPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SaveSlotToPath
func F_SaveSlotToPath(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReplicationSlotDropPtr github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotDropPtr
func F_ReplicationSlotDropPtr(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotsComputeRequiredXmin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotsComputeRequiredXmin
func F_ReplicationSlotsComputeRequiredXmin(m *base.Module, l0 int32)
//go:linkname F_ReplicationSlotsComputeRequiredLSN github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationSlotsComputeRequiredLSN
func F_ReplicationSlotsComputeRequiredLSN(m *base.Module)
//go:linkname F_ReplicationSlotMarkDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReplicationSlotMarkDirty
func F_ReplicationSlotMarkDirty(m *base.Module)
//go:linkname F_ReplicationSlotSave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReplicationSlotSave
func F_ReplicationSlotSave(m *base.Module)
//go:linkname F_CheckSlotRequirements github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckSlotRequirements
func F_CheckSlotRequirements(m *base.Module)
//go:linkname F_CheckSlotPermissions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CheckSlotPermissions
func F_CheckSlotPermissions(m *base.Module)
//go:linkname F_ReplicationSlotReserveWal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReplicationSlotReserveWal
func F_ReplicationSlotReserveWal(m *base.Module)
//go:linkname F_InvalidateObsoleteReplicationSlots github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_InvalidateObsoleteReplicationSlots
func F_InvalidateObsoleteReplicationSlots(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_SyncRepGetCandidateStandbys github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SyncRepGetCandidateStandbys
func F_SyncRepGetCandidateStandbys(m *base.Module, l0 int32) int32
//go:linkname F_WalRcvRunning github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WalRcvRunning
func F_WalRcvRunning(m *base.Module) int32
//go:linkname F_WalSndWakeup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_WalSndWakeup
func F_WalSndWakeup(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WalSndKeepalive github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_WalSndKeepalive
func F_WalSndKeepalive(m *base.Module, l0 int32, l1 int64)
//go:linkname F_PhysicalReplicationSlotNewXmin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PhysicalReplicationSlotNewXmin
func F_PhysicalReplicationSlotNewXmin(m *base.Module, l0 int32, l1 int32)
//go:linkname F_DefineQueryRewrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DefineQueryRewrite
func F_DefineQueryRewrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_build_column_default github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_build_column_default
func F_build_column_default(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RewriteQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RewriteQuery
func F_RewriteQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fireRIRrules github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fireRIRrules
func F_fireRIRrules(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_contain_aggs_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_contain_aggs_of_level
func F_contain_aggs_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_locate_agg_of_level github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_locate_agg_of_level
func F_locate_agg_of_level(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_locate_windowfunc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_locate_windowfunc
func F_locate_windowfunc(m *base.Module, l0 int32) int32
//go:linkname F_checkExprHasSubLink github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_checkExprHasSubLink
func F_checkExprHasSubLink(m *base.Module, l0 int32) int32
//go:linkname F_CombineRangeTables github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CombineRangeTables
func F_CombineRangeTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_OffsetVarNodes github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_OffsetVarNodes
func F_OffsetVarNodes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ChangeVarNodesExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ChangeVarNodesExtended
func F_ChangeVarNodesExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_adjust_relid_set github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_adjust_relid_set
func F_adjust_relid_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rangeTableEntry_used github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_rangeTableEntry_used
func F_rangeTableEntry_used(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getInsertSelectQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getInsertSelectQuery
func F_getInsertSelectQuery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_add_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_add_nulling_relids
func F_add_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_remove_nulling_relids github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_remove_nulling_relids
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_rte_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_replace_rte_variables
func F_replace_rte_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_map_variable_attnos github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_variable_attnos
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_generate_dependencies_recurse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_generate_dependencies_recurse
func F_generate_dependencies_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fetch_statentries_for_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fetch_statentries_for_relation
func F_fetch_statentries_for_relation(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_multi_sort_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multi_sort_init
func F_multi_sort_init(m *base.Module, l0 int32) int32
//go:linkname F_multi_sort_add_dimension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_multi_sort_add_dimension
func F_multi_sort_add_dimension(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_multi_sort_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_multi_sort_compare
func F_multi_sort_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_build_sorted_items github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_build_sorted_items
func F_build_sorted_items(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_stats_check_required_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_stats_check_required_arg
func F_stats_check_required_arg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgaio_io_reclaim github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_io_reclaim
func F_pgaio_io_reclaim(m *base.Module, l0 int32)
//go:linkname F_pgaio_wref_wait github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_wref_wait
func F_pgaio_wref_wait(m *base.Module, l0 int32)
//go:linkname F_pgaio_error_cleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgaio_error_cleanup
func F_pgaio_error_cleanup(m *base.Module)
//go:linkname F_pgaio_closing_fd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgaio_closing_fd
func F_pgaio_closing_fd(m *base.Module, l0 int32)
//go:linkname F_pgaio_result_report github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgaio_result_report
func F_pgaio_result_report(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_read_stream_begin_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_read_stream_begin_relation
func F_read_stream_begin_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_read_stream_next_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read_stream_next_buffer
func F_read_stream_next_buffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_stream_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_read_stream_reset
func F_read_stream_reset(m *base.Module, l0 int32)
//go:linkname F_read_stream_end github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_stream_end
func F_read_stream_end(m *base.Module, l0 int32)
//go:linkname F_BufTableLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufTableLookup
func F_BufTableLookup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufTableInsert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufTableInsert
func F_BufTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_UnpinBufferNoOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnpinBufferNoOwner
func F_UnpinBufferNoOwner(m *base.Module, l0 int32)
//go:linkname F_PrefetchBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PrefetchBuffer
func F_PrefetchBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ReservePrivateRefCountEntry github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReservePrivateRefCountEntry
func F_ReservePrivateRefCountEntry(m *base.Module)
//go:linkname F_PinBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PinBuffer
func F_PinBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ReadBuffer
func F_ReadBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadBufferExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadBufferExtended
func F_ReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ExtendBufferedRel github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ExtendBufferedRel
func F_ExtendBufferedRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetVictimBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetVictimBuffer
func F_GetVictimBuffer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReadBufferWithoutRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReadBufferWithoutRelcache
func F_ReadBufferWithoutRelcache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseBuffer
func F_ReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_AsyncReadBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AsyncReadBuffers
func F_AsyncReadBuffers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MarkBufferDirty github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MarkBufferDirty
func F_MarkBufferDirty(m *base.Module, l0 int32)
//go:linkname F_RelationGetNumberOfBlocksInFork github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetNumberOfBlocksInFork
func F_RelationGetNumberOfBlocksInFork(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufferGetLSNAtomic github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufferGetLSNAtomic
func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64
//go:linkname F_FindAndDropRelationBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FindAndDropRelationBuffers
func F_FindAndDropRelationBuffers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_InvalidateBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_InvalidateBuffer
func F_InvalidateBuffer(m *base.Module, l0 int32)
//go:linkname F_FlushRelationBuffers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FlushRelationBuffers
func F_FlushRelationBuffers(m *base.Module, l0 int32)
//go:linkname F_RelationCopyStorageUsingBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationCopyStorageUsingBuffer
func F_RelationCopyStorageUsingBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_UnlockReleaseBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UnlockReleaseBuffer
func F_UnlockReleaseBuffer(m *base.Module, l0 int32)
//go:linkname F_LockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockBuffer
func F_LockBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_IncrBufferRefCount github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_IncrBufferRefCount
func F_IncrBufferRefCount(m *base.Module, l0 int32)
//go:linkname F_MarkBufferDirtyHint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkBufferDirtyHint
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockBuffer
func F_ConditionalLockBuffer(m *base.Module, l0 int32) int32
//go:linkname F_CheckBufferIsPinnedOnce github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckBufferIsPinnedOnce
func F_CheckBufferIsPinnedOnce(m *base.Module, l0 int32)
//go:linkname F_LockBufferForCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockBufferForCleanup
func F_LockBufferForCleanup(m *base.Module, l0 int32)
//go:linkname F_ConditionalLockBufferForCleanup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionalLockBufferForCleanup
func F_ConditionalLockBufferForCleanup(m *base.Module, l0 int32) int32
//go:linkname F_IsBufferCleanupOK github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_IsBufferCleanupOK
func F_IsBufferCleanupOK(m *base.Module, l0 int32) int32
//go:linkname F_StrategyFreeBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StrategyFreeBuffer
func F_StrategyFreeBuffer(m *base.Module, l0 int32)
//go:linkname F_GetAccessStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetAccessStrategy
func F_GetAccessStrategy(m *base.Module, l0 int32) int32
//go:linkname F_IOContextForStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_IOContextForStrategy
func F_IOContextForStrategy(m *base.Module, l0 int32) int32
//go:linkname F_LocalBufferAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LocalBufferAlloc
func F_LocalBufferAlloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_InvalidateLocalBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateLocalBuffer
func F_InvalidateLocalBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BufFileCreateTemp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileCreateTemp
func F_BufFileCreateTemp(m *base.Module, l0 int32) int32
//go:linkname F_BufFileCreateFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileCreateFileSet
func F_BufFileCreateFileSet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BufFileOpenFileSet github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BufFileOpenFileSet
func F_BufFileOpenFileSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_BufFileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileClose
func F_BufFileClose(m *base.Module, l0 int32)
//go:linkname F_BufFileReadExact github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BufFileReadExact
func F_BufFileReadExact(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileReadMaybeEOF github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileReadMaybeEOF
func F_BufFileReadMaybeEOF(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_BufFileWrite github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileWrite
func F_BufFileWrite(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BufFileSeek github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BufFileSeek
func F_BufFileSeek(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_BufFileSeekBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BufFileSeekBlock
func F_BufFileSeekBlock(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_BufFileSize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BufFileSize
func F_BufFileSize(m *base.Module, l0 int32) int64
//go:linkname F_copy_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_copy_file
func F_copy_file(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_file_exists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_file_exists
func F_pg_file_exists(m *base.Module, l0 int32) int32
//go:linkname F_fsync_fname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fsync_fname
func F_fsync_fname(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fsync_fname_ext github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fsync_fname_ext
func F_fsync_fname_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_CloseTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CloseTransientFile
func F_CloseTransientFile(m *base.Module, l0 int32) int32
//go:linkname F_OpenTransientFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OpenTransientFile
func F_OpenTransientFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileClose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileClose
func F_FileClose(m *base.Module, l0 int32)
//go:linkname F_PathNameOpenFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PathNameOpenFile
func F_PathNameOpenFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AllocateDir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_AllocateDir
func F_AllocateDir(m *base.Module, l0 int32) int32
//go:linkname F_FreeDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FreeDir
func F_FreeDir(m *base.Module, l0 int32)
//go:linkname F_TempTablespacePath github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TempTablespacePath
func F_TempTablespacePath(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PathNameCreateTemporaryFile github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PathNameCreateTemporaryFile
func F_PathNameCreateTemporaryFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileAccess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileAccess
func F_FileAccess(m *base.Module, l0 int32) int32
//go:linkname F_FileReadV github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FileReadV
func F_FileReadV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_FileWriteV github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FileWriteV
func F_FileWriteV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_FileSync github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FileSync
func F_FileSync(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FileSize github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_FileSize
func F_FileSize(m *base.Module, l0 int32) int64
//go:linkname F_AllocateFile github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AllocateFile
func F_AllocateFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_reserveAllocatedDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_reserveAllocatedDesc
func F_reserveAllocatedDesc(m *base.Module) int32
//go:linkname F_OpenPipeStream github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OpenPipeStream
func F_OpenPipeStream(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FreeFile github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FreeFile
func F_FreeFile(m *base.Module, l0 int32) int32
//go:linkname F_ReadDir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReadDir
func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_AtEOXact_Files github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_Files
func F_AtEOXact_Files(m *base.Module, l0 int32)
//go:linkname F_RemovePgTempFilesInDir github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RemovePgTempFilesInDir
func F_RemovePgTempFilesInDir(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_RemovePgTempRelationFiles github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RemovePgTempRelationFiles
func F_RemovePgTempRelationFiles(m *base.Module, l0 int32)
//go:linkname F_fsm_readbuf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fsm_readbuf
func F_fsm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RecordPageWithFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RecordPageWithFreeSpace
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_FreeSpaceMapVacuum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeSpaceMapVacuum
func F_FreeSpaceMapVacuum(m *base.Module, l0 int32)
//go:linkname F_FreeSpaceMapVacuumRange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreeSpaceMapVacuumRange
func F_FreeSpaceMapVacuumRange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetFreeIndexPage
func F_GetFreeIndexPage(m *base.Module, l0 int32) int32
//go:linkname F_RecordFreeIndexPage github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RecordFreeIndexPage
func F_RecordFreeIndexPage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BarrierArriveAndDetach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BarrierArriveAndDetach
func F_BarrierArriveAndDetach(m *base.Module, l0 int32) int32
//go:linkname F_BarrierAttach github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_BarrierAttach
func F_BarrierAttach(m *base.Module, l0 int32) int32
//go:linkname F_dsm_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_attach
func F_dsm_attach(m *base.Module, l0 int32) int32
//go:linkname F_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsm_detach
func F_dsm_detach(m *base.Module, l0 int32)
//go:linkname F_on_dsm_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_on_dsm_detach
func F_on_dsm_detach(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dsm_impl_op github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_dsm_impl_op
func F_dsm_impl_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_proc_exit
func F_proc_exit(m *base.Module, l0 int32)
//go:linkname F_on_proc_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_on_proc_exit
func F_on_proc_exit(m *base.Module, l0 int32)
//go:linkname F_before_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_before_shmem_exit
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_on_shmem_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_on_shmem_exit
func F_on_shmem_exit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_WaitLatch github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_WaitLatch
func F_WaitLatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SendPostmasterSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SendPostmasterSignal
func F_SendPostmasterSignal(m *base.Module, l0 int32)
//go:linkname F_PostmasterIsAliveInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PostmasterIsAliveInternal
func F_PostmasterIsAliveInternal(m *base.Module) int32
//go:linkname F_KnownAssignedXidsAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_KnownAssignedXidsAdd
func F_KnownAssignedXidsAdd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_KnownAssignedXidsDisplay github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_KnownAssignedXidsDisplay
func F_KnownAssignedXidsDisplay(m *base.Module, l0 int32)
//go:linkname F_TransactionIdIsInProgress github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_TransactionIdIsInProgress
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32
//go:linkname F_GlobalVisHorizonKindForRel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GlobalVisHorizonKindForRel
func F_GlobalVisHorizonKindForRel(m *base.Module, l0 int32) int32
//go:linkname F_GetOldestTransactionIdConsideredRunning github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetOldestTransactionIdConsideredRunning
func F_GetOldestTransactionIdConsideredRunning(m *base.Module) int32
//go:linkname F_GetVirtualXIDsDelayingChkpt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetVirtualXIDsDelayingChkpt
func F_GetVirtualXIDsDelayingChkpt(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_HaveVirtualXIDsDelayingChkpt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_HaveVirtualXIDsDelayingChkpt
func F_HaveVirtualXIDsDelayingChkpt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_BackendPidGetProc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BackendPidGetProc
func F_BackendPidGetProc(m *base.Module, l0 int32) int32
//go:linkname F_GlobalVisCheckRemovableFullXid github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GlobalVisCheckRemovableFullXid
func F_GlobalVisCheckRemovableFullXid(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_GlobalVisCheckRemovableXid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GlobalVisCheckRemovableXid
func F_GlobalVisCheckRemovableXid(m *base.Module, l0 int32) int32
//go:linkname F_SendProcSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendProcSignal
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ProcessProcSignalBarrier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessProcSignalBarrier
func F_ProcessProcSignalBarrier(m *base.Module)
//go:linkname F_shm_mq_set_receiver github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_set_receiver
func F_shm_mq_set_receiver(m *base.Module, l0 int32, l1 int32)
//go:linkname F_shm_mq_set_sender github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_mq_set_sender
func F_shm_mq_set_sender(m *base.Module, l0 int32, l1 int32)
//go:linkname F_shm_mq_get_sender github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_get_sender
func F_shm_mq_get_sender(m *base.Module, l0 int32) int32
//go:linkname F_shm_mq_attach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_attach
func F_shm_mq_attach(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_mq_send github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_shm_mq_send
func F_shm_mq_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_shm_mq_wait_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_wait_internal
func F_shm_mq_wait_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_shm_mq_receive github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_shm_mq_receive
func F_shm_mq_receive(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_shm_mq_detach github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_mq_detach
func F_shm_mq_detach(m *base.Module, l0 int32)
//go:linkname F_shm_toc_allocate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_shm_toc_allocate
func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_shm_toc_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_shm_toc_insert
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_shm_toc_lookup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_shm_toc_lookup
func F_shm_toc_lookup(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_add_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_size
func F_add_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mul_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mul_size
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveSharedInvalidMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveSharedInvalidMessages
func F_ReceiveSharedInvalidMessages(m *base.Module)
//go:linkname F_SICleanupQueue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SICleanupQueue
func F_SICleanupQueue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_StandbyReleaseAllLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_StandbyReleaseAllLocks
func F_StandbyReleaseAllLocks(m *base.Module)
//go:linkname F_StandbyReleaseXidEntryLocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_StandbyReleaseXidEntryLocks
func F_StandbyReleaseXidEntryLocks(m *base.Module, l0 int32)
//go:linkname F_ResolveRecoveryConflictWithSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResolveRecoveryConflictWithSnapshot
func F_ResolveRecoveryConflictWithSnapshot(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResolveRecoveryConflictWithSnapshotFullXid github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ResolveRecoveryConflictWithSnapshotFullXid
func F_ResolveRecoveryConflictWithSnapshotFullXid(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_LogStandbySnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogStandbySnapshot
func F_LogStandbySnapshot(m *base.Module) int64
//go:linkname F_ModifyWaitEvent github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ModifyWaitEvent
func F_ModifyWaitEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_inv_seek github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_inv_seek
func F_inv_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_ConditionVariablePrepareToSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariablePrepareToSleep
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableSleep github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ConditionVariableSleep
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionVariableSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionVariableSignal
func F_ConditionVariableSignal(m *base.Module, l0 int32)
//go:linkname F_ConditionVariableBroadcast github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ConditionVariableBroadcast
func F_ConditionVariableBroadcast(m *base.Module, l0 int32)
//go:linkname F_LockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRelationOid
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ConditionalLockRelationOid
func F_ConditionalLockRelationOid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LockRelationId github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockRelationId
func F_LockRelationId(m *base.Module, l0 int32)
//go:linkname F_UnlockRelationId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UnlockRelationId
func F_UnlockRelationId(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationOid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationOid
func F_UnlockRelationOid(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ConditionalLockRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ConditionalLockRelation
func F_ConditionalLockRelation(m *base.Module, l0 int32) int32
//go:linkname F_UnlockRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelation
func F_UnlockRelation(m *base.Module, l0 int32)
//go:linkname F_LockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockRelationForExtension
func F_LockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_UnlockRelationForExtension github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockRelationForExtension
func F_UnlockRelationForExtension(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockTuple
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_UnlockTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_UnlockTuple
func F_UnlockTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockDatabaseObject github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LockDatabaseObject
func F_LockDatabaseObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockSharedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockSharedObject
func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_LockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockApplyTransactionForSession
func F_LockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_UnlockApplyTransactionForSession github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnlockApplyTransactionForSession
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_LockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LockAcquire
func F_LockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_LockAcquireExtended github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LockAcquireExtended
func F_LockAcquireExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_CleanUpLock github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CleanUpLock
func F_CleanUpLock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_LockRefindAndRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LockRefindAndRelease
func F_LockRefindAndRelease(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_LWLockAcquire github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockAcquire
func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_LWLockQueueSelf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_LWLockQueueSelf
func F_LWLockQueueSelf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LWLockDequeueSelf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LWLockDequeueSelf
func F_LWLockDequeueSelf(m *base.Module, l0 int32)
//go:linkname F_LWLockUpdateVar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LWLockUpdateVar
func F_LWLockUpdateVar(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_LWLockRelease github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LWLockRelease
func F_LWLockRelease(m *base.Module, l0 int32)
//go:linkname F_LWLockReleaseAll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_LWLockReleaseAll
func F_LWLockReleaseAll(m *base.Module)
//go:linkname F_PredicateLockPage github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPage
func F_PredicateLockPage(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_TransferPredicateLocksToHeapRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_TransferPredicateLocksToHeapRelation
func F_TransferPredicateLocksToHeapRelation(m *base.Module, l0 int32)
//go:linkname F_PredicateLockPageSplit github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PredicateLockPageSplit
func F_PredicateLockPageSplit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PredicateLockPageCombine github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PredicateLockPageCombine
func F_PredicateLockPageCombine(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CheckForSerializableConflictIn github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CheckForSerializableConflictIn
func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ProcSendSignal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcSendSignal
func F_ProcSendSignal(m *base.Module, l0 int32)
//go:linkname F_s_lock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_s_lock
func F_s_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_s_lock_stuck github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_s_lock_stuck
func F_s_lock_stuck(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageInit
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageGetTempPageCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PageGetTempPageCopy
func F_PageGetTempPageCopy(m *base.Module, l0 int32) int32
//go:linkname F_PageGetTempPageCopySpecial github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageGetTempPageCopySpecial
func F_PageGetTempPageCopySpecial(m *base.Module, l0 int32) int32
//go:linkname F_PageGetFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageGetFreeSpace
func F_PageGetFreeSpace(m *base.Module, l0 int32) int32
//go:linkname F_PageGetExactFreeSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PageGetExactFreeSpace
func F_PageGetExactFreeSpace(m *base.Module, l0 int32) int32
//go:linkname F_PageIndexTupleDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleDelete
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PageIndexMultiDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PageIndexMultiDelete
func F_PageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_PageIndexTupleOverwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_PageIndexTupleOverwrite
func F_PageIndexTupleOverwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_checksum_page github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_checksum_page
func F_pg_checksum_page(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_start_rel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_start_rel
func F_smgr_bulk_start_rel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgr_bulk_flush github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_smgr_bulk_flush
func F_smgr_bulk_flush(m *base.Module, l0 int32)
//go:linkname F_smgr_bulk_write github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_smgr_bulk_write
func F_smgr_bulk_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_smgr_bulk_get_buf github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgr_bulk_get_buf
func F_smgr_bulk_get_buf(m *base.Module, l0 int32) int32
//go:linkname F_mdopenfork github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mdopenfork
func F_mdopenfork(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mdextend github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mdextend
func F_mdextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__mdfd_segpath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__mdfd_segpath
func F__mdfd_segpath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__mdfd_openseg github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F__mdfd_openseg
func F__mdfd_openseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_mdnblocks github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mdnblocks
func F_mdnblocks(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgropen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgropen
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrclose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrclose
func F_smgrclose(m *base.Module, l0 int32)
//go:linkname F_smgrexists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrexists
func F_smgrexists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrcreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrcreate
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_smgrextend github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_smgrextend
func F_smgrextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_smgrprefetch github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_smgrprefetch
func F_smgrprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_smgrnblocks github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_smgrnblocks
func F_smgrnblocks(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_smgrregistersync github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_smgrregistersync
func F_smgrregistersync(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RegisterSyncRequest github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RegisterSyncRequest
func F_RegisterSyncRequest(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_CreateDestReceiver github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateDestReceiver
func F_CreateDestReceiver(m *base.Module, l0 int32) int32
//go:linkname F_ProcessInterrupts github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessInterrupts
func F_ProcessInterrupts(m *base.Module)
//go:linkname F_pg_parse_query github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_parse_query
func F_pg_parse_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_rewrite_query github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rewrite_query
func F_pg_rewrite_query(m *base.Module, l0 int32) int32
//go:linkname F_pg_analyze_and_rewrite_withcb github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_analyze_and_rewrite_withcb
func F_pg_analyze_and_rewrite_withcb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_process_postgres_switches github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_process_postgres_switches
func F_process_postgres_switches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_PostgresMainLongJmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PostgresMainLongJmp
func F_PostgresMainLongJmp(m *base.Module)
//go:linkname F_PostgresMainLoopOnce github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PostgresMainLoopOnce
func F_PostgresMainLoopOnce(m *base.Module)
//go:linkname F_CreateQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CreateQueryDesc
func F_CreateQueryDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_FreeQueryDesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FreeQueryDesc
func F_FreeQueryDesc(m *base.Module, l0 int32)
//go:linkname F_ChoosePortalStrategy github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ChoosePortalStrategy
func F_ChoosePortalStrategy(m *base.Module, l0 int32) int32
//go:linkname F_FetchStatementTargetList github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_FetchStatementTargetList
func F_FetchStatementTargetList(m *base.Module, l0 int32) int32
//go:linkname F_PortalRunMulti github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PortalRunMulti
func F_PortalRunMulti(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_PortalRunUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalRunUtility
func F_PortalRunUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_PortalRunFetch github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PortalRunFetch
func F_PortalRunFetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64
//go:linkname F_CommandIsReadOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CommandIsReadOnly
func F_CommandIsReadOnly(m *base.Module, l0 int32) int32
//go:linkname F_ProcessUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessUtility
func F_ProcessUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_standard_ProcessUtility github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_standard_ProcessUtility
func F_standard_ProcessUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_CreateCommandTag github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateCommandTag
func F_CreateCommandTag(m *base.Module, l0 int32) int32
//go:linkname F_UtilityTupleDescriptor github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_UtilityTupleDescriptor
func F_UtilityTupleDescriptor(m *base.Module, l0 int32) int32
//go:linkname F_addCompoundAffixFlagValue github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_addCompoundAffixFlagValue
func F_addCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_getNextFlagFromString github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_getNextFlagFromString
func F_getNextFlagFromString(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setCompoundAffixFlagValue github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_setCompoundAffixFlagValue
func F_setCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_mkSPNode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_mkSPNode
func F_mkSPNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_mkANode github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mkANode
func F_mkANode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_mkVoidAffix github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mkVoidAffix
func F_mkVoidAffix(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_make_tsvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_make_tsvector
func F_make_tsvector(m *base.Module, l0 int32) int32
//go:linkname F_t_isalnum_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_t_isalnum_cstr
func F_t_isalnum_cstr(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_begin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline_begin
func F_tsearch_readline_begin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tsearch_readline github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsearch_readline
func F_tsearch_readline(m *base.Module, l0 int32) int32
//go:linkname F_tsearch_readline_end github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tsearch_readline_end
func F_tsearch_readline_end(m *base.Module, l0 int32)
//go:linkname F_parsetext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parsetext
func F_parsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tsquery_opr_selec github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tsquery_opr_selec
func F_tsquery_opr_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32) float64
//go:linkname F_get_tsearch_config_filename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_tsearch_config_filename
func F_get_tsearch_config_filename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_prs_setup_firstcall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_prs_setup_firstcall
func F_prs_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_prs_process_call github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_prs_process_call
func F_prs_process_call(m *base.Module, l0 int32) int32
//go:linkname F_TParserGet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TParserGet
func F_TParserGet(m *base.Module, l0 int32) int32
//go:linkname F_hlCover github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hlCover
func F_hlCover(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_pgstat_report_activity github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_report_activity
func F_pgstat_report_activity(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_get_beentry_by_proc_number github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_beentry_by_proc_number
func F_pgstat_get_beentry_by_proc_number(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_report_stat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_report_stat
func F_pgstat_report_stat(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_reset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_reset
func F_pgstat_reset(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_pgstat_fetch_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_entry
func F_pgstat_fetch_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_pgstat_snapshot_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_snapshot_fixed
func F_pgstat_snapshot_fixed(m *base.Module, l0 int32)
//go:linkname F_pgstat_count_backend_io_op github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_count_backend_io_op
func F_pgstat_count_backend_io_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64)
//go:linkname F_pgstat_flush_backend github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_flush_backend
func F_pgstat_flush_backend(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_fetch_stat_checkpointer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_fetch_stat_checkpointer
func F_pgstat_fetch_stat_checkpointer(m *base.Module) int32
//go:linkname F_pgstat_report_connect github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_report_connect
func F_pgstat_report_connect(m *base.Module)
//go:linkname F_pgstat_init_function_usage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_init_function_usage
func F_pgstat_init_function_usage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_find_funcstat_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_find_funcstat_entry
func F_find_funcstat_entry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_flush_io github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_flush_io
func F_pgstat_flush_io(m *base.Module, l0 int32)
//go:linkname F_pgstat_assoc_relation github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_assoc_relation
func F_pgstat_assoc_relation(m *base.Module, l0 int32)
//go:linkname F_pgstat_report_analyze github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_report_analyze
func F_pgstat_report_analyze(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int64)
//go:linkname F_pgstat_count_heap_insert github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_count_heap_insert
func F_pgstat_count_heap_insert(m *base.Module, l0 int32, l1 int64)
//go:linkname F_pgstat_fetch_stat_tabentry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_fetch_stat_tabentry
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32
//go:linkname F_find_tabstat_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_find_tabstat_entry
func F_find_tabstat_entry(m *base.Module, l0 int32) int32
//go:linkname F_pgstat_release_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_release_entry_ref
func F_pgstat_release_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pgstat_get_entry_ref github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_get_entry_ref
func F_pgstat_get_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_pgstat_lock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_lock_entry
func F_pgstat_lock_entry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_unlock_entry github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgstat_unlock_entry
func F_pgstat_unlock_entry(m *base.Module, l0 int32)
//go:linkname F_pgstat_get_entry_ref_locked github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgstat_get_entry_ref_locked
func F_pgstat_get_entry_ref_locked(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_pgstat_drop_entry_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgstat_drop_entry_internal
func F_pgstat_drop_entry_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgstat_report_subscription_error github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgstat_report_subscription_error
func F_pgstat_report_subscription_error(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pgstat_get_xact_stack_level github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgstat_get_xact_stack_level
func F_pgstat_get_xact_stack_level(m *base.Module, l0 int32) int32
//go:linkname F_create_drop_transactional_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_create_drop_transactional_internal
func F_create_drop_transactional_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_aclupdate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_aclupdate
func F_aclupdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_aclnewowner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_aclnewowner
func F_aclnewowner(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_roles_is_member_of github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_roles_is_member_of
func F_roles_is_member_of(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_aclmembers github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_aclmembers
func F_aclmembers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_role_oid_or_public github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_role_oid_or_public
func F_get_role_oid_or_public(m *base.Module, l0 int32) int32
//go:linkname F_check_can_set_role github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_can_set_role
func F_check_can_set_role(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_rolespec_oid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_rolespec_oid
func F_get_rolespec_oid(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DatumGetExpandedArray github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DatumGetExpandedArray
func F_DatumGetExpandedArray(m *base.Module, l0 int32) int32
//go:linkname F_deconstruct_expanded_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_expanded_array
func F_deconstruct_expanded_array(m *base.Module, l0 int32)
//go:linkname F_mcelem_array_contain_overlap_selec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mcelem_array_contain_overlap_selec
func F_mcelem_array_contain_overlap_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_construct_empty_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_empty_array
func F_construct_empty_array(m *base.Module, l0 int32) int32
//go:linkname F_ArrayCastAndSet github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ArrayCastAndSet
func F_ArrayCastAndSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_array_get_element github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_get_element
func F_array_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_array_set_element github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_array_set_element
func F_array_set_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_construct_md_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_md_array
func F_construct_md_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_deconstruct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_construct_array github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_construct_array
func F_construct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_construct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_construct_array_builtin
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deconstruct_array_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_deconstruct_array_builtin
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_array_contains_nulls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_contains_nulls
func F_array_contains_nulls(m *base.Module, l0 int32) int32
//go:linkname F_array_contain_compare github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_contain_compare
func F_array_contain_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_array_create_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_create_iterator
func F_array_create_iterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_array_free_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_array_free_iterator
func F_array_free_iterator(m *base.Module, l0 int32)
//go:linkname F_initArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_initArrayResult
func F_initArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_initArrayResultWithSize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_initArrayResultWithSize
func F_initArrayResultWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_accumArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_accumArrayResult
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResult github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_makeArrayResult
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initArrayResultArr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initArrayResultArr
func F_initArrayResultArr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_accumArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accumArrayResultAny
func F_accumArrayResultAny(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_makeArrayResultAny github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_makeArrayResultAny
func F_makeArrayResultAny(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ArrayGetNItems github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ArrayGetNItems
func F_ArrayGetNItems(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_encode_to_ascii github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_encode_to_ascii
func F_encode_to_ascii(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_date2j github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_date2j
func F_date2j(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_j2date github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_j2date
func F_j2date(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_GetCurrentDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentDateTime
func F_GetCurrentDateTime(m *base.Module, l0 int32)
//go:linkname F_GetCurrentTimeUsec github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetCurrentTimeUsec
func F_GetCurrentTimeUsec(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ParseDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ParseDateTime
func F_ParseDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DecodeDateTime github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_DecodeDateTime
func F_DecodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DecodeNumber github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DecodeNumber
func F_DecodeNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_DecodeTimezoneAbbrev github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DecodeTimezoneAbbrev
func F_DecodeTimezoneAbbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DetermineTimeZoneOffset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DetermineTimeZoneOffset
func F_DetermineTimeZoneOffset(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_DetermineTimeZoneAbbrevOffsetTS github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DetermineTimeZoneAbbrevOffsetTS
func F_DetermineTimeZoneAbbrevOffsetTS(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DateTimeParseError github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DateTimeParseError
func F_DateTimeParseError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_DecodeTimezoneNameToTz github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DecodeTimezoneNameToTz
func F_DecodeTimezoneNameToTz(m *base.Module, l0 int32) int32
//go:linkname F_datumGetSize github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datumGetSize
func F_datumGetSize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_datumIsEqual github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_datumIsEqual
func F_datumIsEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_datumRestore github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_datumRestore
func F_datumRestore(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_db_dir_size github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_db_dir_size
func F_db_dir_size(m *base.Module, l0 int32) int64
//go:linkname F_calculate_relation_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_calculate_relation_size
func F_calculate_relation_size(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_calculate_table_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_calculate_table_size
func F_calculate_table_size(m *base.Module, l0 int32) int64
//go:linkname F_domain_state_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_domain_state_setup
func F_domain_state_setup(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_domain_check_input github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_domain_check_input
func F_domain_check_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errdatatype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdatatype
func F_errdatatype(m *base.Module, l0 int32)
//go:linkname F_domain_check github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_domain_check
func F_domain_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_domain_check_safe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_domain_check_safe
func F_domain_check_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_EOH_get_flat_size github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_EOH_get_flat_size
func F_EOH_get_flat_size(m *base.Module, l0 int32) int32
//go:linkname F_EOH_flatten_into github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EOH_flatten_into
func F_EOH_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_DeleteExpandedObject github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DeleteExpandedObject
func F_DeleteExpandedObject(m *base.Module, l0 int32)
//go:linkname F_make_expanded_record_from_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_make_expanded_record_from_tupdesc
func F_make_expanded_record_from_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_make_expanded_record_from_exprecord github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_expanded_record_from_exprecord
func F_make_expanded_record_from_exprecord(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_expanded_record_fetch_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_fetch_tupdesc
func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32
//go:linkname F_expanded_record_set_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_expanded_record_set_tuple
func F_expanded_record_set_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_expanded_record_get_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_get_tuple
func F_expanded_record_get_tuple(m *base.Module, l0 int32) int32
//go:linkname F_expanded_record_lookup_field github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_expanded_record_lookup_field
func F_expanded_record_lookup_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expanded_record_fetch_field github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_fetch_field
func F_expanded_record_fetch_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_expanded_record_set_field_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_expanded_record_set_field_internal
func F_expanded_record_set_field_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_float_overflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_overflow_error
func F_float_overflow_error(m *base.Module)
//go:linkname F_float_underflow_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_float_underflow_error
func F_float_underflow_error(m *base.Module)
//go:linkname F_float_zero_divide_error github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_float_zero_divide_error
func F_float_zero_divide_error(m *base.Module)
//go:linkname F_init_degree_constants github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_init_degree_constants
func F_init_degree_constants(m *base.Module)
//go:linkname F_format_type_be github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_type_be
func F_format_type_be(m *base.Module, l0 int32) int32
//go:linkname F_format_type_with_typemod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_type_with_typemod
func F_format_type_with_typemod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_str_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_str_tolower
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_str_toupper github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_str_toupper
func F_str_toupper(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_do_to_timestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_do_to_timestamp
func F_do_to_timestamp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_NUM_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_cache
func F_NUM_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_NUM_processor github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_NUM_processor
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_int_to_roman github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int_to_roman
func F_int_to_roman(m *base.Module, l0 int32) int32
//go:linkname F_read_binary_file github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_read_binary_file
func F_read_binary_file(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pg_ls_dir_files github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_ls_dir_files
func F_pg_ls_dir_files(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pair_decode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pair_decode
func F_pair_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_box_ar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_box_ar
func F_box_ar(m *base.Module, l0 int32) float64
//go:linkname F_point_dt github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_point_dt
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_point_sl github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_point_sl
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_lseg_interpt_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lseg_interpt_lseg
func F_lseg_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_interpt_line github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_lseg_interpt_line
func F_lseg_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lseg_closept_point github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lseg_closept_point
func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_box_interpt_lseg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_box_interpt_lseg
func F_box_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dist_ppoly_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_dist_ppoly_internal
func F_dist_ppoly_internal(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_poly_overlap_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_poly_overlap_internal
func F_poly_overlap_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lseg_inside_poly github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lseg_inside_poly
func F_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_poly_to_circle github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_poly_to_circle
func F_poly_to_circle(m *base.Module, l0 int32, l1 int32)
//go:linkname F_int8dec github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_int8dec
func F_int8dec(m *base.Module, l0 int32) int32
//go:linkname F_escape_json github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_escape_json
func F_escape_json(m *base.Module, l0 int32, l1 int32)
//go:linkname F_datum_to_json_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_datum_to_json_internal
func F_datum_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_escape_json_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_escape_json_with_len
func F_escape_json_with_len(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_build_object_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_build_object_worker
func F_json_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_json_build_array_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_build_array_worker
func F_json_build_array_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_json_validate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_json_validate
func F_json_validate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbToCStringWorker github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbToCStringWorker
func F_JsonbToCStringWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_JsonbTypeName github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbTypeName
func F_JsonbTypeName(m *base.Module, l0 int32) int32
//go:linkname F_JsonbExtractScalar github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbExtractScalar
func F_JsonbExtractScalar(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_datum_to_jsonb_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_datum_to_jsonb_internal
func F_datum_to_jsonb_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_jsonb_build_object_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_jsonb_build_object_worker
func F_jsonb_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_jsonb_build_array_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonb_build_array_worker
func F_jsonb_build_array_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_cannotCastJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_cannotCastJsonbValue
func F_cannotCastJsonbValue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_make_scalar_key github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_scalar_key
func F_make_scalar_key(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_extract_jsp_path_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_extract_jsp_path_expr
func F_extract_jsp_path_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_extract_jsp_bool_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_extract_jsp_bool_expr
func F_extract_jsp_bool_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_emit_jsp_gin_entries github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_emit_jsp_gin_entries
func F_emit_jsp_gin_entries(m *base.Module, l0 int32, l1 int32)
//go:linkname F_execute_jsp_gin_node github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_execute_jsp_gin_node
func F_execute_jsp_gin_node(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbValueToJsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_JsonbValueToJsonb
func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32
//go:linkname F_pushJsonbValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pushJsonbValue
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_iteratorFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_iteratorFromContainer
func F_iteratorFromContainer(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbIteratorNext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_JsonbIteratorNext
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compareJsonbContainers github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_compareJsonbContainers
func F_compareJsonbContainers(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_findJsonbValueFromContainer github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_findJsonbValueFromContainer
func F_findJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_JsonbDeepContains github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbDeepContains
func F_JsonbDeepContains(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_JsonbHashScalarValue github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_JsonbHashScalarValue
func F_JsonbHashScalarValue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_parse_json_or_errsave github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_parse_json_or_errsave
func F_pg_parse_json_or_errsave(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_json_errsave_error github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_json_errsave_error
func F_json_errsave_error(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_worker
func F_get_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_jsonb_path_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_jsonb_path_all
func F_get_jsonb_path_all(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_each_worker_jsonb github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_each_worker_jsonb
func F_each_worker_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_populate_record_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_worker
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_populate_record_field github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_populate_record_field
func F_populate_record_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_prepare_column_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_prepare_column_cache
func F_prepare_column_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_parse_jsonb_index_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_jsonb_index_flags
func F_parse_jsonb_index_flags(m *base.Module, l0 int32) int32
//go:linkname F_iterate_jsonb_values github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_iterate_jsonb_values
func F_iterate_jsonb_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_iterate_json_values github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iterate_json_values
func F_iterate_json_values(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_categorize_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_json_categorize_type
func F_json_categorize_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_populate_record github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_populate_record
func F_populate_record(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_jsonPathFromCstring github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jsonPathFromCstring
func F_jsonPathFromCstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_jspInitByBuffer github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jspInitByBuffer
func F_jspInitByBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_jspInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jspInit
func F_jspInit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jsonb_path_query_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_jsonb_path_query_internal
func F_jsonb_path_query_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_executeItemOptUnwrapTarget github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_executeItemOptUnwrapTarget
func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_executeItemOptUnwrapResult github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_executeItemOptUnwrapResult
func F_executeItemOptUnwrapResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_yy_fatal_error_5 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_yy_fatal_error_5
func F_yy_fatal_error_5(m *base.Module, l0 int32)
//go:linkname F_jsonpath_yyerror github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_jsonpath_yyerror
func F_jsonpath_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addUnicodeChar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_addUnicodeChar
func F_addUnicodeChar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GenericMatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GenericMatchText
func F_GenericMatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SB_MatchText github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SB_MatchText
func F_SB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_Generic_Text_IC_like github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_Generic_Text_IC_like
func F_Generic_Text_IC_like(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_like_regex_support github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_like_regex_support
func F_like_regex_support(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_patternsel_common github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_patternsel_common
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64
//go:linkname F_current_database github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_current_database
func F_current_database(m *base.Module, l0 int32) int32
//go:linkname F_make_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_make_multirange
func F_make_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_multirange_get_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_multirange_get_range
func F_multirange_get_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_get_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_multirange_get_bounds
func F_multirange_get_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_multirange_intersect_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_multirange_intersect_internal
func F_multirange_intersect_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_multirange_contains_elem_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_multirange_contains_elem_internal
func F_multirange_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_multirange_contains_range_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_multirange_contains_range_internal
func F_multirange_contains_range_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_contains_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_contains_multirange_internal
func F_range_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overlaps_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_overlaps_multirange_internal
func F_range_overlaps_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overright_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_overright_multirange_internal
func F_range_overright_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_before_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_before_multirange_internal
func F_range_before_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_after_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_after_multirange_internal
func F_range_after_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_adjacent_multirange_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_adjacent_multirange_internal
func F_range_adjacent_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_namestrcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_namestrcmp
func F_namestrcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_current_user github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_current_user
func F_current_user(m *base.Module, l0 int32) int32
//go:linkname F_current_schema github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_current_schema
func F_current_schema(m *base.Module, l0 int32) int32
//go:linkname F_convert_network_to_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_convert_network_to_scalar
func F_convert_network_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_inet_hist_value_sel github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_inet_hist_value_sel
func F_inet_hist_value_sel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64
//go:linkname F_make_result_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_make_result_opt_error
func F_make_result_opt_error(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_var_from_str github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_var_from_str
func F_set_var_from_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_mul_var github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_mul_var
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_add_var github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_add_var
func F_add_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_apply_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_apply_typmod
func F_apply_typmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_str_from_var github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_str_from_var
func F_get_str_from_var(m *base.Module, l0 int32) int32
//go:linkname F_div_var github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_div_var
func F_div_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_cmp_abs_common github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_cmp_abs_common
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sub_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sub_var
func F_sub_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_div_var_int github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_div_var_int
func F_div_var_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_numeric_mul_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_numeric_mul_opt_error
func F_numeric_mul_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_numeric_mod_opt_error github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_numeric_mod_opt_error
func F_numeric_mod_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_gcd_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_gcd_var
func F_gcd_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sqrt_var github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sqrt_var
func F_sqrt_var(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_int64_to_numeric github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_int64_to_numeric
func F_int64_to_numeric(m *base.Module, l0 int64) int32
//go:linkname F_do_numeric_accum github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_do_numeric_accum
func F_do_numeric_accum(m *base.Module, l0 int32, l1 int32)
//go:linkname F_accum_sum_add github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_accum_sum_add
func F_accum_sum_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_accum_sum_final github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_accum_sum_final
func F_accum_sum_final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numericvar_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_numericvar_serialize
func F_numericvar_serialize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_numeric_stddev_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_numeric_stddev_internal
func F_numeric_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_numeric_poly_stddev_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_numeric_poly_stddev_internal
func F_numeric_poly_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_ultoa_n github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_ultoa_n
func F_pg_ultoa_n(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_buildoidvector github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_buildoidvector
func F_buildoidvector(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_check_valid_oidvector github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_valid_oidvector
func F_check_valid_oidvector(m *base.Module, l0 int32)
//go:linkname F_ordered_set_startup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ordered_set_startup
func F_ordered_set_startup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_percentile_cont_multi_final_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_percentile_cont_multi_final_common
func F_percentile_cont_multi_final_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hypothetical_check_argtypes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hypothetical_check_argtypes
func F_hypothetical_check_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_check_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_locale
func F_check_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_PGLC_localeconv github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PGLC_localeconv
func F_PGLC_localeconv(m *base.Module) int32
//go:linkname F_pg_newlocale_from_collation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_newlocale_from_collation
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32
//go:linkname F_pg_strncoll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strncoll
func F_pg_strncoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pg_strnxfrm github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_strnxfrm
func F_pg_strnxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_builtin_locale_encoding github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_builtin_locale_encoding
func F_builtin_locale_encoding(m *base.Module, l0 int32) int32
//go:linkname F_builtin_validate_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_builtin_validate_locale
func F_builtin_validate_locale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_icu_language_tag github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_icu_language_tag
func F_icu_language_tag(m *base.Module) int32
//go:linkname F_icu_validate_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_icu_validate_locale
func F_icu_validate_locale(m *base.Module)
//go:linkname F_get_collation_actual_version_builtin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_collation_actual_version_builtin
func F_get_collation_actual_version_builtin(m *base.Module, l0 int32) int32
//go:linkname F_char2wchar github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_char2wchar
func F_char2wchar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_pg_stat_wal_build_tuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_stat_wal_build_tuple
func F_pg_stat_wal_build_tuple(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_quote_literal_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_quote_literal_cstr
func F_quote_literal_cstr(m *base.Module, l0 int32) int32
//go:linkname F_get_range_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_range_io_data
func F_get_range_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_range
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_serialize github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_serialize
func F_range_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_range_deserialize github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_deserialize
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_range_get_typcache github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_get_typcache
func F_range_get_typcache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_range_contains_elem_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_contains_elem_internal
func F_range_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_eq_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_eq_internal
func F_range_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_cmp_bounds github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_cmp_bounds
func F_range_cmp_bounds(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_contains_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_contains_internal
func F_range_contains_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_contained_by_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_contained_by_internal
func F_range_contained_by_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_before_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_range_before_internal
func F_range_before_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_adjacent_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_adjacent_internal
func F_range_adjacent_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overleft_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_range_overleft_internal
func F_range_overleft_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_overright_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_range_overright_internal
func F_range_overright_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_range_minus_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_range_minus_internal
func F_range_minus_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_make_empty_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_make_empty_range
func F_make_empty_range(m *base.Module, l0 int32) int32
//go:linkname F_range_super_union github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_range_super_union
func F_range_super_union(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_compile_and_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RE_compile_and_cache
func F_RE_compile_and_cache(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RE_wchar_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RE_wchar_execute
func F_RE_wchar_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_re_flags github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_parse_re_flags
func F_parse_re_flags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_similar_escape_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_similar_escape_internal
func F_similar_escape_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setup_regexp_matches github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_setup_regexp_matches
func F_setup_regexp_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_regexp_split_to_array github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_regexp_split_to_array
func F_regexp_split_to_array(m *base.Module, l0 int32) int32
//go:linkname F_format_procedure github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_procedure
func F_format_procedure(m *base.Module, l0 int32) int32
//go:linkname F_format_operator_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_format_operator_extended
func F_format_operator_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_format_operator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_format_operator
func F_format_operator(m *base.Module, l0 int32) int32
//go:linkname F_ri_FetchConstraintInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ri_FetchConstraintInfo
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ri_PlanCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ri_PlanCheck
func F_ri_PlanCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ri_PerformCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ri_PerformCheck
func F_ri_PerformCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_ri_restrict github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ri_restrict
func F_ri_restrict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_record_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_record_cmp
func F_record_cmp(m *base.Module, l0 int32) int32
//go:linkname F_record_image_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_record_image_cmp
func F_record_image_cmp(m *base.Module, l0 int32) int32
//go:linkname F_pg_get_ruledef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_get_ruledef_worker
func F_pg_get_ruledef_worker(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_quote_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_quote_identifier
func F_quote_identifier(m *base.Module, l0 int32) int32
//go:linkname F_generate_relation_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_relation_name
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_rule_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_rule_expr
func F_get_rule_expr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_rtable_names github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_set_rtable_names
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_set_simple_column_names github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_set_simple_column_names
func F_set_simple_column_names(m *base.Module, l0 int32)
//go:linkname F_pg_get_indexdef_worker github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_get_indexdef_worker
func F_pg_get_indexdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_get_reloptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_reloptions
func F_get_reloptions(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_rule_windowspec github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rule_windowspec
func F_get_rule_windowspec(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_deparse_context_for github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_deparse_context_for
func F_deparse_context_for(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_deparse_expression github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_deparse_expression
func F_deparse_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_print_function_arguments github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_print_function_arguments
func F_print_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_quote_qualified_identifier github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_quote_qualified_identifier
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_set_deparse_context_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_set_deparse_context_plan
func F_set_deparse_context_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_generate_operator_clause github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_operator_clause
func F_generate_operator_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_colname_is_unique github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_colname_is_unique
func F_colname_is_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_json_behavior github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_json_behavior
func F_get_json_behavior(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_simple_binary_op_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_simple_binary_op_name
func F_get_simple_binary_op_name(m *base.Module, l0 int32) int32
//go:linkname F_eqsel_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_eqsel_internal
func F_eqsel_internal(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_get_restriction_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_restriction_variable
func F_get_restriction_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_mcv_selectivity github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_mcv_selectivity
func F_mcv_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64
//go:linkname F_get_actual_variable_endpoint github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_actual_variable_endpoint
func F_get_actual_variable_endpoint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_convert_numeric_to_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_convert_numeric_to_scalar
func F_convert_numeric_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_convert_string_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_convert_string_datum
func F_convert_string_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_all_rows_selectable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_all_rows_selectable
func F_all_rows_selectable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_join_variables github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_join_variables
func F_get_join_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_AdjustTimestampForTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AdjustTimestampForTypmod
func F_AdjustTimestampForTypmod(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_timestamp2tm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_timestamp2tm
func F_timestamp2tm(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_timestamp2timestamptz_opt_overflow github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp2timestamptz_opt_overflow
func F_timestamp2timestamptz_opt_overflow(m *base.Module, l0 int64, l1 int32) int64
//go:linkname F_GetCurrentTimestamp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCurrentTimestamp
func F_GetCurrentTimestamp(m *base.Module) int64
//go:linkname F_timestamptz_to_str github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_timestamptz_to_str
func F_timestamptz_to_str(m *base.Module, l0 int64) int32
//go:linkname F_timestamp_cmp_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamp_cmp_internal
func F_timestamp_cmp_internal(m *base.Module, l0 int64, l1 int64) int32
//go:linkname F_interval_um_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_interval_um_internal
func F_interval_um_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_timestamptz_pl_interval_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_timestamptz_pl_interval_internal
func F_timestamptz_pl_interval_internal(m *base.Module, l0 int64, l1 int32, l2 int32) int64
//go:linkname F_pushValue github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pushValue
func F_pushValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dofindsubquery github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dofindsubquery
func F_dofindsubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_QTNFree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNFree
func F_QTNFree(m *base.Module, l0 int32)
//go:linkname F_QTNSort github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_QTNSort
func F_QTNSort(m *base.Module, l0 int32)
//go:linkname F_QTNTernary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTNTernary
func F_QTNTernary(m *base.Module, l0 int32)
//go:linkname F_QTNBinary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTNBinary
func F_QTNBinary(m *base.Module, l0 int32)
//go:linkname F_QTN2QT github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_QTN2QT
func F_QTN2QT(m *base.Module, l0 int32) int32
//go:linkname F_getWeights github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getWeights
func F_getWeights(m *base.Module, l0 int32, l1 int32)
//go:linkname F_calc_rank github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_calc_rank
func F_calc_rank(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32
//go:linkname F_calc_rank_cd github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_calc_rank_cd
func F_calc_rank_cd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32
//go:linkname F_tsvector_delete_by_indices github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tsvector_delete_by_indices
func F_tsvector_delete_by_indices(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_TS_phrase_execute github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_TS_phrase_execute
func F_TS_phrase_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tsvector_update_trigger github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tsvector_update_trigger
func F_tsvector_update_trigger(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_uuidv7 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_generate_uuidv7
func F_generate_uuidv7(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_anybit_typmodin github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_anybit_typmodin
func F_anybit_typmodin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bit_catenate github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_bit_catenate
func F_bit_catenate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_bitsubstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_bitsubstring
func F_bitsubstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_bpchar_input github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bpchar_input
func F_bpchar_input(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_anychar_typmodin github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_anychar_typmodin
func F_anychar_typmodin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_cstring_to_text github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_cstring_to_text
func F_cstring_to_text(m *base.Module, l0 int32) int32
//go:linkname F_cstring_to_text_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_cstring_to_text_with_len
func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_to_cstring github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_text_to_cstring
func F_text_to_cstring(m *base.Module, l0 int32) int32
//go:linkname F_text_to_cstring_buffer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_text_to_cstring_buffer
func F_text_to_cstring_buffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_text_catenate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_catenate
func F_text_catenate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_text_substring github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_substring
func F_text_substring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_text_position_setup github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_position_setup
func F_text_position_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_text_position_next github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_text_position_next
func F_text_position_next(m *base.Module, l0 int32) int32
//go:linkname F_varstr_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstr_cmp
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_varstr_sortsupport github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_varstr_sortsupport
func F_varstr_sortsupport(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_varstrfastcmp_locale github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_varstrfastcmp_locale
func F_varstrfastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_textToQualifiedNameList github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_textToQualifiedNameList
func F_textToQualifiedNameList(m *base.Module, l0 int32) int32
//go:linkname F_text_format github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_text_format
func F_text_format(m *base.Module, l0 int32) int32
//go:linkname F_xmlconcat github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_xmlconcat
func F_xmlconcat(m *base.Module) int32
//go:linkname F_xmlparse github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_xmlparse
func F_xmlparse(m *base.Module) int32
//go:linkname F_map_sql_identifier_to_xml_name github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_map_sql_identifier_to_xml_name
func F_map_sql_identifier_to_xml_name(m *base.Module) int32
//go:linkname F_map_sql_value_to_xml_value github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_map_sql_value_to_xml_value
func F_map_sql_value_to_xml_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_map_sql_typecoll_to_xmlschema_types github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_map_sql_typecoll_to_xmlschema_types
func F_map_sql_typecoll_to_xmlschema_types(m *base.Module, l0 int32) int32
//go:linkname F_CreateCacheMemoryContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CreateCacheMemoryContext
func F_CreateCacheMemoryContext(m *base.Module)
//go:linkname F_CatalogCacheInitializeCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CatalogCacheInitializeCache
func F_CatalogCacheInitializeCache(m *base.Module, l0 int32)
//go:linkname F_SearchCatCacheInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_SearchCatCacheInternal
func F_SearchCatCacheInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ReleaseCatCache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReleaseCatCache
func F_ReleaseCatCache(m *base.Module, l0 int32)
//go:linkname F_ReleaseCatCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseCatCacheList
func F_ReleaseCatCacheList(m *base.Module, l0 int32)
//go:linkname F_EventCacheLookup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_EventCacheLookup
func F_EventCacheLookup(m *base.Module, l0 int32) int32
//go:linkname F_LocalExecuteInvalidationMessage github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LocalExecuteInvalidationMessage
func F_LocalExecuteInvalidationMessage(m *base.Module, l0 int32)
//go:linkname F_ProcessCommittedInvalidationMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ProcessCommittedInvalidationMessages
func F_ProcessCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_CommandEndInvalidationMessages github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CommandEndInvalidationMessages
func F_CommandEndInvalidationMessages(m *base.Module)
//go:linkname F_CacheInvalidateHeapTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CacheInvalidateHeapTuple
func F_CacheInvalidateHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CacheInvalidateRelcache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CacheInvalidateRelcache
func F_CacheInvalidateRelcache(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateRelcacheAll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CacheInvalidateRelcacheAll
func F_CacheInvalidateRelcacheAll(m *base.Module)
//go:linkname F_CacheInvalidateRelcacheByTuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CacheInvalidateRelcacheByTuple
func F_CacheInvalidateRelcacheByTuple(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateRelcacheByRelid github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheInvalidateRelcacheByRelid
func F_CacheInvalidateRelcacheByRelid(m *base.Module, l0 int32)
//go:linkname F_CacheInvalidateSmgr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CacheInvalidateSmgr
func F_CacheInvalidateSmgr(m *base.Module, l0 int32)
//go:linkname F_CacheRegisterSyscacheCallback github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CacheRegisterSyscacheCallback
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_op_in_opfamily github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_op_in_opfamily
func F_op_in_opfamily(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_op_opfamily_properties github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_op_opfamily_properties
func F_get_op_opfamily_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_get_opfamily_member github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_opfamily_member
func F_get_opfamily_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_opfamily_member_for_cmptype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_opfamily_member_for_cmptype
func F_get_opfamily_member_for_cmptype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_equality_op_for_ordering_op github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_equality_op_for_ordering_op
func F_get_equality_op_for_ordering_op(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_mergejoin_opfamilies github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_mergejoin_opfamilies
func F_get_mergejoin_opfamilies(m *base.Module, l0 int32) int32
//go:linkname F_get_compatible_hash_operators github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_compatible_hash_operators
func F_get_compatible_hash_operators(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_opfamily_proc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opfamily_proc
func F_get_opfamily_proc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_get_negator github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_negator
func F_get_negator(m *base.Module, l0 int32) int32
//go:linkname F_get_attname github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_attname
func F_get_attname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_get_attnum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attnum
func F_get_attnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_atttype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_atttype
func F_get_atttype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_attoptions github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_attoptions
func F_get_attoptions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_constraint_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_constraint_index
func F_get_constraint_index(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_family github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opclass_family
func F_get_opclass_family(m *base.Module, l0 int32) int32
//go:linkname F_get_opclass_input_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opclass_input_type
func F_get_opclass_input_type(m *base.Module, l0 int32) int32
//go:linkname F_get_opfamily_name github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_opfamily_name
func F_get_opfamily_name(m *base.Module, l0 int32) int32
//go:linkname F_get_opcode github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_opcode
func F_get_opcode(m *base.Module, l0 int32) int32
//go:linkname F_op_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_op_strict
func F_op_strict(m *base.Module, l0 int32) int32
//go:linkname F_func_strict github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_func_strict
func F_func_strict(m *base.Module, l0 int32) int32
//go:linkname F_func_volatile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_func_volatile
func F_func_volatile(m *base.Module, l0 int32) int32
//go:linkname F_get_commutator github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_commutator
func F_get_commutator(m *base.Module, l0 int32) int32
//go:linkname F_get_func_name github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_name
func F_get_func_name(m *base.Module, l0 int32) int32
//go:linkname F_get_func_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_rettype
func F_get_func_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_func_prokind github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_prokind
func F_get_func_prokind(m *base.Module, l0 int32) int32
//go:linkname F_get_func_leakproof github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_func_leakproof
func F_get_func_leakproof(m *base.Module, l0 int32) int32
//go:linkname F_get_func_support github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_func_support
func F_get_func_support(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_rel_name
func F_get_rel_name(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_type_id github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_rel_type_id
func F_get_rel_type_id(m *base.Module, l0 int32) int32
//go:linkname F_get_rel_persistence github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_rel_persistence
func F_get_rel_persistence(m *base.Module, l0 int32) int32
//go:linkname F_get_typisdefined github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typisdefined
func F_get_typisdefined(m *base.Module, l0 int32) int32
//go:linkname F_get_typlen github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typlen
func F_get_typlen(m *base.Module, l0 int32) int32
//go:linkname F_get_typbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typbyval
func F_get_typbyval(m *base.Module, l0 int32) int32
//go:linkname F_get_typlenbyval github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_typlenbyval
func F_get_typlenbyval(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typlenbyvalalign github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typlenbyvalalign
func F_get_typlenbyvalalign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_type_io_data github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_type_io_data
func F_get_type_io_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_getBaseType github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseType
func F_getBaseType(m *base.Module, l0 int32) int32
//go:linkname F_getBaseTypeAndTypmod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getBaseTypeAndTypmod
func F_getBaseTypeAndTypmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_typtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_typtype
func F_get_typtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_rowtype github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_rowtype
func F_type_is_rowtype(m *base.Module, l0 int32) int32
//go:linkname F_type_is_enum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_enum
func F_type_is_enum(m *base.Module, l0 int32) int32
//go:linkname F_type_is_range github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_type_is_range
func F_type_is_range(m *base.Module, l0 int32) int32
//go:linkname F_type_is_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_type_is_multirange
func F_type_is_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_type_category_preferred github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_type_category_preferred
func F_get_type_category_preferred(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_array_type github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_array_type
func F_get_array_type(m *base.Module, l0 int32) int32
//go:linkname F_get_base_element_type github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_base_element_type
func F_get_base_element_type(m *base.Module, l0 int32) int32
//go:linkname F_getTypeOutputInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_getTypeOutputInfo
func F_getTypeOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_get_typcollation github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_typcollation
func F_get_typcollation(m *base.Module, l0 int32) int32
//go:linkname F_type_is_collatable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_type_is_collatable
func F_type_is_collatable(m *base.Module, l0 int32) int32
//go:linkname F_get_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_attstatsslot
func F_get_attstatsslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_free_attstatsslot github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_free_attstatsslot
func F_free_attstatsslot(m *base.Module, l0 int32)
//go:linkname F_get_namespace_name github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_namespace_name
func F_get_namespace_name(m *base.Module, l0 int32) int32
//go:linkname F_get_namespace_name_or_temp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_namespace_name_or_temp
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32
//go:linkname F_get_range_multirange github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_range_multirange
func F_get_range_multirange(m *base.Module, l0 int32) int32
//go:linkname F_get_index_isvalid github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_index_isvalid
func F_get_index_isvalid(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetPartitionKey github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetPartitionKey
func F_RelationGetPartitionKey(m *base.Module, l0 int32) int32
//go:linkname F_GetCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetCachedPlan
func F_GetCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_RevalidateCachedQuery github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RevalidateCachedQuery
func F_RevalidateCachedQuery(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReleaseCachedPlan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ReleaseCachedPlan
func F_ReleaseCachedPlan(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CachedPlanAllowsSimpleValidityCheck github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_CachedPlanAllowsSimpleValidityCheck
func F_CachedPlanAllowsSimpleValidityCheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_RelationInitTableAccessMethod github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationInitTableAccessMethod
func F_RelationInitTableAccessMethod(m *base.Module, l0 int32)
//go:linkname F_RelationInitPhysicalAddr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationInitPhysicalAddr
func F_RelationInitPhysicalAddr(m *base.Module, l0 int32)
//go:linkname F_RelationClearRelation github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationClearRelation
func F_RelationClearRelation(m *base.Module, l0 int32)
//go:linkname F_RelationGetStatExtList github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationGetStatExtList
func F_RelationGetStatExtList(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetPrimaryKeyIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetPrimaryKeyIndex
func F_RelationGetPrimaryKeyIndex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_RelationGetReplicaIndex github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RelationGetReplicaIndex
func F_RelationGetReplicaIndex(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexExpressions github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexExpressions
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexPredicate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RelationGetIndexPredicate
func F_RelationGetIndexPredicate(m *base.Module, l0 int32) int32
//go:linkname F_RelationGetIndexAttrBitmap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationGetIndexAttrBitmap
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_errtable github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errtable
func F_errtable(m *base.Module, l0 int32)
//go:linkname F_RelationCacheInitFilePreInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationCacheInitFilePreInvalidate
func F_RelationCacheInitFilePreInvalidate(m *base.Module)
//go:linkname F_unlink_initfile github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_unlink_initfile
func F_unlink_initfile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RelationCacheInitFilePostInvalidate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_RelationCacheInitFilePostInvalidate
func F_RelationCacheInitFilePostInvalidate(m *base.Module)
//go:linkname F_RelationMapOidToFilenumberForDatabase github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_RelationMapOidToFilenumberForDatabase
func F_RelationMapOidToFilenumberForDatabase(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_read_relmap_file github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_read_relmap_file
func F_read_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_write_relmap_file github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_write_relmap_file
func F_write_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_RelationMapUpdateMap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_RelationMapUpdateMap
func F_RelationMapUpdateMap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_get_tablespace github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_tablespace
func F_get_tablespace(m *base.Module, l0 int32) int32
//go:linkname F_get_tablespace_maintenance_io_concurrency github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_get_tablespace_maintenance_io_concurrency
func F_get_tablespace_maintenance_io_concurrency(m *base.Module, l0 int32) int32
//go:linkname F_SearchSysCache1 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCache1
func F_SearchSysCache1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCache2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache2
func F_SearchSysCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCache4 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCache4
func F_SearchSysCache4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SearchSysCacheCopy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SearchSysCacheCopy
func F_SearchSysCacheCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheLockedCopy1 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheLockedCopy1
func F_SearchSysCacheLockedCopy1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SearchSysCacheExists github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_SearchSysCacheExists
func F_SearchSysCacheExists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_GetSysCacheOid github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetSysCacheOid
func F_GetSysCacheOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_SysCacheGetAttr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SysCacheGetAttr
func F_SysCacheGetAttr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SysCacheGetAttrNotNull github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SysCacheGetAttrNotNull
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_SearchSysCacheList github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SearchSysCacheList
func F_SearchSysCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lookup_ts_parser_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_ts_parser_cache
func F_lookup_ts_parser_cache(m *base.Module, l0 int32) int32
//go:linkname F_getTSCurrentConfig github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_getTSCurrentConfig
func F_getTSCurrentConfig(m *base.Module) int32
//go:linkname F_lookup_type_cache github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_type_cache
func F_lookup_type_cache(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_lookup_rowtype_tupdesc_internal
func F_lookup_rowtype_tupdesc_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookup_rowtype_tupdesc_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_lookup_rowtype_tupdesc_domain
func F_lookup_rowtype_tupdesc_domain(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_assign_record_type_typmod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_assign_record_type_typmod
func F_assign_record_type_typmod(m *base.Module, l0 int32)
//go:linkname F_load_enum_cache_data github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_load_enum_cache_data
func F_load_enum_cache_data(m *base.Module, l0 int32)
//go:linkname F_errstart github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_errstart
func F_errstart(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_write_stderr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_write_stderr
func F_write_stderr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errmsg_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errmsg_internal
func F_errmsg_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_EmitErrorReport github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_EmitErrorReport
func F_EmitErrorReport(m *base.Module)
//go:linkname F_FreeErrorDataContents github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_FreeErrorDataContents
func F_FreeErrorDataContents(m *base.Module, l0 int32)
//go:linkname F_pg_re_throw github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_re_throw
func F_pg_re_throw(m *base.Module)
//go:linkname F_errsave_start github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errsave_start
func F_errsave_start(m *base.Module, l0 int32) int32
//go:linkname F_errsave_finish github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errsave_finish
func F_errsave_finish(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errcode_for_file_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errcode_for_file_access
func F_errcode_for_file_access(m *base.Module)
//go:linkname F_errcode_for_socket_access github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errcode_for_socket_access
func F_errcode_for_socket_access(m *base.Module)
//go:linkname F_errmsg github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errmsg
func F_errmsg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errmsg_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errmsg_plural
func F_errmsg_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_errdetail_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_errdetail_internal
func F_errdetail_internal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_errdetail_plural github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_errdetail_plural
func F_errdetail_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_set_errcontext_domain github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_set_errcontext_domain
func F_set_errcontext_domain(m *base.Module, l0 int32)
//go:linkname F_errhidestmt github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_errhidestmt
func F_errhidestmt(m *base.Module)
//go:linkname F_errhidecontext github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errhidecontext
func F_errhidecontext(m *base.Module)
//go:linkname F_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_errposition
func F_errposition(m *base.Module, l0 int32) int32
//go:linkname F_internalerrposition github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internalerrposition
func F_internalerrposition(m *base.Module, l0 int32)
//go:linkname F_err_generic_string github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_err_generic_string
func F_err_generic_string(m *base.Module, l0 int32, l1 int32)
//go:linkname F_format_elog_string github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_format_elog_string
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_FlushErrorState github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_FlushErrorState
func F_FlushErrorState(m *base.Module)
//go:linkname F_load_external_function github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_load_external_function
func F_load_external_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_find_in_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_find_in_path
func F_find_in_path(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_substitute_path_macro github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_substitute_path_macro
func F_substitute_path_macro(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fmgr_info github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info
func F_fmgr_info(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fetch_finfo_record github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fetch_finfo_record
func F_fetch_finfo_record(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fmgr_info_cxt github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_fmgr_info_cxt
func F_fmgr_info_cxt(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fmgr_info_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fmgr_info_copy
func F_fmgr_info_copy(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pg_detoast_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_detoast_datum
func F_pg_detoast_datum(m *base.Module, l0 int32) int32
//go:linkname F_DirectFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall1Coll
func F_DirectFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DirectFunctionCall2Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall2Coll
func F_DirectFunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_DirectFunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_DirectFunctionCall3Coll
func F_DirectFunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_DirectFunctionCall4Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_DirectFunctionCall4Coll
func F_DirectFunctionCall4Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_FunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall1Coll
func F_FunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_FunctionCall3Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_FunctionCall3Coll
func F_FunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_OidFunctionCall1Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_OidFunctionCall1Coll
func F_OidFunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_OidFunctionCall6Coll github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_OidFunctionCall6Coll
func F_OidFunctionCall6Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_InputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InputFunctionCall
func F_InputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_InputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InputFunctionCallSafe
func F_InputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_DirectInputFunctionCallSafe github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_DirectInputFunctionCallSafe
func F_DirectInputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_OutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OutputFunctionCall
func F_OutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ReceiveFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ReceiveFunctionCall
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_SendFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SendFunctionCall
func F_SendFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_OidInputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_OidInputFunctionCall
func F_OidInputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_OidOutputFunctionCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_OidOutputFunctionCall
func F_OidOutputFunctionCall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_Int64GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_Int64GetDatum
func F_Int64GetDatum(m *base.Module, l0 int64) int32
//go:linkname F_Float8GetDatum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_Float8GetDatum
func F_Float8GetDatum(m *base.Module, l0 float64) int32
//go:linkname F_pg_detoast_datum_copy github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_detoast_datum_copy
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32
//go:linkname F_pg_detoast_datum_packed github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_detoast_datum_packed
func F_pg_detoast_datum_packed(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_rettype github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_fn_expr_rettype
func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32
//go:linkname F_get_fn_expr_argtype github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_expr_argtype
func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_get_fn_opclass_options github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_fn_opclass_options
func F_get_fn_opclass_options(m *base.Module, l0 int32) int32
//go:linkname F_CheckFunctionValidatorAccess github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_CheckFunctionValidatorAccess
func F_CheckFunctionValidatorAccess(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_InitMaterializedSRF github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_InitMaterializedSRF
func F_InitMaterializedSRF(m *base.Module, l0 int32, l1 int32)
//go:linkname F_internal_get_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internal_get_result_type
func F_internal_get_result_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_get_call_result_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_call_result_type
func F_get_call_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_init_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_init_MultiFuncCall
func F_init_MultiFuncCall(m *base.Module, l0 int32) int32
//go:linkname F_end_MultiFuncCall github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_end_MultiFuncCall
func F_end_MultiFuncCall(m *base.Module, l0 int32)
//go:linkname F_get_expr_result_tupdesc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_get_expr_result_tupdesc
func F_get_expr_result_tupdesc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_create
func F_hash_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_destroy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hash_destroy
func F_hash_destroy(m *base.Module, l0 int32)
//go:linkname F_get_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_get_hash_value
func F_get_hash_value(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hash_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_search
func F_hash_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_search_with_hash_value github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_search_with_hash_value
func F_hash_search_with_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hash_seq_init github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_hash_seq_init
func F_hash_seq_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hash_seq_search github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_search
func F_hash_seq_search(m *base.Module, l0 int32) int32
//go:linkname F_hash_seq_term github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_hash_seq_term
func F_hash_seq_term(m *base.Module, l0 int32)
//go:linkname F_AtEOXact_HashTables github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_AtEOXact_HashTables
func F_AtEOXact_HashTables(m *base.Module, l0 int32)
//go:linkname F_GetUserNameFromId github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetUserNameFromId
func F_GetUserNameFromId(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_CreateLockFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_CreateLockFile
func F_CreateLockFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_BaseInit github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BaseInit
func F_BaseInit(m *base.Module)
//go:linkname F_InitPostgres github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_InitPostgres
func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_SwitchToUntrustedUser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_SwitchToUntrustedUser
func F_SwitchToUntrustedUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_RestoreUserContext github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_RestoreUserContext
func F_RestoreUserContext(m *base.Module, l0 int32)
//go:linkname F_local2local github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_local2local
func F_local2local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_latin2mic github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_latin2mic
func F_latin2mic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_mic2latin github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic2latin
func F_mic2latin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_latin2mic_with_table github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_latin2mic_with_table
func F_latin2mic_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_mic2latin_with_table github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mic2latin_with_table
func F_mic2latin_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_UtfToLocal github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UtfToLocal
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_LocalToUtf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LocalToUtf
func F_LocalToUtf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_perform_default_encoding_conversion github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_perform_default_encoding_conversion
func F_perform_default_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_unicode_to_server_noerror github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_unicode_to_server_noerror
func F_pg_unicode_to_server_noerror(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mb2wchar_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mb2wchar_with_len
func F_pg_mb2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mblen_cstr github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_cstr
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32
//go:linkname F_report_invalid_encoding_int github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_report_invalid_encoding_int
func F_report_invalid_encoding_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_mblen_range github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mblen_range
func F_pg_mblen_range(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mblen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_mblen_with_len
func F_pg_mblen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mblen_unbounded github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_mblen_unbounded
func F_pg_mblen_unbounded(m *base.Module, l0 int32) int32
//go:linkname F_pg_mbstrlen_with_len github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbstrlen_with_len
func F_pg_mbstrlen_with_len(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_mbcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_mbcliplen
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_mbcharcliplen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_mbcharcliplen
func F_pg_mbcharcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_encoding_conversion_args github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_check_encoding_conversion_args
func F_check_encoding_conversion_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_appendStringInfoStringQuoted github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_appendStringInfoStringQuoted
func F_appendStringInfoStringQuoted(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_find_option github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_option
func F_find_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_set_config_with_handle github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_set_config_with_handle
func F_set_config_with_handle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_GetConfigOption github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetConfigOption
func F_GetConfigOption(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_guc_strdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_guc_strdup
func F_guc_strdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SetConfigOption github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_SetConfigOption
func F_SetConfigOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_guc_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_guc_malloc
func F_guc_malloc(m *base.Module, l0 int32) int32
//go:linkname F_RestrictSearchPath github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RestrictSearchPath
func F_RestrictSearchPath(m *base.Module)
//go:linkname F_AtEOXact_GUC github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AtEOXact_GUC
func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32)
//go:linkname F_BeginReportingGUCOptions github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BeginReportingGUCOptions
func F_BeginReportingGUCOptions(m *base.Module)
//go:linkname F_init_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_init_custom_variable
func F_init_custom_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_define_custom_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_define_custom_variable
func F_define_custom_variable(m *base.Module, l0 int32)
//go:linkname F_GetConfigOptionByName github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetConfigOptionByName
func F_GetConfigOptionByName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GUCArrayAdd github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GUCArrayAdd
func F_GUCArrayAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_GUCArrayDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GUCArrayDelete
func F_GUCArrayDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ProcessConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ProcessConfigFile
func F_ProcessConfigFile(m *base.Module, l0 int32)
//go:linkname F_ParseConfigFile github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_ParseConfigFile
func F_ParseConfigFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_pg_rusage_show github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_rusage_show
func F_pg_rusage_show(m *base.Module, l0 int32) int32
//go:linkname F_register_ENR github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_register_ENR
func F_register_ENR(m *base.Module, l0 int32, l1 int32)
//go:linkname F_check_enable_rls github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_check_enable_rls
func F_check_enable_rls(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_stack_depth github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_check_stack_depth
func F_check_stack_depth(m *base.Module)
//go:linkname F_stack_is_too_deep github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_stack_is_too_deep
func F_stack_is_too_deep(m *base.Module) int32
//go:linkname F_superuser github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_superuser
func F_superuser(m *base.Module) int32
//go:linkname F_superuser_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_superuser_arg
func F_superuser_arg(m *base.Module, l0 int32) int32
//go:linkname F_schedule_alarm github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_schedule_alarm
func F_schedule_alarm(m *base.Module, l0 int64)
//go:linkname F_AllocSetContextCreateInternal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_AllocSetContextCreateInternal
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_attach_internal github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_attach_internal
func F_attach_internal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_allocate_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_allocate_extended
func F_dsa_allocate_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dsa_free github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dsa_free
func F_dsa_free(m *base.Module, l0 int32, l1 int32)
//go:linkname F_get_segment_by_index github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_get_segment_by_index
func F_get_segment_by_index(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dsa_get_address github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dsa_get_address
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextReset github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextReset
func F_MemoryContextReset(m *base.Module, l0 int32)
//go:linkname F_MemoryContextResetOnly github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextResetOnly
func F_MemoryContextResetOnly(m *base.Module, l0 int32)
//go:linkname F_MemoryContextDelete github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextDelete
func F_MemoryContextDelete(m *base.Module, l0 int32)
//go:linkname F_GetMemoryChunkContext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_GetMemoryChunkContext
func F_GetMemoryChunkContext(m *base.Module, l0 int32) int32
//go:linkname F_GetMemoryChunkSpace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetMemoryChunkSpace
func F_GetMemoryChunkSpace(m *base.Module, l0 int32) int32
//go:linkname F_MemoryContextStats github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextStats
func F_MemoryContextStats(m *base.Module, l0 int32)
//go:linkname F_MemoryContextAlloc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_MemoryContextAlloc
func F_MemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextAllocZero github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_MemoryContextAllocZero
func F_MemoryContextAllocZero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ProcessLogMemoryContextInterrupt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_ProcessLogMemoryContextInterrupt
func F_ProcessLogMemoryContextInterrupt(m *base.Module)
//go:linkname F_palloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc
func F_palloc(m *base.Module, l0 int32) int32
//go:linkname F_palloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_palloc0
func F_palloc0(m *base.Module, l0 int32) int32
//go:linkname F_palloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_palloc_extended
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_palloc_aligned github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_palloc_aligned
func F_palloc_aligned(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pfree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pfree
func F_pfree(m *base.Module, l0 int32)
//go:linkname F_repalloc_extended github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc_extended
func F_repalloc_extended(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_repalloc0 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc0
func F_repalloc0(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_repalloc_huge github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_repalloc_huge
func F_repalloc_huge(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_MemoryContextStrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_MemoryContextStrdup
func F_MemoryContextStrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pstrdup
func F_pstrdup(m *base.Module, l0 int32) int32
//go:linkname F_pnstrdup github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pnstrdup
func F_pnstrdup(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetPortalByName github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetPortalByName
func F_GetPortalByName(m *base.Module, l0 int32) int32
//go:linkname F_PortalDrop github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_PortalDrop
func F_PortalDrop(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateNewPortal github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateNewPortal
func F_CreateNewPortal(m *base.Module) int32
//go:linkname F_MarkPortalFailed github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_MarkPortalFailed
func F_MarkPortalFailed(m *base.Module, l0 int32)
//go:linkname F_HoldPortal github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HoldPortal
func F_HoldPortal(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerCreate
func F_ResourceOwnerCreate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ResourceOwnerEnlarge github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerEnlarge
func F_ResourceOwnerEnlarge(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerRemember github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerRemember
func F_ResourceOwnerRemember(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ResourceOwnerForget github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ResourceOwnerForget
func F_ResourceOwnerForget(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_CreateAuxProcessResourceOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_CreateAuxProcessResourceOwner
func F_CreateAuxProcessResourceOwner(m *base.Module)
//go:linkname F_ReleaseAuxProcessResources github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_ReleaseAuxProcessResources
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32)
//go:linkname F_ResourceOwnerForgetLock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ResourceOwnerForgetLock
func F_ResourceOwnerForgetLock(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LogicalTapeSetClose github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeSetClose
func F_LogicalTapeSetClose(m *base.Module, l0 int32)
//go:linkname F_LogicalTapeCreate github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeCreate
func F_LogicalTapeCreate(m *base.Module, l0 int32) int32
//go:linkname F_LogicalTapeClose github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LogicalTapeClose
func F_LogicalTapeClose(m *base.Module, l0 int32)
//go:linkname F_ltsWriteBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ltsWriteBlock
func F_ltsWriteBlock(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_LogicalTapeRewindForRead github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_LogicalTapeRewindForRead
func F_LogicalTapeRewindForRead(m *base.Module, l0 int32, l1 int32)
//go:linkname F_LogicalTapeRead github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_LogicalTapeRead
func F_LogicalTapeRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_LogicalTapeFreeze github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_LogicalTapeFreeze
func F_LogicalTapeFreeze(m *base.Module, l0 int32, l1 int32)
//go:linkname F_qsort_interruptible github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_qsort_interruptible
func F_qsort_interruptible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_qsort_interruptible_med3 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_qsort_interruptible_med3
func F_qsort_interruptible_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_sts_end_parallel_scan github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_sts_end_parallel_scan
func F_sts_end_parallel_scan(m *base.Module, l0 int32)
//go:linkname F_PrepareSortSupportComparisonShim github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PrepareSortSupportComparisonShim
func F_PrepareSortSupportComparisonShim(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PrepareSortSupportFromOrderingOp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareSortSupportFromOrderingOp
func F_PrepareSortSupportFromOrderingOp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_PrepareSortSupportFromIndexRel github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_PrepareSortSupportFromIndexRel
func F_PrepareSortSupportFromIndexRel(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_begin_batch github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_batch
func F_tuplesort_begin_batch(m *base.Module, l0 int32)
//go:linkname F_tuplesort_end github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_end
func F_tuplesort_end(m *base.Module, l0 int32)
//go:linkname F_dumptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_dumptuples
func F_dumptuples(m *base.Module, l0 int32, l1 int32)
//go:linkname F_inittapes github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_inittapes
func F_inittapes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_performsort github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_performsort
func F_tuplesort_performsort(m *base.Module, l0 int32)
//go:linkname F_qsort_tuple_unsigned github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_qsort_tuple_unsigned
func F_qsort_tuple_unsigned(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_qsort_tuple_int32 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_qsort_tuple_int32
func F_qsort_tuple_int32(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_qsort_ssup github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_qsort_ssup
func F_qsort_ssup(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_beginmerge github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_beginmerge
func F_beginmerge(m *base.Module, l0 int32)
//go:linkname F_tuplesort_gettuple_common github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettuple_common
func F_tuplesort_gettuple_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplesort_merge_order github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplesort_merge_order
func F_tuplesort_merge_order(m *base.Module, l0 int64) int32
//go:linkname F_tuplesort_estimate_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_estimate_shared
func F_tuplesort_estimate_shared(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_initialize_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_initialize_shared
func F_tuplesort_initialize_shared(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_attach_shared github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_attach_shared
func F_tuplesort_attach_shared(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_begin_heap
func F_tuplesort_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_tuplesort_begin_index_btree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplesort_begin_index_btree
func F_tuplesort_begin_index_btree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplesort_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_puttupleslot
func F_tuplesort_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplesort_putgintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_putgintuple
func F_tuplesort_putgintuple(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_putdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplesort_putdatum
func F_tuplesort_putdatum(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_tuplesort_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_gettupleslot
func F_tuplesort_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_tuplesort_getheaptuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_getheaptuple
func F_tuplesort_getheaptuple(m *base.Module, l0 int32) int32
//go:linkname F_tuplesort_getgintuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplesort_getgintuple
func F_tuplesort_getgintuple(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_tuplesort_getdatum github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_tuplesort_getdatum
func F_tuplesort_getdatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_tuplestore_begin_heap github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_begin_heap
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tuplestore_set_eflags github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_set_eflags
func F_tuplestore_set_eflags(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_end github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_end
func F_tuplestore_end(m *base.Module, l0 int32)
//go:linkname F_tuplestore_select_read_pointer github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_select_read_pointer
func F_tuplestore_select_read_pointer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_puttupleslot
func F_tuplestore_puttupleslot(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_puttuple github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_tuplestore_puttuple
func F_tuplestore_puttuple(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tuplestore_putvalues github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_tuplestore_putvalues
func F_tuplestore_putvalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_tuplestore_gettupleslot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_gettupleslot
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tuplestore_skiptuples github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_skiptuples
func F_tuplestore_skiptuples(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_tuplestore_rescan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_tuplestore_rescan
func F_tuplestore_rescan(m *base.Module, l0 int32)
//go:linkname F_HeapTupleHeaderAdjustCmax github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_HeapTupleHeaderAdjustCmax
func F_HeapTupleHeaderAdjustCmax(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_GetTransactionSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_GetTransactionSnapshot
func F_GetTransactionSnapshot(m *base.Module) int32
//go:linkname F_InvalidateCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_InvalidateCatalogSnapshot
func F_InvalidateCatalogSnapshot(m *base.Module)
//go:linkname F_GetLatestSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetLatestSnapshot
func F_GetLatestSnapshot(m *base.Module) int32
//go:linkname F_GetCatalogSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_GetCatalogSnapshot
func F_GetCatalogSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_PushCopiedSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_PushCopiedSnapshot
func F_PushCopiedSnapshot(m *base.Module, l0 int32)
//go:linkname F_UpdateActiveSnapshotCommandId github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_UpdateActiveSnapshotCommandId
func F_UpdateActiveSnapshotCommandId(m *base.Module)
//go:linkname F_PopActiveSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_PopActiveSnapshot
func F_PopActiveSnapshot(m *base.Module)
//go:linkname F_RegisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_RegisterSnapshot
func F_RegisterSnapshot(m *base.Module, l0 int32) int32
//go:linkname F_UnregisterSnapshot github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_UnregisterSnapshot
func F_UnregisterSnapshot(m *base.Module, l0 int32)
//go:linkname F_UnregisterSnapshotFromOwner github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_UnregisterSnapshotFromOwner
func F_UnregisterSnapshotFromOwner(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tzparse github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_tzparse
func F_tzparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_localtime github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_localtime
func F_pg_localtime(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_gmtime github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_gmtime
func F_pg_gmtime(m *base.Module, l0 int32) int32
//go:linkname F_pg_next_dst_boundary github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_next_dst_boundary
func F_pg_next_dst_boundary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_pg_tz_acceptable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_tz_acceptable
func F_pg_tz_acceptable(m *base.Module, l0 int32) int32
//go:linkname F_pg_tzset_offset github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_tzset_offset
func F_pg_tzset_offset(m *base.Module, l0 int32) int32
//go:linkname F__fmt github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F__fmt
func F__fmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_jit_compile_expr github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_jit_compile_expr
func F_jit_compile_expr(m *base.Module, l0 int32) int32
//go:linkname F_binaryheap_add_unordered github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_add_unordered
func F_binaryheap_add_unordered(m *base.Module, l0 int32, l1 int32)
//go:linkname F_binaryheap_build github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_binaryheap_build
func F_binaryheap_build(m *base.Module, l0 int32)
//go:linkname F_binaryheap_remove_first github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_binaryheap_remove_first
func F_binaryheap_remove_first(m *base.Module, l0 int32) int32
//go:linkname F_binaryheap_replace_first github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_binaryheap_replace_first
func F_binaryheap_replace_first(m *base.Module, l0 int32, l1 int32)
//go:linkname F_CreateEmptyBlockRefTable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateEmptyBlockRefTable
func F_CreateEmptyBlockRefTable(m *base.Module) int32
//go:linkname F_BlockRefTableSetLimitBlock github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BlockRefTableSetLimitBlock
func F_BlockRefTableSetLimitBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_BlockRefTableMarkBlockModified github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_BlockRefTableMarkBlockModified
func F_BlockRefTableMarkBlockModified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_CreateBlockRefTableReader github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_CreateBlockRefTableReader
func F_CreateBlockRefTableReader(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_BlockRefTableRead github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_BlockRefTableRead
func F_BlockRefTableRead(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_BlockRefTableReaderGetBlocks github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_BlockRefTableReaderGetBlocks
func F_BlockRefTableReaderGetBlocks(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DestroyBlockRefTableReader github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_DestroyBlockRefTableReader
func F_DestroyBlockRefTableReader(m *base.Module, l0 int32)
//go:linkname F_pg_checksum_update github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_checksum_update
func F_pg_checksum_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_update_controlfile github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_update_controlfile
func F_update_controlfile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_char_to_encoding_private github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_char_to_encoding_private
func F_pg_char_to_encoding_private(m *base.Module, l0 int32) int32
//go:linkname F_normalize_exec_path github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_normalize_exec_path
func F_normalize_exec_path(m *base.Module, l0 int32) int32
//go:linkname F_get_dirent_type github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_dirent_type
func F_get_dirent_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hash_bytes_uint32 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_hash_bytes_uint32
func F_hash_bytes_uint32(m *base.Module, l0 int32) int32
//go:linkname F_pg_getnameinfo_all github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_getnameinfo_all
func F_pg_getnameinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_json_lex_number github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_json_lex_number
func F_json_lex_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_makeJsonLexContextCstringLen github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeJsonLexContextCstringLen
func F_makeJsonLexContextCstringLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_parse_object github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_parse_object
func F_parse_object(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_array github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_parse_array
func F_parse_array(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_scalar github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_parse_scalar
func F_parse_scalar(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replace_percent_placeholders github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_replace_percent_placeholders
func F_replace_percent_placeholders(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_psprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_psprintf
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pvsnprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pvsnprintf
func F_pvsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_GetDatabasePath github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_GetDatabasePath
func F_GetDatabasePath(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_GetRelationPath github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_GetRelationPath
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_rmtree github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_rmtree
func F_rmtree(m *base.Module, l0 int32) int32
//go:linkname F_makeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_makeStringInfo
func F_makeStringInfo(m *base.Module) int32
//go:linkname F_initStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_initStringInfo
func F_initStringInfo(m *base.Module, l0 int32)
//go:linkname F_appendStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_appendStringInfo
func F_appendStringInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_enlargeStringInfo github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_enlargeStringInfo
func F_enlargeStringInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_appendStringInfoChar github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_appendStringInfoChar
func F_appendStringInfoChar(m *base.Module, l0 int32, l1 int32)
//go:linkname F_case_index github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_case_index
func F_case_index(m *base.Module, l0 int32) int32
//go:linkname F_pg_u_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_u_isalnum
func F_pg_u_isalnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_wait_result_to_str github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_wait_result_to_str
func F_wait_result_to_str(m *base.Module, l0 int32) int32
//go:linkname F_pg_encoding_mblen github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_encoding_mblen
func F_pg_encoding_mblen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_encoding_mblen_or_incomplete github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_encoding_mblen_or_incomplete
func F_pg_encoding_mblen_or_incomplete(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_encoding_verifymbchar github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_encoding_verifymbchar
func F_pg_encoding_verifymbchar(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_encoding_max_length github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_encoding_max_length
func F_pg_encoding_max_length(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_cryptohash_create
func F_pg_cryptohash_create(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_init github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_cryptohash_init
func F_pg_cryptohash_init(m *base.Module, l0 int32) int32
//go:linkname F_pg_cryptohash_update github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_cryptohash_update
func F_pg_cryptohash_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bsearch_arg github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_bsearch_arg
func F_bsearch_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_join_path_components github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_join_path_components
func F_join_path_components(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_canonicalize_path_enc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_canonicalize_path_enc
func F_canonicalize_path_enc(m *base.Module, l0 int32)
//go:linkname F_get_share_path github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_get_share_path
func F_get_share_path(m *base.Module, l0 int32)
//go:linkname F_pg_usleep github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_usleep
func F_pg_usleep(m *base.Module, l0 int32)
//go:linkname F_pg_strcasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strcasecmp
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pg_strncasecmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_strncasecmp
func F_pg_strncasecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_tolower github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_tolower
func F_pg_tolower(m *base.Module, l0 int32) int32
//go:linkname F_pqsignal_be github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pqsignal_be
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_qsort github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pg_qsort
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_vsnprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_vsnprintf
func F_pg_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dostr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dostr
func F_dostr(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dopr_outchmulti github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_dopr_outchmulti
func F_dopr_outchmulti(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_leading_pad github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_leading_pad
func F_leading_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pg_snprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pg_snprintf
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pg_sprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pg_sprintf
func F_pg_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pg_printf github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pg_printf
func F_pg_printf(m *base.Module, l0 int32, l1 int32)
//go:linkname F_pg_strfromd github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pg_strfromd
func F_pg_strfromd(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_pgl_system github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_system
func F_pgl_system(m *base.Module, l0 int32) int32
//go:linkname F_pgl_exit github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pgl_exit
func F_pgl_exit(m *base.Module, l0 int32)
//go:linkname F_pgl_shmctl github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pgl_shmctl
func F_pgl_shmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_copy_plpgsql_datums github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_copy_plpgsql_datums
func F_copy_plpgsql_datums(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_stmt_block github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_exec_stmt_block
func F_exec_stmt_block(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_plpgsql_create_econtext github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_plpgsql_create_econtext
func F_plpgsql_create_econtext(m *base.Module, l0 int32)
//go:linkname F_exec_eval_datum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_exec_eval_datum
func F_exec_eval_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_instantiate_empty_record_variable github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_instantiate_empty_record_variable
func F_instantiate_empty_record_variable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exec_prepare_plan github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exec_prepare_plan
func F_exec_prepare_plan(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2
func F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dump_block github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_dump_block
func F_dump_block(m *base.Module, l0 int32)
//go:linkname F_plpgsql_scanner_errposition github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_plpgsql_scanner_errposition
func F_plpgsql_scanner_errposition(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_SN_create_env github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_SN_create_env
func F_SN_create_env(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_r_consonant_pair_1 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_r_consonant_pair_1
func F_r_consonant_pair_1(m *base.Module, l0 int32) int32
//go:linkname F_find_among_b github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_find_among_b
func F_find_among_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replace_s github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_replace_s
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_slice_from_s github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_slice_from_s
func F_slice_from_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_slice_del github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_slice_del
func F_slice_del(m *base.Module, l0 int32) int32
//go:linkname F_slice_to github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_slice_to
func F_slice_to(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_mbuf_append github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_mbuf_append
func F_mbuf_append(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pullf_create github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pullf_create
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pullf_free github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_free
func F_pullf_free(m *base.Module, l0 int32)
//go:linkname F_pullf_read github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_pullf_read
func F_pullf_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pullf_read_max github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_pullf_read_max
func F_pullf_read_max(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pullf_read_fixed github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pullf_read_fixed
func F_pullf_read_fixed(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_decompress_filter github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_decompress_filter
func F_pgp_decompress_filter(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pgp_parse_pkt_hdr github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_parse_pkt_hdr
func F_pgp_parse_pkt_hdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_mpi_check github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mpi_check
func F_mpi_check(m *base.Module, l0 int32) int32
//go:linkname F_pgp_mpi_alloc github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pgp_mpi_alloc
func F_pgp_mpi_alloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pgp_mpi_free github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_pgp_mpi_free
func F_pgp_mpi_free(m *base.Module, l0 int32) int32
//go:linkname F_pgp_load_digest github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_pgp_load_digest
func F_pgp_load_digest(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_px_debug github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_px_debug
func F_px_debug(m *base.Module, l0 int32, l1 int32)
//go:linkname F_citextcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_citextcmp
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_internal_citext_pattern_cmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_internal_citext_pattern_cmp
func F_internal_citext_pattern_cmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generate_trgm_only github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_generate_trgm_only
func F_generate_trgm_only(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hstoreUpgrade github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_hstoreUpgrade
func F_hstoreUpgrade(m *base.Module, l0 int32) int32
//go:linkname F_array_iterator github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_array_iterator
func F_array_iterator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ltree_crc32_sz github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_ltree_crc32_sz
func F_ltree_crc32_sz(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parse_ltree github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_parse_ltree
func F_parse_ltree(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_finish_nodeitem github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_finish_nodeitem
func F_finish_nodeitem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lca_inner github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_lca_inner
func F_lca_inner(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_queryin github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_queryin
func F_queryin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___memcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memset github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__Exit github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F__Exit
func F__Exit(m *base.Module, l0 int32)
//go:linkname F_access github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_access
func F_access(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_atan github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_atan
func F_atan(m *base.Module, l0 float64) float64
//go:linkname F___isspace github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___isspace
func F___isspace(m *base.Module, l0 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32)
//go:linkname F___cos github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___cos
func F___cos(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F___rem_pio2 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___rem_pio2
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F___sin github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F_memmove github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memmove
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___time github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___time
func F___time(m *base.Module) int64
//go:linkname F___gettimeofday github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___gettimeofday
func F___gettimeofday(m *base.Module, l0 int32)
//go:linkname F_exp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_exp
func F_exp(m *base.Module, l0 float64) float64
//go:linkname F_fclose github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fclose
func F_fclose(m *base.Module, l0 int32) int32
//go:linkname F_fflush github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F___uflow github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___uflow
func F___uflow(m *base.Module, l0 int32) int32
//go:linkname F_do_getc github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_do_getc
func F_do_getc(m *base.Module, l0 int32) int32
//go:linkname F_fgets github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fgets
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fopen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fopen
func F_fopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_fread
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___fstatat github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fsync github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F___ftello_unlocked github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___ftello_unlocked
func F___ftello_unlocked(m *base.Module, l0 int32) int64
//go:linkname F_fwrite github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getrusage github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_getrusage
func F_getrusage(m *base.Module, l0 int32)
//go:linkname F_isalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_iswalnum github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_iswalnum
func F_iswalnum(m *base.Module, l0 int32) int32
//go:linkname F_iswalpha github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_iswalpha
func F_iswalpha(m *base.Module, l0 int32) int32
//go:linkname F_iswprint github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_iswprint
func F_iswprint(m *base.Module, l0 int32) int32
//go:linkname F_iswspace github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_iswspace
func F_iswspace(m *base.Module, l0 int32) int32
//go:linkname F_kill github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_kill
func F_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ldexp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_ldexp
func F_ldexp(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_memchr github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mkdir github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_mkdir
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_posix_fadvise github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_posix_fadvise
func F_posix_fadvise(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32
//go:linkname F_pow github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pow
func F_pow(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F_pread github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_pread
func F_pread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_read github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_readdir github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_readdir
func F_readdir(m *base.Module, l0 int32) int32
//go:linkname F_rename github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_rename
func F_rename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sigemptyset github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32)
//go:linkname F_sigprocmask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_sigprocmask
func F_sigprocmask(m *base.Module, l0 int32, l1 int32)
//go:linkname F_stat github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_stat
func F_stat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strcpy
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strlcpy github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strlcpy
func F_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strlen github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F_strncmp github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncmp
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strspn github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_strspn
func F_strspn(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strtod github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtox_2 github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F___syscall_ret github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F_towlower github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_towlower
func F_towlower(m *base.Module, l0 int32) int32
//go:linkname F_casemap github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_casemap
func F_casemap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_towupper github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F_towupper
func F_towupper(m *base.Module, l0 int32) int32
//go:linkname F_umask github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_umask
func F_umask(m *base.Module, l0 int32) int32
//go:linkname F_unlink github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F_vfscanf github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_vfscanf
func F_vfscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsnprintf github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F_vsnprintf
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_write github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/pgmem/internal/aot/pgaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_emscripten_builtin_realloc github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F_emscripten_builtin_realloc
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___divtf3 github.com/shibukawa/pgmem/internal/aot/pgaot/p4.F___divtf3
func F___divtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___multf3 github.com/shibukawa/pgmem/internal/aot/pgaot/p3.F___multf3
func F___multf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___multi3 github.com/shibukawa/pgmem/internal/aot/pgaot/p0.F___multi3
func F___multi3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___udivmodti4 github.com/shibukawa/pgmem/internal/aot/pgaot/p5.F___udivmodti4
func F___udivmodti4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
