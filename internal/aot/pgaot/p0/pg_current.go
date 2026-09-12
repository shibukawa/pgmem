package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_snapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v38 int64
	_ = v38
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int64
	_ = v102
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v11 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[40]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			v23 = F_palloc(m, v18<<(uint(int32(3))%32)+int32(24))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if base.Ui32(v25) <= base.Ui32(int32(2)) {
					v43 = base.I64_extend_i32_u(v25)
				} else {
					v30 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v25) {
						v38 = (v30 - int64(1)) & int64(4294967295)
					} else {
						v38 = v30
					}
					v43 = base.I64_extend_i32_u(v25) | v38<<(uint(int64(32))%64)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v43
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				if base.Ui32(v45) <= base.Ui32(int32(2)) {
					v63 = base.I64_extend_i32_u(v45)
				} else {
					v50 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v45) {
						v58 = (v50 - int64(1)) & int64(4294967295)
					} else {
						v58 = v50
					}
					v63 = base.I64_extend_i32_u(v45) | v58<<(uint(int64(32))%64)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v18
				*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v63
				if v18 == int32(0) {
					v172 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v172<<(uint(int32(5))%32) + int32(96)
					return v23
				} else {
					v72 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					v80 = int32(0)
					for {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						v93 = int32(2)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v80<<(uint(v93)%32))))
						if base.Ui32(v96) <= base.Ui32(v93) {
							v107 = base.I64_extend_i32_u(v96)
						} else {
							if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v96) {
								v102 = (v72 - int64(1)) & int64(4294967295)
							} else {
								v102 = v72
							}
							v107 = base.I64_extend_i32_u(v96) | v102<<(uint(int64(32))%64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(24)+v80<<(uint(int32(3))%32)))) = v107
						v110 = v80 + int32(1)
						if v110 != v18 {
							v80 = v110
							continue
						} else {
							break
						}
						break
					}
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					if base.Ui32(v112) < base.Ui32(int32(2)) {
						v172 = v112
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = v172<<(uint(int32(5))%32) + int32(96)
						return v23
					} else {
						v116 = v23 + int32(24)
						F_pg_qsort(m, v116, v112, int32(8), int32(1562))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if base.Ui32(int32(2)) <= base.Ui32(v121) {
								v126 = int32(0)
								v127 = int32(1)
								for {
									v136 = int32(3)
									v139 = *(*int64)(unsafe.Add(mBase, uint32(v116+v127<<(uint(v136)%32))))
									v143 = *(*int64)(unsafe.Add(mBase, uint32(v116+v126<<(uint(v136)%32))))
									if v139 == v143 {
										v152 = v126
									} else {
										v146 = v126 + int32(1)
										if v127 == v146 {
											v152 = v127
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v116+v146<<(uint(int32(3))%32)))) = v139
											v152 = v146
										}
									}
									v155 = v127 + int32(1)
									if v155 != v121 {
										v126 = v152
										v127 = v155
										continue
									} else {
										break
									}
									break
								}
								v161 = v152 + int32(1)
							} else {
								v161 = v121
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v161
							v172 = v161
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = v172<<(uint(int32(5))%32) + int32(96)
							return v23
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(106108), int32(0))
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493183), int32(380), int32(86201))
					mBase = m.M
					v198 = m.ExcPending
					if v198 != 0 {
						return int32(0)
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
func F_pg_current_wal_insert_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(127733), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(556132), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493093), int32(303), int32(243964))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
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
		v37 = F_GetXLogInsertRecPtr(m)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = F_Int64GetDatum(m, v37)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		}
	}
}
func F_pg_current_xact_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	F_PreventCommandDuringRecovery(m, int32(667384))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, _consts[35]))
		if base.I32_wrap_i64(v9) != 0 {
			v16 = v9
			v17 = F_Int64GetDatum(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		} else {
			F_AssignTransactionId(m, int32(4384296))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v15 = *(*int64)(unsafe.Add(mBase, _consts[35]))
				v16 = v15
				v17 = F_Int64GetDatum(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
