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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
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
			v23 = int32(4)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v25&int32(254) == int32(2) {
				v34 = v23
			} else {
				v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
			}
			if v25 == int32(1) {
				v37 = v23
			} else {
				v37 = v34
			}
			v50 = v37
		} else {
			v38 = int32(1)
			if v20&v38 != 0 {
				v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v53 = base.I64_extend_i32_s(v50) << (uint(int64(3)) % 64)
		if base.B2i32(int64(0) <= v17)&base.B2i32(v17 < v53) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v53 - int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
					F_errmsg(m, int32(427451), v9)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496954), int32(3356), int32(103134))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
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
			v80 = int32(1)
			if v20&v80 != 0 {
				v84 = v80
			} else {
				v84 = int32(4)
			}
			v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(3))%64)))+(v12+v84)))))
			m.G0 = v9 + int32(16)
			return int32(base.Ui32(v87)>>(uint(base.I32_wrap_i64(v17)&int32(7))%32)) & int32(1)
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
					F_errmsg(m, int32(399961), int32(0))
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496954), int32(3240), int32(26387))
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
				v26 = F_pg_detoast_datum_slice(m, l0, v22-int32(1), int32(-1))
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
				F_errmsg(m, int32(437039), int32(0))
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496954), int32(3236), int32(26387))
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(1)
		v14 = v9 + v13
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		v19 = v17 & v13
		if v19 != 0 {
			v20 = v14
		} else {
			v20 = v9 + int32(4)
		}
		if v17 == int32(1) {
			v23 = int32(4)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v25&int32(254) == int32(2) {
				v34 = v23
			} else {
				v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
			}
			if v25 == int32(1) {
				v37 = v23
			} else {
				v37 = v34
			}
			v48 = v37
		} else {
			v38 = int32(1)
			if v19 != 0 {
				v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v50 = v48 + int32(4)
		v51 = F_palloc(m, v50)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v51))) = v50 << (uint(int32(2)) % 32)
			v56 = v20 + v48
			if base.Ui32(v56) <= base.Ui32(v20) {
			} else {
				v62 = v48 + v51 + int32(4)
				v64 = v48 & int32(7)
				if v64 != 0 {
					v66 = v20
					v67 = v62
					v68 = int32(0)
					for {
						v73 = int32(1)
						v74 = v67 - v73
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
						*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v75)
						v78 = v66 + v73
						v80 = v68 + v73
						if v80 != v64 {
							v66 = v78
							v67 = v74
							v68 = v80
							continue
						} else {
							break
						}
						break
					}
					v82 = v78
					v83 = v74
				} else {
					v82 = v20
					v83 = v62
				}
				if base.Ui32(v48-int32(1)) < base.Ui32(int32(7)) {
				} else {
					v91 = v82
					v92 = v83
					for {
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(1)))) = uint8(v100)
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(2)))) = uint8(v104)
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(3)))) = uint8(v108)
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(4)))) = uint8(v112)
						v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(5)))) = uint8(v116)
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(6)))) = uint8(v120)
						v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+6)))
						*(*uint8)(unsafe.Add(mBase, uint32(v92-int32(7)))) = uint8(v124)
						v126 = int32(8)
						v127 = v92 - v126
						v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+7)))
						*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v128)
						v131 = v91 + v126
						if v131 != v56 {
							v91 = v131
							v92 = v127
							continue
						} else {
							break
						}
						break
					}
				}
			}
			return v51
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
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
		v190 = v18
		if v190 == int32(0) {
			v198 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v198)
			v201 = int32(0)
		} else {
			v201 = v190
		}
		m.G0 = v12 + int32(16)
		return v201
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
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(61279), int32(0))
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496954), int32(5432), int32(351794))
							mBase = m.M
							v218 = m.ExcPending
							if v218 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v61 = int32(4487040)
					v62 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v64
					v66 = F_makeStringInfo(m)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v62
						v70 = v66
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
						if v72 != 0 {
							v151 = int32(1)
							v152 = v21 + v151
							v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							v157 = v155 & v151
							if v157 != 0 {
								v158 = v152
							} else {
								v158 = v21 + int32(4)
							}
							if v155 == int32(1) {
								v161 = int32(4)
								v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
								if v163&int32(254) == int32(2) {
									v172 = v161
								} else {
									v172 = base.B2i32(v163 == int32(18)) << (uint(v161) % 32)
								}
								if v163 == int32(1) {
									v175 = v161
								} else {
									v175 = v172
								}
								v186 = v175
							} else {
								v176 = int32(1)
								if v157 != 0 {
									v186 = int32(base.Ui32(v155)>>(uint(v176)%32)) - v176
								} else {
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v186 = int32(base.Ui32(v180)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, v70, v158, v186)
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int32(0)
							} else {
								v190 = v70
								if v190 == int32(0) {
									v198 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v198)
									v201 = int32(0)
								} else {
									v201 = v190
								}
								m.G0 = v12 + int32(16)
								return v201
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
									v86 = int32(4)
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
									if v88&int32(254) == int32(2) {
										v97 = v86
									} else {
										v97 = base.B2i32(v88 == int32(18)) << (uint(v86) % 32)
									}
									if v88 == int32(1) {
										v100 = v86
									} else {
										v100 = v97
									}
									v111 = v100
								} else {
									v101 = int32(1)
									if v82 != 0 {
										v111 = int32(base.Ui32(v80)>>(uint(v101)%32)) - v101
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								F_appendBinaryStringInfo(m, v70, v83, v111)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									if v18 != 0 {
									} else {
										v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
										if v114 == int32(1) {
											v117 = int32(4)
											v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
											if v119&int32(254) == int32(2) {
												v128 = v117
											} else {
												v128 = base.B2i32(v119 == int32(18)) << (uint(v117) % 32)
											}
											if v119 == int32(1) {
												v131 = v117
											} else {
												v131 = v128
											}
											v144 = v131
										} else {
											v132 = int32(1)
											if v114&v132 != 0 {
												v144 = int32(base.Ui32(v114)>>(uint(v132)%32)) - v132
											} else {
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
												v144 = int32(base.Ui32(v138)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v144
									}
									v151 = int32(1)
									v152 = v21 + v151
									v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									v157 = v155 & v151
									if v157 != 0 {
										v158 = v152
									} else {
										v158 = v21 + int32(4)
									}
									if v155 == int32(1) {
										v161 = int32(4)
										v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
										if v163&int32(254) == int32(2) {
											v172 = v161
										} else {
											v172 = base.B2i32(v163 == int32(18)) << (uint(v161) % 32)
										}
										if v163 == int32(1) {
											v175 = v161
										} else {
											v175 = v172
										}
										v186 = v175
									} else {
										v176 = int32(1)
										if v157 != 0 {
											v186 = int32(base.Ui32(v155)>>(uint(v176)%32)) - v176
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
											v186 = int32(base.Ui32(v180)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									F_appendBinaryStringInfo(m, v70, v158, v186)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										v190 = v70
										if v190 == int32(0) {
											v198 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v198)
											v201 = int32(0)
										} else {
											v201 = v190
										}
										m.G0 = v12 + int32(16)
										return v201
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
					v151 = int32(1)
					v152 = v21 + v151
					v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					v157 = v155 & v151
					if v157 != 0 {
						v158 = v152
					} else {
						v158 = v21 + int32(4)
					}
					if v155 == int32(1) {
						v161 = int32(4)
						v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
						if v163&int32(254) == int32(2) {
							v172 = v161
						} else {
							v172 = base.B2i32(v163 == int32(18)) << (uint(v161) % 32)
						}
						if v163 == int32(1) {
							v175 = v161
						} else {
							v175 = v172
						}
						v186 = v175
					} else {
						v176 = int32(1)
						if v157 != 0 {
							v186 = int32(base.Ui32(v155)>>(uint(v176)%32)) - v176
						} else {
							v180 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v186 = int32(base.Ui32(v180)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_appendBinaryStringInfo(m, v70, v158, v186)
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
						return int32(0)
					} else {
						v190 = v70
						if v190 == int32(0) {
							v198 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v198)
							v201 = int32(0)
						} else {
							v201 = v190
						}
						m.G0 = v12 + int32(16)
						return v201
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
							v86 = int32(4)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
							if v88&int32(254) == int32(2) {
								v97 = v86
							} else {
								v97 = base.B2i32(v88 == int32(18)) << (uint(v86) % 32)
							}
							if v88 == int32(1) {
								v100 = v86
							} else {
								v100 = v97
							}
							v111 = v100
						} else {
							v101 = int32(1)
							if v82 != 0 {
								v111 = int32(base.Ui32(v80)>>(uint(v101)%32)) - v101
							} else {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
								v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						F_appendBinaryStringInfo(m, v70, v83, v111)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							if v18 != 0 {
							} else {
								v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
								if v114 == int32(1) {
									v117 = int32(4)
									v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
									if v119&int32(254) == int32(2) {
										v128 = v117
									} else {
										v128 = base.B2i32(v119 == int32(18)) << (uint(v117) % 32)
									}
									if v119 == int32(1) {
										v131 = v117
									} else {
										v131 = v128
									}
									v144 = v131
								} else {
									v132 = int32(1)
									if v114&v132 != 0 {
										v144 = int32(base.Ui32(v114)>>(uint(v132)%32)) - v132
									} else {
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v144 = int32(base.Ui32(v138)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v144
							}
							v151 = int32(1)
							v152 = v21 + v151
							v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							v157 = v155 & v151
							if v157 != 0 {
								v158 = v152
							} else {
								v158 = v21 + int32(4)
							}
							if v155 == int32(1) {
								v161 = int32(4)
								v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
								if v163&int32(254) == int32(2) {
									v172 = v161
								} else {
									v172 = base.B2i32(v163 == int32(18)) << (uint(v161) % 32)
								}
								if v163 == int32(1) {
									v175 = v161
								} else {
									v175 = v172
								}
								v186 = v175
							} else {
								v176 = int32(1)
								if v157 != 0 {
									v186 = int32(base.Ui32(v155)>>(uint(v176)%32)) - v176
								} else {
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v186 = int32(base.Ui32(v180)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, v70, v158, v186)
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int32(0)
							} else {
								v190 = v70
								if v190 == int32(0) {
									v198 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v198)
									v201 = int32(0)
								} else {
									v201 = v190
								}
								m.G0 = v12 + int32(16)
								return v201
							}
						}
					}
				}
			}
		}
	}
}
