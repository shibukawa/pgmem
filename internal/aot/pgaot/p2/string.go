package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_appendStringInfoSpaces(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	if int32(0) < l1 {
		F_enlargeStringInfo(m, l0, l1)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v12 = F__emscripten_memset_bulkmem(m, v7+v8, base.I32_extend8_s(int32(32)), l1)
			mBase = m.M
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = v13 + l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v16+v14))) = uint8(v18)
			return
		}
	} else {
		return
	}
}
func F_appendStringInfoString(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_strlen(m, l1)
	mBase = m.M
	F_enlargeStringInfo(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v4 != 0 {
			v10 = F__emscripten_memcpy_bulkmem(m, v7+v8, l1, v4)
			mBase = m.M
		} else {
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = v4 + v12
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15+v13))) = uint8(v17)
		return
	}
}
func F_appendStringInfoVA(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v8 - v9
	if int32(16) <= v10 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_pvsnprintf(m, v13+v9, v10, l1, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if base.Ui32(v15) < base.Ui32(v10) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + v15
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v25+v26))) = uint8(v28)
				v30 = v15
				return v30
			}
		}
	} else {
		v30 = int32(32)
		return v30
	}
}
func F_makeStringConst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(2010044694600)
		return v5
	}
}
func F_stringToNode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = int32(0)
	v3 = int32(4443900)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[483]))
	*(*int32)(unsafe.Add(mBase, _consts[483])) = l0
	v9 = F_nodeRead(m, v2, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[483])) = v4
		return v9
	}
}
func F_string_agg_transfn(m *base.Module, l0 int32) int32 {
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
						F_errmsg_internal(m, int32(62819), int32(0))
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512786), int32(5432), int32(362061))
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
					v61 = int32(4536272)
					v62 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v64
					v66 = F_makeStringInfo(m)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v62
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
func F_string_field_used(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if l1 == v6 {
		v28 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v28
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if l1 == v8 {
		v28 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if l1 == v10 {
		v28 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = l0 + int32(56)
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v18 = int32(0)
	v19 = base.B2i32(v17 != v18)
	if v17 == v18 {
		v28 = v19
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v28 = v19
	goto L1
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if l1 == v22 {
		v28 = v19
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	if l1 != v24 {
		v14 = v17
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
}
func F_string_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v9 = l2 + int32(256)
	v13 = base.B2i32(v9 != v4)
	if v6&int32(3) == v4 {
		v39 = v6
		v41 = v9
		v42 = v13
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v112 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v112 = int32(0)
	goto L1
L3:
	;
	v90 = v83
	v92 = v85
	goto L21
L4:
	;
	if v42 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L5:
	;
	if v9 == int32(0) {
		v39 = v6
		v41 = v9
		v42 = v13
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = v6
	v24 = v9
	goto L7
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v27 == int32(0) {
		v83 = v22
		v85 = v24
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v39 = v34
	v41 = v30
	v42 = v32
	goto L4
L9:
	;
	v29 = int32(1)
	v30 = v24 - v29
	v31 = int32(0)
	v32 = base.B2i32(v30 != v31)
	v34 = v22 + v29
	if v34&int32(3) == v31 {
		v39 = v34
		v41 = v30
		v42 = v32
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v30 != 0 {
		v22 = v34
		v24 = v30
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v46 == int32(0) {
		v76 = v39
		v78 = v41
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v78 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L14:
	;
	if base.Ui32(v41) < base.Ui32(int32(4)) {
		v76 = v39
		v78 = v41
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v56 = v39
	v58 = v41
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v63 = v62 ^ int32(0)
	v66 = int32(-2139062144)
	if (int32(16843008)-v63|v63)&v66 != v66 {
		v83 = v56
		v85 = v58
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v76 = v71
	v78 = v73
	goto L13
L18:
	;
	v70 = int32(4)
	v71 = v56 + v70
	v73 = v58 - v70
	if base.Ui32(int32(3)) < base.Ui32(v73) {
		v56 = v71
		v58 = v73
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v83 = v76
	v85 = v78
	goto L3
L21:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if int32(0) == v95 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L2
L23:
	;
	v112 = v90
	goto L1
L24:
	;
	goto L25
L25:
	;
	v97 = int32(1)
	v100 = v92 - v97
	if v100 != 0 {
		v90 = v90 + v97
		v92 = v100
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v114 = v112 - v6
	goto L29
L28:
	;
	v114 = v9
	goto L29
L29:
	;
	if base.Ui32(v114) < base.Ui32(l2) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v116 = v114
	goto L32
L31:
	;
	v116 = l2
	goto L32
L32:
	;
	if v116 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v119 = v114 + v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116 + v6
	return v116
L34:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, l1, v6, v116)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
}
