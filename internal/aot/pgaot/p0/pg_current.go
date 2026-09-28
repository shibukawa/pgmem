package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_snapshot(m *base.Module, l0 int32) int64 {
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
	var v31 int64
	_ = v31
	var v38 int64
	_ = v38
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v51 int64
	_ = v51
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
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v11 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_snapshot[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			v23 = F_palloc(m, v18<<(uint(int32(3))%32)+int32(24))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if base.Ui32(v25) <= base.Ui32(int32(2)) {
					v43 = base.I64_extend_i32_u(v25)
				} else {
					v31 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v25) {
						v38 = (v31 - int64(1)) & int64(4294967295)
					} else {
						v38 = v31
					}
					v43 = base.I64_extend_i32_u(v25) | v38<<(uint(int64(32))%64)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v43
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				if base.Ui32(v45) <= base.Ui32(int32(2)) {
					v63 = base.I64_extend_i32_u(v45)
				} else {
					v51 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v45) {
						v58 = (v51 - int64(1)) & int64(4294967295)
					} else {
						v58 = v51
					}
					v63 = base.I64_extend_i32_u(v45) | v58<<(uint(int64(32))%64)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v18
				*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v63
				if v18 == int32(0) {
					v183 = int32(96)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v183
					return base.I64_extend_i32_u(v23)
				} else {
					v72 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					v80 = int32(0)
					for {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						v93 = int32(2)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v80<<(uint(v93)%32))))
						if base.Ui32(v96) <= base.Ui32(v93) {
							v106 = base.I64_extend_i32_u(v96)
						} else {
							if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v96) {
								v102 = (v72 - int64(1)) & int64(4294967295)
							} else {
								v102 = v72
							}
							v106 = base.I64_extend_i32_u(v96) | v102<<(uint(int64(32))%64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(24)+v80<<(uint(int32(3))%32)))) = v106
						v109 = v80 + int32(1)
						if v109 != v18 {
							v80 = v109
							continue
						} else {
							break
						}
						break
					}
					v111 = int32(1)
					if v18 == v111 {
						v183 = int32(128)
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = v183
						return base.I64_extend_i32_u(v23)
					} else {
						v116 = v23 + int32(24)
						F_pg_qsort(m, v116, v18, int32(8), int32(1759))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int64(0)
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if base.Ui32(int32(2)) <= base.Ui32(v121) {
								v125 = int32(0)
								v126 = v111
								for {
									v135 = int32(3)
									v138 = *(*int64)(unsafe.Add(mBase, uint32(v116+v126<<(uint(v135)%32))))
									v142 = *(*int64)(unsafe.Add(mBase, uint32(v116+v125<<(uint(v135)%32))))
									if v138 == v142 {
										v151 = v125
									} else {
										v145 = v125 + int32(1)
										if v126 == v145 {
											v151 = v126
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v116+v145<<(uint(int32(3))%32)))) = v138
											v151 = v145
										}
									}
									v154 = v126 + int32(1)
									if v154 != v121 {
										v125 = v151
										v126 = v154
										continue
									} else {
										break
									}
									break
								}
								v162 = v151 + int32(1)
							} else {
								v162 = v121
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v162
							v183 = v162<<(uint(int32(5))%32) + int32(96)
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = v183
							return base.I64_extend_i32_u(v23)
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v190 = m.ExcPending
			if v190 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_current_snapshot_0), int32(0))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_current_snapshot_1), int32(381), int32(_a_F_pg_current_snapshot_2))
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return int64(0)
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
func F_pg_current_wal_insert_lsn(m *base.Module, l0 int32) int64 {
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_insert_lsn[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_insert_lsn[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_insert_lsn[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_current_wal_insert_lsn_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_current_wal_insert_lsn_1), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_current_wal_insert_lsn_2), int32(333), int32(_a_F_pg_current_wal_insert_lsn_3))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
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
			return int64(0)
		} else {
			return v37
		}
	}
}
func F_pg_current_xact_id(m *base.Module, l0 int32) int64 {
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
	F_PreventCommandDuringRecovery(m, int32(_a_F_pg_current_xact_id_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, _c_F_pg_current_xact_id[0]))
		if base.I32_wrap_i64(v9) != 0 {
			v16 = v9
			return v16
		} else {
			F_AssignTransactionId(m, int32(_a_F_pg_current_xact_id_1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				v15 = *(*int64)(unsafe.Add(mBase, _c_F_pg_current_xact_id[0]))
				v16 = v15
				return v16
			}
		}
	}
}
