package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildQueryCompletionString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = v6 << (uint(int32(3)) % 32)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_BuildQueryCompletionString[0]))))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_BuildQueryCompletionString[1])))
		base.MemoryCopy(m, l0, v10, v9)
	} else {
	}
	v12 = l0 + v9
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F_BuildQueryCompletionString[2]))))
	if v13&int32(1) != 0 {
		if v6 == int32(158) {
			v18 = int32(_a_F_BuildQueryCompletionString_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v12))) = uint16(v18)
			v22 = v12 + int32(2)
		} else {
			v22 = v12
		}
		v23 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v23)
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		v27 = v22 + int32(1)
		v28 = int32(0)
		if v25 == int64(0) {
			v37 = int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v37)
			v208 = int32(1)
		} else {
			v44 = int32(1233)
			v49 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v25))^int32(63))*v44+v44) >> (uint(int32(12)) % 32))
			v52 = *(*int64)(unsafe.Add(mBase, uint32(v49<<(uint(int32(3))%32))+uint32(_c_F_BuildQueryCompletionString[3])))
			v54 = v49 + base.B2i32(base.Ui64(v52) <= base.Ui64(v25))
			if base.Ui64(int64(100000000)) <= base.Ui64(v25) {
				v58 = v25
				v61 = v28
				for {
					v67 = v27 + v54 - v61
					v68 = int32(8)
					v71 = base.I64_div_u_s(v58, int64(100000000))
					v75 = base.I32_wrap_i64(v58 + v71*int64(4194967296))
					v77 = base.I32_div_u_s(v75, int32(_a_F_BuildQueryCompletionString_1))
					v78 = int32(1)
					v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77<<(uint(v78)%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v67-v68))) = uint16(v80)
					v84 = int32(_a_F_BuildQueryCompletionString_2)
					v85 = base.I32_div_u_s(v75, v84)
					v86 = int32(100)
					v87 = base.I32_rem_u_s(v85, v86)
					v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87<<(uint(v78)%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v67-int32(6)))) = uint16(v90)
					v96 = v75 - v85*v84
					v97 = int32(_a_F_BuildQueryCompletionString_3)
					v100 = base.I32_div_u_s(v96&v97, v86)
					v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100<<(uint(v78)%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v67-int32(4)))) = uint16(v103)
					v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v96-v100*v86)&v97<<(uint(v78)%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v67-int32(2)))) = uint16(v114)
					v117 = v61 + v68
					if base.Ui64(int64(9999999999999999)) < base.Ui64(v58) {
						v58 = v71
						v61 = v117
						continue
					} else {
						break
					}
					break
				}
				v120 = v71
				v123 = v117
			} else {
				v120 = v25
				v123 = v28
			}
			v129 = base.I32_wrap_i64(v120)
			if base.Ui64(int64(10000)) <= base.Ui64(v120) {
				v133 = v27 + v54 - v123
				v134 = int32(4)
				v137 = base.I32_div_u_s(v129, int32(_a_F_BuildQueryCompletionString_2))
				v140 = v129 + v137*int32(-10000)
				v141 = int32(100)
				v142 = base.I32_div_u_s(v140, v141)
				v143 = int32(1)
				v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142<<(uint(v143)%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v133-v134))) = uint16(v145)
				v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v140-v142*v141)<<(uint(v143)%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v133-int32(2)))) = uint16(v154)
				v158 = v137
				v159 = v123 | v134
			} else {
				v158 = v129
				v159 = v123
			}
			if base.Ui32(int32(100)) <= base.Ui32(v158) {
				v167 = int32(2)
				v169 = int32(_a_F_BuildQueryCompletionString_3)
				v171 = int32(100)
				v172 = base.I32_div_u_s(v158&v169, v171)
				v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v158-v172*v171)&v169<<(uint(int32(1))%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v27+v54-v159-v167))) = uint16(v180)
				v184 = v172
				v185 = v159 + v167
			} else {
				v184 = v158
				v185 = v159
			}
			if base.Ui32(int32(10)) <= base.Ui32(v184) {
				v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184<<(uint(int32(1))%32))+uint32(_c_F_BuildQueryCompletionString[4]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v27+v54-v185-int32(2)))) = uint16(v194)
				v208 = v54
			} else {
				v197 = v184 | int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v197)
				v208 = v54
			}
		}
		v211 = v208 + v27
	} else {
		v211 = v12
	}
	v212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v212)
	return v211 - l0
}
func F_query_contains_extern_params_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 != int32(67) {
			if v8 != int32(8) {
				v25 = F_expression_tree_walker_impl(m, l0, int32(531), l1)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					return v25
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				return base.B2i32(v13 == int32(0))
			}
		} else {
			v19 = F_query_tree_walker_impl(m, l0, int32(531), l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	}
}
func F_query_to_xmlschema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = F_text_to_cstring(m, v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v20 = F_pg_detoast_datum_packed(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = F_text_to_cstring(m, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = int32(0)
						v29 = F_SPI_prepare(m, v16, v27, v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							if v29 != 0 {
								v31 = F_SPI_cursor_open(m, v29)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int64(0)
								} else {
									if v31 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v16
											F_errmsg_internal(m, int32(_a_F_query_to_xmlschema_0), v9+int32(16))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_query_to_xmlschema_1), int32(3123), int32(_a_F_query_to_xmlschema_2))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+92))
										v39 = F_map_sql_table_to_xmlschema(m, v35, int32(0), base.B2i32(v18 != int64(0)), v22)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int64(0)
										} else {
											v41 = F_strlen(m, v39)
											mBase = m.M
											v43 = v41 + int32(1)
											v44 = F_SPI_palloc(m, v43)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int64(0)
											} else {
												if v43 != 0 {
													base.MemoryCopy(m, v44, v39, v43)
												} else {
												}
												F_SPI_cursor_close(m, v31)
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return int64(0)
												} else {
													v49 = F_SPI_finish(m)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int64(0)
													} else {
														v51 = F_cstring_to_text(m, v44)
														mBase = m.M
														v52 = m.ExcPending
														if v52 != 0 {
															return int64(0)
														} else {
															m.G0 = v9 + int32(32)
															return base.I64_extend_i32_u(v51)
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
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
									F_errmsg_internal(m, int32(_a_F_query_to_xmlschema_3), v9)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_query_to_xmlschema_1), int32(3120), int32(_a_F_query_to_xmlschema_2))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
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
				}
			}
		}
	}
}
