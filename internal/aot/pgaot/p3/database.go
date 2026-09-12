package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetDatabaseHasLoginEventTriggers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[130]))
		F_LockSharedObject(m, int32(1262), v15, int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[130]))
			v22 = F_SearchSysCacheLockedCopy1(m, int32(21), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 != 0 {
					v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+8)))
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v24)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
					v30 = v28 + v29
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+79)))
					if v31 == int32(0) {
						v34 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v30)+79)) = uint8(v34)
						F_CatalogTupleUpdate(m, v11, v7+int32(8), v22)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_CommandCounterIncrement(m)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								F_UnlockTuple(m, v11, v7+int32(8), int32(7))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									F_sequence_close(m, v11, int32(3))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							}
						}
					} else {
						F_UnlockTuple(m, v11, v7+int32(8), int32(7))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							F_sequence_close(m, v11, int32(3))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_pfree(m, v22)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, _consts[130]))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v60
						F_errmsg_internal(m, int32(54369), v7)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errfinish(m, int32(521275), int32(409), int32(143708))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
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
}
