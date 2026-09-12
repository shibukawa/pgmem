package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IvfflatInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v3 = F_add_reloption_kind(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[1456])) = v3
		F_add_int_reloption(m, v3, int32(125198), int32(125179), int32(100), int32(1), int32(32768))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_DefineCustomIntVariable(m, int32(181564), int32(181579), int32(609544), int32(4718564), int32(1), int32(32768))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_DefineCustomEnumVariable(m, int32(297307), int32(159033), int32(0), int32(4718568), int32(4116224), int32(6))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v33 = int32(32768)
					F_DefineCustomIntVariable(m, int32(181545), int32(158926), int32(0), int32(4718572), v33, v33)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_MarkGUCPrefixReserved(m, int32(119550))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_IvfflatParallelBuildMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v13 = F_shm_toc_lookup(m, l1, int64(-6917529027641081852), int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[30])) = v13
		F_pgstat_report_activity(m, int32(3), v13)
		mBase = m.M
		v20 = F_shm_toc_lookup(m, l1, int64(-6917529027641081855), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
			if v25 != 0 {
				v26 = int32(4)
			} else {
				v26 = int32(5)
			}
			v27 = F_table_open(m, v22, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				if v25 != 0 {
					v32 = int32(3)
				} else {
					v32 = int32(8)
				}
				v33 = F_index_open(m, v29, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = F_palloc0(m, int32(12))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v27
						v42 = F_shm_toc_lookup(m, l1, int64(-6917529027641081854), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_tuplesort_attach_shared(m, v42, l0)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v48 = F_shm_toc_lookup(m, l1, int64(-6917529027641081853), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, _consts[49]))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
									v53 = base.I32_div_s(v51, v52)
									F_IvfflatParallelScanAndSort(m, v36, v20, v42, v48, v53, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										F_relation_close(m, v33, v32)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											F_sequence_close(m, v27, v26)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ivfflat_bit_support(m *base.Module, l0 int32) int32 {
	return int32(4116352)
}
func F_ivfflat_halfvec_support(m *base.Module, l0 int32) int32 {
	return int32(4116328)
}
