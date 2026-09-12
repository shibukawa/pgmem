package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_copy_plpgsql_datums(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v17 = F_palloc(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+512))
	v21 = F_palloc(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v14 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+508))
	v29 = int32(0)
	v30 = v21
	v33 = v21
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v12 + int32(16)
	return
L7:
	;
	v38 = v29 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26+v38)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	switch v41 {
	case 0, 4:
		goto L11
	case 1, 3:
		v90 = v40
		v91 = v33
		goto L9
	case 2:
		goto L13
	default:
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v38))) = v90
	v95 = v29 + int32(1)
	if v95 != v14 {
		v29 = v95
		v30 = v91
		v33 = v91
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v90 = v30
	v91 = v89
	goto L9
L11:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v81
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v85
	v89 = v30 + int32(56)
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v44
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v50
	v89 = v30 + int32(40)
	goto L10
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v60
	F_errmsg_internal(m, int32(463320), v12)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(478287), int32(1367), int32(141363))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	goto L8
}
func F_plpgsql_adddatum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1042]))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
	if v11 == v14 {
		*(*int32)(unsafe.Add(mBase, _consts[1044])) = v11 << (uint(int32(1)) % 32)
		v26 = F_repalloc(m, v8, v11<<(uint(int32(3))%32))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1042])) = v26
			v31 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
			v32 = v31
			v33 = v26
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32
			*(*int32)(unsafe.Add(mBase, _consts[1043])) = v32 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v33+v32<<(uint(int32(2))%32)))) = l0
			return
		}
	} else {
		v32 = v11
		v33 = v8
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32
		*(*int32)(unsafe.Add(mBase, _consts[1043])) = v32 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v33+v32<<(uint(int32(2))%32)))) = l0
		return
	}
}
func F_plpgsql_build_datatype_arrayof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = F_get_array_type(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(526891))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v48 = F_format_type_be(m, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v48
							F_errmsg(m, int32(183751), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(474316), int32(2102), int32(322918))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
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
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v23 = F_SearchSysCache1(m, int32(82), v14)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(526891))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v14
							F_errmsg_internal(m, int32(48046), v8+int32(16))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(474316), int32(1960), int32(349826))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = F_build_datatype(m, v23, v21, v20, int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v28
								m.G0 = v8 + int32(32)
								return v32
							}
						}
					}
				}
			}
		}
	} else {
		v32 = l0
		m.G0 = v8 + int32(32)
		return v32
	}
}
func F_plpgsql_compile_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		v11 = F_function_parse_error_transpose(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 != 0 {
				m.G0 = v7 + int32(16)
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[1041]))
				if v15 != 0 {
					F_set_errcontext_domain(m, int32(526891))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _consts[1041]))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+196))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
						F_errcontext_msg(m, int32(456213), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[1041]))
		if v17 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			F_set_errcontext_domain(m, int32(526891))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[1041]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+196))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
				F_errcontext_msg(m, int32(456213), v7)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_plpgsql_exec_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v10 != 0 {
		v19 = v10 + int32(12)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v21 != 0 {
			if v20 <= int32(0) {
				F_set_errcontext_domain(m, int32(526891))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
					v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v130
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v129
					F_errcontext_msg(m, int32(171046), v6+int32(-48))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						m.G0 = v8 - int32(-64)
						return
					}
				}
			} else {
				F_set_errcontext_domain(m, int32(526891))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v30
					F_errcontext_msg(m, int32(187600), v6+int32(-16))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v8 - int32(-64)
						return
					}
				}
			}
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
			if v41 == int32(0) {
				F_set_errcontext_domain(m, int32(526891))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v116
					F_errcontext_msg(m, int32(174098), v8)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						m.G0 = v8 - int32(-64)
						return
					}
				}
			} else {
				if v20 <= int32(0) {
					F_set_errcontext_domain(m, int32(526891))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v116
						F_errcontext_msg(m, int32(174098), v8)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							m.G0 = v8 - int32(-64)
							return
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(526891))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						switch v56 {
						case 0:
							v96 = int32(302300)
							v98 = v96
						case 1:
							v98 = int32(89349)
						case 2:
							v98 = int32(514601)
						case 3:
							v98 = int32(516919)
						case 4:
							v98 = int32(504491)
						case 5:
							v98 = int32(517874)
						case 6:
							v98 = int32(378995)
						case 7:
							v98 = int32(106536)
						case 8:
							v98 = int32(200249)
						case 9:
							v98 = int32(23292)
						case 10:
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
							if v68 != 0 {
								v69 = int32(498389)
							} else {
								v69 = int32(515126)
							}
							v98 = v69
						case 11:
							v98 = int32(506116)
						case 12:
							v98 = int32(494757)
						case 13:
							v98 = int32(486519)
						case 14:
							v98 = int32(516902)
						case 15:
							v98 = int32(495155)
						case 16:
							v98 = int32(90234)
						case 17:
							v98 = int32(515425)
						case 18:
							v98 = int32(90248)
						case 19:
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
							if v80 != 0 {
								v81 = int32(501591)
							} else {
								v81 = int32(501537)
							}
							v98 = v81
						case 20:
							v98 = int32(508408)
						case 21:
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+32)))
							if v85 != 0 {
								v86 = int32(515012)
							} else {
								v86 = int32(513584)
							}
							v98 = v86
						case 22:
							v98 = int32(516789)
						case 23:
							v98 = int32(508937)
						case 24:
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)))
							if v91 != 0 {
								v92 = int32(511149)
							} else {
								v92 = int32(505276)
							}
							v98 = v92
						case 25:
							v98 = int32(498732)
						case 26:
							v98 = int32(511739)
						default:
							v96 = int32(232230)
							v98 = v96
						}
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v98
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v52
						F_errcontext_msg(m, int32(170159), v6+int32(-32))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							m.G0 = v8 - int32(-64)
							return
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		if v13 == int32(0) {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v108 != 0 {
				F_set_errcontext_domain(m, int32(526891))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
					v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v130
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v129
					F_errcontext_msg(m, int32(171046), v6+int32(-48))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						m.G0 = v8 - int32(-64)
						return
					}
				}
			} else {
				F_set_errcontext_domain(m, int32(526891))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v116
					F_errcontext_msg(m, int32(174098), v8)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						m.G0 = v8 - int32(-64)
						return
					}
				}
			}
		} else {
			v19 = v13 + int32(4)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v21 != 0 {
				if v20 <= int32(0) {
					F_set_errcontext_domain(m, int32(526891))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return
					} else {
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
						v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v130
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v129
						F_errcontext_msg(m, int32(171046), v6+int32(-48))
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return
						} else {
							m.G0 = v8 - int32(-64)
							return
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(526891))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v30
						F_errcontext_msg(m, int32(187600), v6+int32(-16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							m.G0 = v8 - int32(-64)
							return
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				if v41 == int32(0) {
					F_set_errcontext_domain(m, int32(526891))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v116
						F_errcontext_msg(m, int32(174098), v8)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							m.G0 = v8 - int32(-64)
							return
						}
					}
				} else {
					if v20 <= int32(0) {
						F_set_errcontext_domain(m, int32(526891))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v116
							F_errcontext_msg(m, int32(174098), v8)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								m.G0 = v8 - int32(-64)
								return
							}
						}
					} else {
						F_set_errcontext_domain(m, int32(526891))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							switch v56 {
							case 0:
								v96 = int32(302300)
								v98 = v96
							case 1:
								v98 = int32(89349)
							case 2:
								v98 = int32(514601)
							case 3:
								v98 = int32(516919)
							case 4:
								v98 = int32(504491)
							case 5:
								v98 = int32(517874)
							case 6:
								v98 = int32(378995)
							case 7:
								v98 = int32(106536)
							case 8:
								v98 = int32(200249)
							case 9:
								v98 = int32(23292)
							case 10:
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
								if v68 != 0 {
									v69 = int32(498389)
								} else {
									v69 = int32(515126)
								}
								v98 = v69
							case 11:
								v98 = int32(506116)
							case 12:
								v98 = int32(494757)
							case 13:
								v98 = int32(486519)
							case 14:
								v98 = int32(516902)
							case 15:
								v98 = int32(495155)
							case 16:
								v98 = int32(90234)
							case 17:
								v98 = int32(515425)
							case 18:
								v98 = int32(90248)
							case 19:
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
								if v80 != 0 {
									v81 = int32(501591)
								} else {
									v81 = int32(501537)
								}
								v98 = v81
							case 20:
								v98 = int32(508408)
							case 21:
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+32)))
								if v85 != 0 {
									v86 = int32(515012)
								} else {
									v86 = int32(513584)
								}
								v98 = v86
							case 22:
								v98 = int32(516789)
							case 23:
								v98 = int32(508937)
							case 24:
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)))
								if v91 != 0 {
									v92 = int32(511149)
								} else {
									v92 = int32(505276)
								}
								v98 = v92
							case 25:
								v98 = int32(498732)
							case 26:
								v98 = int32(511739)
							default:
								v96 = int32(232230)
								v98 = v96
							}
							*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v98
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v20
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v52
							F_errcontext_msg(m, int32(170159), v6+int32(-32))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								m.G0 = v8 - int32(-64)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_plpgsql_exec_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	v6 = l5
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_plpgsql_estate_setup(m, v12+int32(32), l0, v16, l2, l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+63)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(5438)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(336889)
	v29 = int32(4436760)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v12 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v30
	v36 = v12 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v36
	F_copy_plpgsql_datums(m, v36, l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(157548)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if int32(0) < v45 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v49 = l1 + int32(20)
	v55 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(11435)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190+v191<<(uint(int32(2))%32))))
	v196 = int32(0)
	F_assign_simple_var(m, v12+int32(32), v195, v196, v196, v196)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	v62 = int32(2)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)+v55<<(uint(v62)%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61+v65<<(uint(v62)%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v70 {
	case 0:
		goto L13
	default:
		goto L11
	case 2:
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	v174 = v55 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v174 < v175 {
		v55 = v174
		goto L7
	} else {
		goto L39
	}
L10:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v86)+2))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	F_MemoryContextSetParent(m, v162, v160)
	mBase = m.M
	goto L37
