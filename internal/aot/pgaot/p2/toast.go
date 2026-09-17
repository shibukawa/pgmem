package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsToastNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	if l0 != int32(99) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_IsToastNamespace[0]))
		v12 = base.B2i32(v6 != int32(0)) & base.B2i32(l0 == v6)
	} else {
		v12 = int32(1)
	}
	return v12
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
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
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v39 = int32(4)
		} else {
			if v15 == int32(18) {
				v26 = int32(16)
			} else {
				v26 = int32(0)
			}
			v39 = v26
		}
	} else {
		v27 = int32(1)
		if v11&v27 != 0 {
			v39 = int32(base.Ui32(v11)>>(uint(v27)%32)) - v27
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_toast_compress_datum[0])))
	if l1 != 0 {
		v42 = l1
	} else {
		v42 = v41
	}
	switch v42&int32(255) - int32(108) {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_toast_compress_datum_0), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					F_errdetail(m, int32(_a_F_toast_compress_datum_1), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_toast_compress_datum_2), int32(142), int32(_a_F_toast_compress_datum_3))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
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
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = base.I32_extend8_s(v42)
			F_errmsg_internal(m, int32(_a_F_toast_compress_datum_4), v9)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_toast_compress_datum_5), int32(75), int32(_a_F_toast_compress_datum_6))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v83 == int32(1) {
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
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
			v112 = v99
		} else {
			v100 = int32(1)
			if v83&v100 != 0 {
				v112 = int32(base.Ui32(v83)>>(uint(v100)%32)) - v100
			} else {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v114 = *(*int32)(unsafe.Add(mBase, _c_F_toast_compress_datum[1]))
		v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
		if v112 < v115 {
			v145 = v3
			v148 = v145
			if v148 != 0 {
				v149 = int32(2)
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
				if base.Ui32(int32(base.Ui32(v151)>>(uint(v149)%32))) < base.Ui32(v39-v149) {
					*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v39
					v159 = v148
					m.G0 = v9 + int32(16)
					return v159
				} else {
					F_pfree(m, v148)
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return int32(0)
					} else {
						v159 = int32(0)
						m.G0 = v9 + int32(16)
						return v159
					}
				}
			} else {
				v159 = int32(0)
				m.G0 = v9 + int32(16)
				return v159
			}
		} else {
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
			if v117 < v112 {
				v145 = v3
				v148 = v145
				if v148 != 0 {
					v149 = int32(2)
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
					if base.Ui32(int32(base.Ui32(v151)>>(uint(v149)%32))) < base.Ui32(v39-v149) {
						*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v39
						v159 = v148
						m.G0 = v9 + int32(16)
						return v159
					} else {
						F_pfree(m, v148)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							v159 = int32(0)
							m.G0 = v9 + int32(16)
							return v159
						}
					}
				} else {
					v159 = int32(0)
					m.G0 = v9 + int32(16)
					return v159
				}
			} else {
				v121 = F_palloc(m, v112+int32(12))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					v123 = int32(1)
					v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v125&v123 != 0 {
						v128 = v123
					} else {
						v128 = int32(4)
					}
					v132 = int32(0)
					v133 = F_pglz_compress(m, l0+v128, v112, v121+int32(8), v132)
					mBase = m.M
					if v133 < v132 {
						F_pfree(m, v121)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							v148 = int32(0)
							if v148 != 0 {
								v149 = int32(2)
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
								if base.Ui32(int32(base.Ui32(v151)>>(uint(v149)%32))) < base.Ui32(v39-v149) {
									*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v39
									v159 = v148
									m.G0 = v9 + int32(16)
									return v159
								} else {
									F_pfree(m, v148)
									mBase = m.M
									v157 = m.ExcPending
									if v157 != 0 {
										return int32(0)
									} else {
										v159 = int32(0)
										m.G0 = v9 + int32(16)
										return v159
									}
								}
							} else {
								v159 = int32(0)
								m.G0 = v9 + int32(16)
								return v159
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v121))) = v133<<(uint(int32(2))%32) + int32(34)
						v145 = v121
						v148 = v145
						if v148 != 0 {
							v149 = int32(2)
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
							if base.Ui32(int32(base.Ui32(v151)>>(uint(v149)%32))) < base.Ui32(v39-v149) {
								*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v39
								v159 = v148
								m.G0 = v9 + int32(16)
								return v159
							} else {
								F_pfree(m, v148)
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									v159 = int32(0)
									m.G0 = v9 + int32(16)
									return v159
								}
							}
						} else {
							v159 = int32(0)
							m.G0 = v9 + int32(16)
							return v159
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
			F_errmsg_internal(m, int32(_a_F_toast_fetch_datum_0), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_toast_fetch_datum_1), int32(351), int32(_a_F_toast_fetch_datum_2))
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
				F_errmsg_internal(m, int32(_a_F_toast_fetch_datum_0), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_toast_fetch_datum_1), int32(351), int32(_a_F_toast_fetch_datum_2))
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
							F_relation_close(m, v35, int32(1))
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
	F_relation_close(m, v11, int32(0))
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
