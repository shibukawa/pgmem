package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaGetBit(m *base.Module, l0 int32) int32 {
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
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v20 == int32(1) {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v26 == int32(18) {
				v29 = int32(16)
			} else {
				v29 = int32(0)
			}
			if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v36 = int32(4)
			} else {
				v36 = v29
			}
			v49 = v36
		} else {
			v37 = int32(1)
			if v20&v37 != 0 {
				v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v52 = base.I64_extend_i32_s(v49) << (uint(int64(3)) % 64)
		if base.B2i32(int64(0) <= v17)&base.B2i32(v17 < v52) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v52 - int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
					F_errmsg(m, int32(_a_F_byteaGetBit_0), v9)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaGetBit_1), int32(3356), int32(_a_F_byteaGetBit_2))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v79 = int32(1)
			if v20&v79 != 0 {
				v83 = v79
			} else {
				v83 = int32(4)
			}
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(3))%64)))+(v12+v83)))))
			m.G0 = v9 + int32(16)
			return int32(base.Ui32(v86)>>(uint(base.I32_wrap_i64(v17)&int32(7))%32)) & int32(1)
		}
	}
}
func F_bytea_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	if int32(0) < l2 {
		v9 = l2 + l3
		if base.B2i32(l3 < int32(0)) != base.B2i32(v9 < l2) {
			F_errstart_cold(m, int32(21), int32(0))
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bytea_overlay_0), int32(0))
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bytea_overlay_1), int32(3240), int32(_a_F_bytea_overlay_2))
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v12 = int32(1)
			v15 = F_bytea_substring(m, l0, v12, l2-v12)
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = int32(1)
				if v9 <= v19 {
					v22 = v19
				} else {
					v22 = v9
				}
				v26 = F_detoast_attr_slice(m, l0, v22-int32(1), int32(-1))
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = F_text_catenate(m, v15, l1)
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = F_text_catenate(m, v28, v26)
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							return v30
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_bytea_overlay_3), int32(0))
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bytea_overlay_1), int32(3236), int32(_a_F_bytea_overlay_2))
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_bytea_reverse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v14 == v12 {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v22 == int32(18) {
				v25 = int32(16)
			} else {
				v25 = int32(0)
			}
			if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v32 = int32(4)
			} else {
				v32 = v25
			}
			v43 = v32
		} else {
			v33 = int32(1)
			if v16 != 0 {
				v43 = int32(base.Ui32(v14)>>(uint(v33)%32)) - v33
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v45 = v43 + int32(4)
		v46 = F_palloc(m, v45)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v46))) = v45 << (uint(int32(2)) % 32)
			if int32(0) < v43 {
				if v16 != 0 {
					v55 = v13
				} else {
					v55 = v8 + int32(4)
				}
				v60 = v55
				v61 = v43 + v46 + int32(4)
				for {
					v66 = int32(1)
					v67 = v61 - v66
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
					*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
					v71 = v60 + v66
					if base.Ui32(v71) < base.Ui32(v55+v43) {
						v60 = v71
						v61 = v67
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return v46
		}
	}
}
func F_bytea_string_agg_transfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v14 == v2 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v18 = v17
	} else {
		v18 = v2
	}
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v19 != 0 {
		v187 = v18
		if v187 == int32(0) {
			v195 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
			v198 = int32(0)
		} else {
			v198 = v187
		}
		m.G0 = v12 + int32(16)
		return v198
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v28 = v12 + int32(12)
				v29 = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v30 == v29 {
					v47 = int32(0)
					if v28 == v47 {
						v55 = v47
					} else {
						v50 = v47
						v51 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
						v55 = v51
					}
					v58 = v55
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					switch v33 - int32(429) {
					case 0:
						if v28 == int32(0) {
							v58 = int32(1)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
							v50 = v40
							v51 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
							v55 = v51
							v58 = v55
						}
					case 1:
						if v28 == int32(0) {
							v58 = int32(2)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+368))
							v50 = v45
							v51 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
							v55 = v51
							v58 = v55
						}
					default:
						v47 = int32(0)
						if v28 == v47 {
							v55 = v47
						} else {
							v50 = v47
							v51 = v29
							*(*int32)(unsafe.Add(mBase, uint32(v28))) = v50
							v55 = v51
						}
						v58 = v55
					}
				}
				if v58 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_bytea_string_agg_transfn_0), int32(0))
						mBase = m.M
						v210 = m.ExcPending
						if v210 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bytea_string_agg_transfn_1), int32(_a_F_bytea_string_agg_transfn_2), int32(_a_F_bytea_string_agg_transfn_3))
							mBase = m.M
							v215 = m.ExcPending
							if v215 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v61 = int32(_a_F_bytea_string_agg_transfn_4)
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_bytea_string_agg_transfn[0]))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					*(*int32)(unsafe.Add(mBase, _c_F_bytea_string_agg_transfn[0])) = v64
					v66 = F_makeStringInfo(m)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_bytea_string_agg_transfn[0])) = v62
						v70 = v66
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
						if v72 != 0 {
							v149 = int32(1)
							v150 = v21 + v149
							v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							v155 = v153 & v149
							if v155 != 0 {
								v156 = v150
							} else {
								v156 = v21 + int32(4)
							}
							if v153 == int32(1) {
								v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
								if v162 == int32(18) {
									v165 = int32(16)
								} else {
									v165 = int32(0)
								}
								if base.Ui32((v162-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v172 = int32(4)
								} else {
									v172 = v165
								}
								v183 = v172
							} else {
								v173 = int32(1)
								if v155 != 0 {
									v183 = int32(base.Ui32(v153)>>(uint(v173)%32)) - v173
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v183 = int32(base.Ui32(v177)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, v70, v156, v183)
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int32(0)
							} else {
								v187 = v70
								if v187 == int32(0) {
									v195 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
									v198 = int32(0)
								} else {
									v198 = v187
								}
								m.G0 = v12 + int32(16)
								return v198
							}
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v74 = F_pg_detoast_datum_packed(m, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v76 = int32(1)
								v77 = v74 + v76
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								v82 = v80 & v76
								if v82 != 0 {
									v83 = v77
								} else {
									v83 = v74 + int32(4)
								}
								if v80 == int32(1) {
									v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
									if v89 == int32(18) {
										v92 = int32(16)
									} else {
										v92 = int32(0)
									}
									if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v99 = int32(4)
									} else {
										v99 = v92
									}
									v110 = v99
								} else {
									v100 = int32(1)
									if v82 != 0 {
										v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								F_appendBinaryStringInfo(m, v70, v83, v110)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									if v18 != 0 {
									} else {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
										if v113 == int32(1) {
											v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
											if v119 == int32(18) {
												v122 = int32(16)
											} else {
												v122 = int32(0)
											}
											if base.Ui32((v119-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v129 = int32(4)
											} else {
												v129 = v122
											}
											v142 = v129
										} else {
											v130 = int32(1)
											if v113&v130 != 0 {
												v142 = int32(base.Ui32(v113)>>(uint(v130)%32)) - v130
											} else {
												v136 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
												v142 = int32(base.Ui32(v136)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v142
									}
									v149 = int32(1)
									v150 = v21 + v149
									v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									v155 = v153 & v149
									if v155 != 0 {
										v156 = v150
									} else {
										v156 = v21 + int32(4)
									}
									if v153 == int32(1) {
										v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
										if v162 == int32(18) {
											v165 = int32(16)
										} else {
											v165 = int32(0)
										}
										if base.Ui32((v162-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v172 = int32(4)
										} else {
											v172 = v165
										}
										v183 = v172
									} else {
										v173 = int32(1)
										if v155 != 0 {
											v183 = int32(base.Ui32(v153)>>(uint(v173)%32)) - v173
										} else {
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
											v183 = int32(base.Ui32(v177)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									F_appendBinaryStringInfo(m, v70, v156, v183)
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v187 = v70
										if v187 == int32(0) {
											v195 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
											v198 = int32(0)
										} else {
											v198 = v187
										}
										m.G0 = v12 + int32(16)
										return v198
									}
								}
							}
						}
					}
				}
			} else {
				v70 = v18
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v72 != 0 {
					v149 = int32(1)
					v150 = v21 + v149
					v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					v155 = v153 & v149
					if v155 != 0 {
						v156 = v150
					} else {
						v156 = v21 + int32(4)
					}
					if v153 == int32(1) {
						v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
						if v162 == int32(18) {
							v165 = int32(16)
						} else {
							v165 = int32(0)
						}
						if base.Ui32((v162-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v172 = int32(4)
						} else {
							v172 = v165
						}
						v183 = v172
					} else {
						v173 = int32(1)
						if v155 != 0 {
							v183 = int32(base.Ui32(v153)>>(uint(v173)%32)) - v173
						} else {
							v177 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v183 = int32(base.Ui32(v177)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_appendBinaryStringInfo(m, v70, v156, v183)
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						v187 = v70
						if v187 == int32(0) {
							v195 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
							v198 = int32(0)
						} else {
							v198 = v187
						}
						m.G0 = v12 + int32(16)
						return v198
					}
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v74 = F_pg_detoast_datum_packed(m, v73)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = int32(1)
						v77 = v74 + v76
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
						v82 = v80 & v76
						if v82 != 0 {
							v83 = v77
						} else {
							v83 = v74 + int32(4)
						}
						if v80 == int32(1) {
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
							if v89 == int32(18) {
								v92 = int32(16)
							} else {
								v92 = int32(0)
							}
							if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v99 = int32(4)
							} else {
								v99 = v92
							}
							v110 = v99
						} else {
							v100 = int32(1)
							if v82 != 0 {
								v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
								v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						F_appendBinaryStringInfo(m, v70, v83, v110)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							if v18 != 0 {
							} else {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								if v113 == int32(1) {
									v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
									if v119 == int32(18) {
										v122 = int32(16)
									} else {
										v122 = int32(0)
									}
									if base.Ui32((v119-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v129 = int32(4)
									} else {
										v129 = v122
									}
									v142 = v129
								} else {
									v130 = int32(1)
									if v113&v130 != 0 {
										v142 = int32(base.Ui32(v113)>>(uint(v130)%32)) - v130
									} else {
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v142 = int32(base.Ui32(v136)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v142
							}
							v149 = int32(1)
							v150 = v21 + v149
							v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							v155 = v153 & v149
							if v155 != 0 {
								v156 = v150
							} else {
								v156 = v21 + int32(4)
							}
							if v153 == int32(1) {
								v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
								if v162 == int32(18) {
									v165 = int32(16)
								} else {
									v165 = int32(0)
								}
								if base.Ui32((v162-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v172 = int32(4)
								} else {
									v172 = v165
								}
								v183 = v172
							} else {
								v173 = int32(1)
								if v155 != 0 {
									v183 = int32(base.Ui32(v153)>>(uint(v173)%32)) - v173
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v183 = int32(base.Ui32(v177)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, v70, v156, v183)
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int32(0)
							} else {
								v187 = v70
								if v187 == int32(0) {
									v195 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v195)
									v198 = int32(0)
								} else {
									v198 = v187
								}
								m.G0 = v12 + int32(16)
								return v198
							}
						}
					}
				}
			}
		}
	}
}
