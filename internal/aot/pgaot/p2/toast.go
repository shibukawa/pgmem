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
func F_toast_compress_datum(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(1) {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
		if v19 == int32(18) {
			v22 = int32(16)
		} else {
			v22 = int32(0)
		}
		if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v29 = int32(4)
		} else {
			v29 = v22
		}
		v42 = v29
	} else {
		v30 = int32(1)
		if v13&v30 != 0 {
			v42 = int32(base.Ui32(v13)>>(uint(v30)%32)) - v30
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_toast_compress_datum[0])))
	if l1 != 0 {
		v45 = l1
	} else {
		v45 = v44
	}
	switch v45&int32(255) - int32(108) {
	case 0:
		v50 = m.G0
		v52 = v50 - int32(32)
		m.G0 = v52
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = int32(_a_F_toast_compress_datum_0)
				F_errmsg(m, int32(_a_F_toast_compress_datum_1), v52+int32(16))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(_a_F_toast_compress_datum_0)
					v73 = F_errdetail(m, int32(_a_F_toast_compress_datum_2), v52)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_toast_compress_datum_3), int32(142), int32(_a_F_toast_compress_datum_4))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = base.I32_extend8_s(v45)
			F_errmsg_internal(m, int32(_a_F_toast_compress_datum_5), v10)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_toast_compress_datum_6), int32(75), int32(_a_F_toast_compress_datum_7))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v94 == int32(1) {
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v100 == int32(18) {
				v103 = int32(16)
			} else {
				v103 = int32(0)
			}
			if base.Ui32((v100-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v110 = int32(4)
			} else {
				v110 = v103
			}
			v123 = v110
		} else {
			v111 = int32(1)
			if v94&v111 != 0 {
				v123 = int32(base.Ui32(v94)>>(uint(v111)%32)) - v111
			} else {
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v123 = int32(base.Ui32(v117)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v125 = *(*int32)(unsafe.Add(mBase, _c_F_toast_compress_datum[1]))
		v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
		if v123 < v126 {
			v156 = v3
			v159 = v156
			v160 = int64(0)
			if v159 == int32(0) {
				v173 = v160
				m.G0 = v10 + int32(16)
				return v173
			} else {
				v163 = int32(2)
				v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
				if base.Ui32(int32(base.Ui32(v165)>>(uint(v163)%32))) < base.Ui32(v42-v163) {
					*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v42
					v173 = base.I64_extend_i32_u(v159)
					m.G0 = v10 + int32(16)
					return v173
				} else {
					F_pfree(m, v159)
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return int64(0)
					} else {
						v173 = v160
						m.G0 = v10 + int32(16)
						return v173
					}
				}
			}
		} else {
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
			if v128 < v123 {
				v156 = v3
				v159 = v156
				v160 = int64(0)
				if v159 == int32(0) {
					v173 = v160
					m.G0 = v10 + int32(16)
					return v173
				} else {
					v163 = int32(2)
					v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
					if base.Ui32(int32(base.Ui32(v165)>>(uint(v163)%32))) < base.Ui32(v42-v163) {
						*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v42
						v173 = base.I64_extend_i32_u(v159)
						m.G0 = v10 + int32(16)
						return v173
					} else {
						F_pfree(m, v159)
						mBase = m.M
						v172 = m.ExcPending
						if v172 != 0 {
							return int64(0)
						} else {
							v173 = v160
							m.G0 = v10 + int32(16)
							return v173
						}
					}
				}
			} else {
				v132 = F_palloc(m, v123+int32(12))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int64(0)
				} else {
					v134 = int32(1)
					v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					if v136&v134 != 0 {
						v139 = v134
					} else {
						v139 = int32(4)
					}
					v143 = int32(0)
					v144 = F_pglz_compress(m, v12+v139, v123, v132+int32(8), v143)
					mBase = m.M
					if v144 < v143 {
						F_pfree(m, v132)
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int64(0)
						} else {
							v159 = int32(0)
							v160 = int64(0)
							if v159 == int32(0) {
								v173 = v160
								m.G0 = v10 + int32(16)
								return v173
							} else {
								v163 = int32(2)
								v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
								if base.Ui32(int32(base.Ui32(v165)>>(uint(v163)%32))) < base.Ui32(v42-v163) {
									*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v42
									v173 = base.I64_extend_i32_u(v159)
									m.G0 = v10 + int32(16)
									return v173
								} else {
									F_pfree(m, v159)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int64(0)
									} else {
										v173 = v160
										m.G0 = v10 + int32(16)
										return v173
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v132))) = v144<<(uint(int32(2))%32) + int32(34)
						v156 = v132
						v159 = v156
						v160 = int64(0)
						if v159 == int32(0) {
							v173 = v160
							m.G0 = v10 + int32(16)
							return v173
						} else {
							v163 = int32(2)
							v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
							if base.Ui32(int32(base.Ui32(v165)>>(uint(v163)%32))) < base.Ui32(v42-v163) {
								*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v42
								v173 = base.I64_extend_i32_u(v159)
								m.G0 = v10 + int32(16)
								return v173
							} else {
								F_pfree(m, v159)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return int64(0)
								} else {
									v173 = v160
									m.G0 = v10 + int32(16)
									return v173
								}
							}
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
func F_toast_get_valid_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_table_open(m, l0, int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = int32(8)
	v21 = F_toast_open_indexes(m, v12, v16, v9+v16, v9+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v21<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if int32(0) < v29 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pfree(m, v23)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23+v33<<(uint(int32(2))%32))))
	F_relation_close(m, v42, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v47 = v33 + int32(1)
	if v47 != v29 {
		v33 = v47
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_relation_close(m, v12, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	m.G0 = v9 + int32(16)
	return v28
}
