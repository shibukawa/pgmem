package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsToastNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	if l0 != int32(99) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[277]))
		v13 = base.B2i32(v8 != int32(0)) & base.B2i32(l0 == v8)
	} else {
		v13 = int32(1)
	}
	return v13
}
func F_toast_compress_datum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v40 = int32(4)
		} else {
			v40 = base.B2i32(v15&int32(255) == int32(18)) << (uint(int32(4)) % 32)
		}
	} else {
		v28 = int32(1)
		if v11&v28 != 0 {
			v40 = int32(base.Ui32(v11)>>(uint(v28)%32)) - v28
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[17])))
	if l1 != 0 {
		v43 = l1
	} else {
		v43 = v42
	}
	switch v43&int32(255) - int32(108) {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(453759), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errdetail(m, int32(592485), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(507997), int32(142), int32(292205))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = base.I32_extend8_s(v43)
			F_errmsg_internal(m, int32(514269), v9)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(505474), int32(75), int32(292184))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v84 == int32(1) {
			v87 = int32(4)
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v89&int32(254) == int32(2) {
				v98 = v87
			} else {
				v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
			}
			if v89 == int32(1) {
				v101 = v87
			} else {
				v101 = v98
			}
			v114 = v101
		} else {
			v102 = int32(1)
			if v84&v102 != 0 {
				v114 = int32(base.Ui32(v84)>>(uint(v102)%32)) - v102
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v114 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v116 = *(*int32)(unsafe.Add(mBase, _consts[18]))
		v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
		if v114 < v117 {
			v147 = v3
			v150 = v147
			if v150 != 0 {
				v151 = int32(2)
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
				if base.Ui32(int32(base.Ui32(v153)>>(uint(v151)%32))) < base.Ui32(v40-v151) {
					*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v40
					v161 = v150
					m.G0 = v9 + int32(16)
					return v161
				} else {
					F_pfree(m, v150)
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
					} else {
						v161 = int32(0)
						m.G0 = v9 + int32(16)
						return v161
					}
				}
			} else {
				v161 = int32(0)
				m.G0 = v9 + int32(16)
				return v161
			}
		} else {
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
			if v119 < v114 {
				v147 = v3
				v150 = v147
				if v150 != 0 {
					v151 = int32(2)
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
					if base.Ui32(int32(base.Ui32(v153)>>(uint(v151)%32))) < base.Ui32(v40-v151) {
						*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v40
						v161 = v150
						m.G0 = v9 + int32(16)
						return v161
					} else {
						F_pfree(m, v150)
						mBase = m.M
						v159 = m.ExcPending
						if v159 != 0 {
							return int32(0)
						} else {
							v161 = int32(0)
							m.G0 = v9 + int32(16)
							return v161
						}
					}
				} else {
					v161 = int32(0)
					m.G0 = v9 + int32(16)
					return v161
				}
			} else {
				v123 = F_palloc(m, v114+int32(12))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					v125 = int32(1)
					v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v127&v125 != 0 {
						v130 = v125
					} else {
						v130 = int32(4)
					}
					v134 = int32(0)
					v135 = F_pglz_compress(m, l0+v130, v114, v123+int32(8), v134)
					mBase = m.M
					if v135 < v134 {
						F_pfree(m, v123)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int32(0)
						} else {
							v150 = int32(0)
							if v150 != 0 {
								v151 = int32(2)
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
								if base.Ui32(int32(base.Ui32(v153)>>(uint(v151)%32))) < base.Ui32(v40-v151) {
									*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v40
									v161 = v150
									m.G0 = v9 + int32(16)
									return v161
								} else {
									F_pfree(m, v150)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										v161 = int32(0)
										m.G0 = v9 + int32(16)
										return v161
									}
								}
							} else {
								v161 = int32(0)
								m.G0 = v9 + int32(16)
								return v161
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v123))) = v135<<(uint(int32(2))%32) + int32(34)
						v147 = v123
						v150 = v147
						if v150 != 0 {
							v151 = int32(2)
							v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
							if base.Ui32(int32(base.Ui32(v153)>>(uint(v151)%32))) < base.Ui32(v40-v151) {
								*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v40
								v161 = v150
								m.G0 = v9 + int32(16)
								return v161
							} else {
								F_pfree(m, v150)
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int32(0)
								} else {
									v161 = int32(0)
									m.G0 = v9 + int32(16)
									return v161
								}
							}
						} else {
							v161 = int32(0)
							m.G0 = v9 + int32(16)
							return v161
						}
					}
				}
			}
		}
	}
}
func F_toast_fetch_datum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 != int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(154021), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(503542), int32(351), int32(292276))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v10 != int32(18) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(154021), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(503542), int32(351), int32(292276))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+14))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+10))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+6))
			v18 = v16 & int32(1073741823)
			v20 = v18 + int32(4)
			v21 = F_palloc(m, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(2)
				v26 = v20 << (uint(v25) % 32)
				if base.Ui32(v18) < base.Ui32(v13-int32(4)) {
					v32 = v26 | v25
				} else {
					v32 = v26
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
				if v18 != 0 {
					v35 = F_table_open(m, v14, int32(1))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+188))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+160))
						m.T0[v39].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v35, v15, v18, int32(0), v18, v21)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_sequence_close(m, v35, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								return v21
							}
						}
					}
				} else {
					return v21
				}
			}
		}
	}
}
func F_toast_get_valid_index(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_table_open(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_toast_open_indexes(m, v11, l1, v9+int32(8), v9+int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v19<<(uint(int32(2))%32))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if int32(0) < v27 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pfree(m, v21)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21+v32<<(uint(int32(2))%32))))
	F_relation_close(m, v40, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v45 = v32 + int32(1)
	if v45 != v27 {
		v32 = v45
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_sequence_close(m, v11, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	m.G0 = v9 + int32(16)
	return v26
}