L11:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	v108 = v49 + v55<<(uint(int32(3))%32)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+4)))
	if v109 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v75 = v49 + v55<<(uint(int32(3))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)))
	F_assign_simple_var(m, v12+int32(32), v69, v76, v77, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+44)))
	if v81 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)))
	if v83 != int32(65535) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+40))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v87 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+20)))
	if v93 != int32(1) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	switch v90 - int32(2) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L17
	}
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	v100 = F_expand_array(m, v86, v98, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_assign_simple_var(m, v12+int32(32), v69, v100, int32(0), int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	if v123 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	F_exec_move_row_from_datum(m, v12+int32(32), v69, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v119 = int32(0)
	F_exec_move_row(m, v12+int32(32), v69, v119, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L22
L27:
	;
	goto L22
L28:
	;
	F_SPI_freetuptable(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v128 == v126 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	F_MemoryContextReset(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L9
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+508))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v55<<(uint(int32(2))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v145
	F_errmsg_internal(m, int32(463320), v12)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(478287), int32(614), int32(239924))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_assign_simple_var(m, v12+int32(32), v69, v161+int32(12), int32(0), int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L9
L39:
	;
	goto L8
L40:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v203 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v242 != 0 {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v208
	v239 = v208
	goto L41
L43:
	;
	goto L44
L44:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if v210 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	if v232 == int32(0) {
		v239 = v230
		goto L41
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(0)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v215
	v230 = v215
	v231 = v203
	goto L45
L47:
	;
	goto L48
L48:
	;
	m.T0[v210].(func(*base.Module, int32, int32))(m, v12+int32(32), l0)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v226
	if v223 == v224 {
		v239 = v226
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v230 = v226
	v231 = v223
	goto L45
L51:
	;
	m.T0[v232].(func(*base.Module, int32, int32))(m, v12+int32(32), v230)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v239 = v230
	goto L41
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v248 = F_exec_stmt_block(m, v12+int32(32), v239)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = int32(0)
	if v248 == int32(2) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	if v254 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	m.T0[v254].(func(*base.Module, int32, int32))(m, v12+int32(32), v239)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L125
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L121
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(352098)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v268)
	v271 = l1 + int32(16)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+61)))
	if v272 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(0)
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L117
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(93845)
	v405 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v406 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L68:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	if v275 == int32(0) {
		goto L63
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v268&int32(1) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v278 != int32(383) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+12)))
	if v281&int32(2) == int32(0) {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+16)) = int32(2)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v288 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+24)) = v288
	v290 = int32(4443856)
	v291 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v295 = F_CreateTupleDescCopy(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
	v303 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v303)
	goto L67
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+28)) = v295
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v291
	goto L76
