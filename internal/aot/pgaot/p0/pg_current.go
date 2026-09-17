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
	var v74 int64
	_ = v74
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v11 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_snapshot[0]))
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
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(96)
					return v23
				} else {
					v74 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
					v82 = int32(0)
					for {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						v95 = int32(2)
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v82<<(uint(v95)%32))))
						if base.Ui32(v98) <= base.Ui32(v95) {
							v108 = base.I64_extend_i32_u(v98)
						} else {
							if base.Ui32(base.I32_wrap_i64(v11)) < base.Ui32(v98) {
								v104 = (v74 - int64(1)) & int64(4294967295)
							} else {
								v104 = v74
							}
							v108 = base.I64_extend_i32_u(v98) | v104<<(uint(int64(32))%64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(24)+v82<<(uint(int32(3))%32)))) = v108
						v111 = v82 + int32(1)
						if v111 != v18 {
							v82 = v111
							continue
						} else {
							break
						}
						break
					}
					v113 = int32(1)
					if v18 == v113 {
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(128)
						return v23
					} else {
						v120 = v23 + int32(24)
						F_pg_qsort(m, v120, v18, int32(8), int32(1546))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if base.Ui32(int32(2)) <= base.Ui32(v125) {
								v129 = int32(0)
								v130 = v113
								for {
									v139 = int32(3)
									v142 = *(*int64)(unsafe.Add(mBase, uint32(v120+v130<<(uint(v139)%32))))
									v146 = *(*int64)(unsafe.Add(mBase, uint32(v120+v129<<(uint(v139)%32))))
									if v142 == v146 {
										v155 = v129
									} else {
										v149 = v129 + int32(1)
										if v130 == v149 {
											v155 = v130
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v120+v149<<(uint(int32(3))%32)))) = v142
											v155 = v149
										}
									}
									v158 = v130 + int32(1)
									if v158 != v125 {
										v129 = v155
										v130 = v158
										continue
									} else {
										break
									}
									break
								}
								v166 = v155 + int32(1)
							} else {
								v166 = v125
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v166
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = v166<<(uint(int32(5))%32) + int32(96)
							return v23
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v182 = m.ExcPending
			if v182 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_current_snapshot_0), int32(0))
				mBase = m.M
				v186 = m.ExcPending
				if v186 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_current_snapshot_1), int32(380), int32(_a_F_pg_current_snapshot_2))
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_insert_lsn[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_insert_lsn[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
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
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_current_wal_insert_lsn_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_pg_current_wal_insert_lsn_1), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_current_wal_insert_lsn_2), int32(303), int32(_a_F_pg_current_wal_insert_lsn_3))
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
	F_PreventCommandDuringRecovery(m, int32(_a_F_pg_current_xact_id_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, _c_F_pg_current_xact_id[0]))
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
			F_AssignTransactionId(m, int32(_a_F_pg_current_xact_id_1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v15 = *(*int64)(unsafe.Add(mBase, _c_F_pg_current_xact_id[0]))
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
