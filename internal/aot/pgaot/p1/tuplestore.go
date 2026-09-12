package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplestore_putvalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	v6 = int32(4548768)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v9
	v12 = F_heap_form_minimal_tuple(m, l1, l2, l3, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = F_GetMemoryChunkSpace(m, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v16 - base.I64_extend_i32_u(v14)
			F_tuplestore_puttuple_common(m, l0, v12)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
				return
			}
		}
	}
}
func F_tuplestore_select_read_pointer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v6 != l1 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v8) < base.Ui32(int32(2)) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = l1
			return
		} else {
			if v8 == int32(2) {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v14 = int32(24)
				v16 = v13 + l1*v14
				v19 = v13 + v6*v14
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)))
				if v20 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v19+int32(12)))) = v28
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
					v31 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+40)))
					*(*int64)(unsafe.Add(mBase, uint32(v19+int32(16)))) = v30 + v31
				} else {
				}
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
				if v35 == int32(1) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
					v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
					v41 = F_BufFileSeek(m, v34, v38, v39, int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = l1
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									F_errmsg(m, int32(402645), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										F_errfinish(m, int32(518848), int32(552), int32(224637))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					v61 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
					v63 = F_BufFileSeek(m, v34, v60, v61, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						if v63 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = l1
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errmsg(m, int32(402645), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_errfinish(m, int32(518848), int32(562), int32(224637))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(367405), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						F_errfinish(m, int32(518848), int32(566), int32(224637))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