L78:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
	if v310 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
	if v385 != int32(1) {
		goto L67
	} else {
		goto L104
	}
L81:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v313 == int32(2249) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v373 = int32(-1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v376 = F_exec_cast_value(m, v12+int32(32), v372, v271, v309, v373, v374, v373)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L100
	}
L84:
	;
	v326 = F_get_call_result_type(m, l1, v12+int32(16), v12+int32(12))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L92
	}
L85:
	;
	if v313 != v309 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v319 = F_SPI_datumTransfer(m, v317, int32(-1))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v319
	goto L67
L88:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L97
	}
L89:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v349 = F_SPI_datumTransfer(m, v347, int32(-1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L96
	}
L90:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_coerce_function_result_tuple(m, v12+int32(32), v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_coerce_function_result_tuple(m, v12+int32(32), v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	switch v326 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L88
	}
L93:
	;
	goto L67
L94:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v341 = int32(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_domain_check(m, v340, v341, v342, v341, v341)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L67
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v349
	goto L67
L97:
	;
	F_errmsg_internal(m, int32(351051), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(478287), int32(746), int32(239924))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v376
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v379 != 0 {
		goto L67
	} else {
		goto L101
	}
L101:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v380 != 0 {
		goto L67
	} else {
		goto L102
	}
L102:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v382 = F_SPI_datumTransfer(m, v376, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v382
	goto L67
L104:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v392 = int32(-1)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v395 = F_exec_cast_value(m, v12+int32(32), v390, v271, v391, v392, v393, v392)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v395
	goto L67
L106:
	;
	v417 = int32(4537072)
	v418 = *(*int32)(unsafe.Add(mBase, _consts[1046]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	F_pfree(m, v418)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L110
	}
L107:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	if v409 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	m.T0[v409].(func(*base.Module, int32, int32))(m, v12+int32(32), l0)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1046])) = v419
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	F_FreeExprContext(m, v423, int32(1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v427 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v427
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	if v429 == v427 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v444
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	m.G0 = v12 + int32(176)
	return v446
L113:
	;
	F_SPI_freetuptable(m, v429)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v434 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v436 == v434 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v436)+20))
	F_MemoryContextReset(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	F_errcode(m, int32(83887490))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(506020), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(478287), int32(641), int32(239924))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(99802), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(478287), int32(659), int32(239924))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(57994), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(478287), int32(664), int32(239924))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_extra_errors_assign_hook(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[1055])) = v4
	return
}
func F_plpgsql_inline_handler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v82 int64
	_ = v82
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v276 int32
	_ = v276
	var v294 int32
	_ = v294
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v350 int32
	_ = v350
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v378 int32
	_ = v378
	var v390 int32
	_ = v390
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int64
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v23 = v2
	v24 = v2
	v25 = v2
	v26 = v2
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = int32(-1)
	v35 = v18
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v18 + int32(48)
	return v165
L4:
	;
	goto L3
L5:
	;
	if v34 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v404 = int32(m.ExcTag)
	v405 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v404 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L8:
	;
	v38 = int32(32)
	v39 = v35 - v38
	m.G0 = v39
	v42 = v39 - v38
	m.G0 = v42
	v45 = v42 - int32(160)
	m.G0 = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v25
	v56 = v47 + int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v39
	F_SPI_connect_ext(m, v48^int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		v403 = v45
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v136 = v23
	v137 = v24
	v138 = v25
	v139 = v26
	v140 = v27
	v141 = v28
	v142 = v29
	v143 = v30
	v144 = v31
	v145 = v32
	v148 = v35
	v149 = v33
	goto L10
L10:
	;
	if v149 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v39
	v76 = F_plpgsql_compile_inline(m, v65)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		v403 = v45
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v78 + int64(1)
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+12)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v42
	v96 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v28
	v103 = v76 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v39
	v110 = F_CreateExecutorState(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		v403 = v45
		goto L7
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v39
	v124 = F_ResourceOwnerCreate(m, int32(0), int32(137492))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		v403 = v45
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	v129 = *(*int32)(unsafe.Add(mBase, _consts[294]))
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v18 + int32(4)
	goto L18
L16:
	;
	v136 = v124
	v137 = v42
	v138 = v76
	v139 = v39
	v140 = v45
	v141 = v110
	v142 = v103
	v143 = v129
	v144 = v127
	v145 = v56
	v148 = v45
	v149 = int32(0)
	goto L10
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v140
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	v165 = F_plpgsql_exec_function(m, v138, v139, v141, v136, v136, v154)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v144
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	v310 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+8))
	goto L33
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v144
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_FreeExecutorState(m, v141)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_ReleaseAllPlanCacheRefsInOwner(m, v136)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_ResourceOwnerDelete(m, v136)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v207 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_plpgsql_free_function_memory(m, v138)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	v233 = F_SPI_finish(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L27
	}
L27:
	;
	if v233 == int32(2) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	v261 = F_SPI_result_code_string(m, v233)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v261
	F_errmsg_internal(m, int32(194307), v18)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_errfinish(m, int32(473865), int32(427), int32(209455))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	v323 = int32(0)
	F_plpgsql_subxact_cb(m, int32(2), v311, v323, v323)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_FreeExecutorState(m, v141)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_ReleaseAllPlanCacheRefsInOwner(m, v136)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_ResourceOwnerDelete(m, v136)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v363 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_plpgsql_free_function_memory(m, v138)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v139
	F_pg_re_throw(m)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		v403 = v148
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	v409 = int32(v405)
	m.G0 = v403
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	if v18+int32(4) == v416 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	m.ExcPending = 1
	goto L49
L42:
	;
	if v419 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v419 = v418
	goto L45
L44:
	;
	v419 = int32(0)
	goto L45
L45:
	;
	goto L42
L46:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v23 = v427
	v24 = v421
	v25 = v424
	v26 = v420
	v27 = v422
	v28 = v426
	v29 = v425
	v30 = v428
	v31 = v429
	v32 = v423
	v33 = v411
	v34 = v419
	v35 = v403
	goto L2
L47:
	;
	goto L48
L48:
	;
	F___wasm_longjmp(m, v412, v411)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_push_back_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(8)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if int32(4) <= v20 {
		F_errstart_cold(m, int32(21), int32(526891))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(305202), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_errfinish(m, int32(473835), int32(388), int32(269200))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32))+76)) = l0
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
		v50 = v19 + v47*int32(24)
		v51 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		*(*int64)(unsafe.Add(mBase, uint32(v50)+100)) = v51
		v53 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v50)+92)) = v53
		*(*int32)(unsafe.Add(mBase, uint32(v50)+112)) = v42
		*(*int32)(unsafe.Add(mBase, uint32(v50)+108)) = v41
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
		*(*int32)(unsafe.Add(mBase, uint32(v57)+72)) = v58 + int32(1)
		m.G0 = v11 + int32(16)
		return
	}
}
func F_plpgsql_subxact_cb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v7 = int32(1)
	if base.Ui32(v7) < base.Ui32(l0-v7) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1046]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v13
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 != l1 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	F_FreeExprContext(m, v25, base.B2i32(l0 == int32(1)))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v28 = int32(4537072)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1046]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	F_pfree(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1046])) = v30
	if v30 != 0 {
		v20 = v30
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
}
func F_plpgsql_validator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v148 int32
	_ = v148
	var v153 int64
	_ = v153
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_CheckFunctionValidatorAccess(m, v14, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L76
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L71
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L67
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(526891))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L64
	}
L5:
	;
	return int32(0)
L6:
	;
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = F_SearchSysCache1(m, int32(47), v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	m.G0 = v11 + int32(160)
	return int32(0)
L10:
	;
	if v21 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+108))
	v30 = F_get_typtype(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v69 = F_get_func_arg_info(m, v21, v11+int32(156), v11+int32(152), v11+int32(148))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L27
	}
L13:
	;
	if v30 != int32(112) {
		v61 = v2
		v62 = v2
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+108))
	if v34 <= int32(3830) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v61 = int32(0)
	v62 = v58
	goto L12
L16:
	;
	v58 = int32(0)
	goto L15
L17:
	;
	switch v34 - int32(2249) {
	case 0, 28, 29, 34:
		goto L16
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 31, 32, 33:
		goto L1
	case 30:
		v61 = int32(1)
		v62 = v2
		goto L12
	default:
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(v34-int32(5077)) < base.Ui32(int32(4)) {
		goto L16
	} else {
		goto L23
	}
L20:
	;
	if v34 == int32(2776) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if v34 == int32(3500) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L1
L23:
	;
	switch v34 - int32(3831) {
	case 0:
		goto L16
	case 1, 2, 3, 4, 5, 6:
		goto L1
	case 7:
		goto L25
	default:
		goto L24
	}
L24:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v34-int32(4537)) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v58 = int32(1)
	goto L15
L26:
	;
	goto L16
L27:
	;
	if int32(0) < v69 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1056])))
	if v128 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L31:
	;
	v82 = v73 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
	v86 = F_get_typtype(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v117 = v73 + int32(1)
	if v117 != v69 {
		v73 = v117
		goto L31
	} else {
		goto L49
	}
L34:
	;
	if v86 != int32(112) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90+v82)))
	if v92 <= int32(3830) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v92 != int32(2249) {
		goto L2
	} else {
		goto L48
	}
L37:
	;
	if v92 <= int32(2775) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v92-int32(5077)) < base.Ui32(int32(4)) {
		goto L33
	} else {
		goto L45
	}
L40:
	;
	switch v92 - int32(2277) {
	case 0, 6:
		goto L33
	case 1, 2, 3, 4, 5:
		goto L2
	default:
		goto L36
	}
L41:
	;
	goto L42
L42:
	;
	if v92 == int32(2776) {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	if v92 != int32(3500) {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	if base.Ui32(v92-int32(4537)) < base.Ui32(int32(2)) {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	if v92 == int32(3831) {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	goto L2
L48:
	;
	goto L33
L49:
	;
	goto L32
L50:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L63
	}
L53:
	;
	v134 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v11)+140)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v11)+132)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v15
	v148 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v11 + int32(96)
	if v61 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v186 = F_plpgsql_compile(m, v11+int32(128), int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L60
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v181
	goto L54
L56:
	;
	v153 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(442)
	v181 = v11 + int32(52)
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v62 == int32(0) {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(441)
	v181 = v11 + int32(52)
	goto L55
L60:
	;
	v188 = F_SPI_finish(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	if v188 != int32(2) {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	goto L52
L63:
	;
	goto L9
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	F_errmsg_internal(m, int32(42408), v11)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(473865), int32(462), int32(199702))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v225 = F_SPI_result_code_string(m, v188)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v225
	F_errmsg_internal(m, int32(194307), v11+int32(48))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(473865), int32(544), int32(199702))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245+v73<<(uint(int32(2))%32))))
	v250 = F_format_type_be(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v250
	F_errmsg(m, int32(178998), v11+int32(32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(473865), int32(497), int32(199702))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v28)+108))
	v272 = F_format_type_be(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v272
	F_errmsg(m, int32(181793), v11+int32(16))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(473865), int32(481), int32(199702))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_xact_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	if base.Ui32(int32(4)) < base.Ui32(l0) {
		if l0&int32(-2) != int32(2) {
		} else {
			v37 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[1053])) = v37
			*(*int32)(unsafe.Add(mBase, _consts[1046])) = v37
			*(*int32)(unsafe.Add(mBase, _consts[1054])) = int32(0)
		}
		return
	} else {
		if int32(1)<<(uint(l0)%32)&int32(19) == int32(0) {
			if l0&int32(-2) != int32(2) {
			} else {
				v37 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[1053])) = v37
				*(*int32)(unsafe.Add(mBase, _consts[1046])) = v37
				*(*int32)(unsafe.Add(mBase, _consts[1054])) = int32(0)
			}
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1046])) = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, _consts[1053]))
			if v18 != 0 {
				F_FreeExecutorState(m, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[1053])) = v21
					v28 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
					if v28 == v21 {
						*(*int32)(unsafe.Add(mBase, _consts[1054])) = int32(0)
						return
					} else {
						F_ReleaseAllPlanCacheRefsInOwner(m, v28)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1054])) = int32(0)
							return
						}
					}
				}
			} else {
				v21 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[1053])) = v21
				v28 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
				if v28 == v21 {
					*(*int32)(unsafe.Add(mBase, _consts[1054])) = int32(0)
					return
				} else {
					F_ReleaseAllPlanCacheRefsInOwner(m, v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1054])) = int32(0)
						return
					}
				}
			}
		}
	}
}
