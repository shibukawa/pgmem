package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_getkeepalivesinterval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v22 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v22 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
			if v13 != 0 {
				v22 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
				if v14 != 0 {
					v22 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(388))))
					v22 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v22
}
func F_pq_getmsgbytes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	if int32(0) <= l1 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if l1 <= v6-v7 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v7
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			return v31 + v7
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(398571), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485954), int32(515), int32(156379))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(398571), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485954), int32(515), int32(156379))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
func F_pq_getmsgstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = v5 + v6
	if v7&int32(3) == int32(0) {
		v31 = v7
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v65 = v64 + v5
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 <= v65 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v64 = v56 - v7
	goto L1
L3:
	;
	v35 = v31
	goto L12
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v20 = v7
	goto L8
L8:
	;
	v24 = v20 + int32(1)
	if v24&int32(3) == int32(0) {
		v31 = v24
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v56 = v24
	goto L2
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 != 0 {
		v20 = v24
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v44 = int32(-2139062144)
	if (int32(16843008)-v41|v41)&v44 == v44 {
		v35 = v35 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v50 = v35
	goto L15
L14:
	;
	goto L13
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		v50 = v50 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v56 = v50
	goto L2
L17:
	;
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v65 + int32(1)
	v89 = F_pg_client_to_server(m, v7, v64)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L21
	} else {
		goto L26
	}
L21:
	;
	return int32(0)
L22:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(398629), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(485954), int32(595), int32(324719))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	return v89
}
func F_pq_getmsgtext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	if l1 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L31
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7-v8 < l1 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v13 + v8
	v15 = F_pg_client_to_server(m, v14, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v15 != v14 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v15&int32(3) == int32(0) {
		v43 = v15
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L8
L8:
	;
	v81 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L26
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v76
	return v15
L10:
	;
	v76 = v68 - v15
	goto L9
L11:
	;
	v47 = v43
	goto L20
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v27 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v76 = int32(0)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v32 = v15
	goto L16
L16:
	;
	v36 = v32 + int32(1)
	if v36&int32(3) == int32(0) {
		v43 = v36
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v68 = v36
	goto L10
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v41 != 0 {
		v32 = v36
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v56 = int32(-2139062144)
	if (int32(16843008)-v53|v53)&v56 == v56 {
		v47 = v47 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v62 = v47
	goto L23
L22:
	;
	goto L21
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 != 0 {
		v62 = v62 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v68 = v62
	goto L10
L25:
	;
	goto L24
L26:
	;
	if l1 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84+l1))) = uint8(v86)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = l1
	return v84
L28:
	;
	v83 = F__emscripten_memcpy_bulkmem(m, v81, v14, l1)
	mBase = m.M
	v84 = v83
	goto L30
L29:
	;
	v84 = v81
	goto L30
L30:
	;
	goto L27
L31:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(398571), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(485954), int32(554), int32(62070))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v9 = F_palloc0(m, int32(524))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v19 != 0 {
			v20 = F__emscripten_memcpy_bulkmem(m, v9+int32(144), l0+int32(4), v19)
			mBase = m.M
		} else {
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = int32(128)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+272)) = v22
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(12)))))
		if v28 != int32(1) {
			v31 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v31
			v38 = *(*int32)(unsafe.Add(mBase, _consts[575]))
			v40 = m.G0
			v42 = v40 - int32(16)
			m.G0 = v42
			*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v38
			if v9 == int32(0) {
			} else {
				v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v47 == int32(1) {
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+400))
					if v38 == v50 {
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+384))
						if int32(0) < v52 {
							if v38 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+384))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v60
							} else {
							}
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+400)) = v63
						} else {
							v55 = F_pq_getkeepalivesidle(m, v9)
							mBase = m.M
							if int32(0) <= v55 {
								if v38 == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+384))
									*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v60
								} else {
								}
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+400)) = v63
							} else {
							}
						}
					}
				}
			}
			m.G0 = v42 + int32(16)
			v69 = *(*int32)(unsafe.Add(mBase, _consts[576]))
			v71 = m.G0
			v73 = v71 - int32(16)
			m.G0 = v73
			*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v69
			if v9 == int32(0) {
			} else {
				v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v78 == int32(1) {
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+404))
					if v69 == v81 {
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+388))
						if int32(0) < v83 {
							if v69 == int32(0) {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+388))
								*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v91
							} else {
							}
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+404)) = v94
						} else {
							v86 = F_pq_getkeepalivesinterval(m, v9)
							mBase = m.M
							if int32(0) <= v86 {
								if v69 == int32(0) {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+388))
									*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v91
								} else {
								}
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+404)) = v94
							} else {
							}
						}
					}
				}
			}
			m.G0 = v73 + int32(16)
			v100 = *(*int32)(unsafe.Add(mBase, _consts[577]))
			v102 = m.G0
			v104 = v102 - int32(16)
			m.G0 = v104
			*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v100
			if v9 == int32(0) {
			} else {
				v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v109 == int32(1) {
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+408))
					if v100 == v112 {
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
						if int32(0) < v114 {
							if v100 == int32(0) {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
								*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v122
							} else {
							}
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+408)) = v125
						} else {
							v117 = F_pq_getkeepalivescount(m, v9)
							mBase = m.M
							if int32(0) <= v117 {
								if v100 == int32(0) {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
									*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v122
								} else {
								}
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+408)) = v125
							} else {
							}
						}
					}
				}
			}
			m.G0 = v104 + int32(16)
			v131 = *(*int32)(unsafe.Add(mBase, _consts[578]))
			v133 = m.G0
			v135 = v133 - int32(16)
			m.G0 = v135
			*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v131
			if v9 == int32(0) {
			} else {
				v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v140 == int32(1) {
				} else {
					v143 = *(*int32)(unsafe.Add(mBase, uint32(v9)+412))
					if v131 == v143 {
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v9)+396))
						if int32(0) < v145 {
							if v131 == int32(0) {
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+396))
								*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v153
							} else {
							}
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+412)) = v156
						} else {
							v148 = F_pq_gettcpusertimeout(m, v9)
							mBase = m.M
							if int32(0) <= v148 {
								if v131 == int32(0) {
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+396))
									*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v153
								} else {
								}
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+412)) = v156
							} else {
							}
						}
					}
				}
			}
			m.G0 = v135 + int32(16)
		} else {
		}
		v162 = int32(8192)
		*(*int32)(unsafe.Add(mBase, _consts[579])) = v162
		v166 = *(*int32)(unsafe.Add(mBase, _consts[36]))
		v168 = F_MemoryContextAlloc(m, v166, v162)
		mBase = m.M
		v169 = m.ExcPending
		if v169 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[580])) = v168
			v172 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[581])) = v172
			*(*int32)(unsafe.Add(mBase, _consts[582])) = v172
			*(*int32)(unsafe.Add(mBase, _consts[583])) = v172
			*(*int32)(unsafe.Add(mBase, _consts[584])) = v172
			*(*uint8)(unsafe.Add(mBase, _consts[585])) = uint8(v172)
			*(*uint8)(unsafe.Add(mBase, _consts[586])) = uint8(v172)
			F_on_proc_exit(m, int32(800))
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return int32(0)
			} else {
				v194 = m.G0
				v196 = v194 - int32(16)
				m.G0 = v196
				*(*int32)(unsafe.Add(mBase, uint32(v196))) = int32(2048)
				m.G0 = v196 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(1)
				v217 = F_CreateWaitEventSet(m, int32(0), int32(3))
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[587])) = v217
					v221 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					F_AddWaitEventToSet(m, v217, int32(4), v221, int32(0))
					mBase = m.M
					v224 = m.ExcPending
					if v224 != 0 {
						return int32(0)
					} else {
						v226 = *(*int32)(unsafe.Add(mBase, _consts[587]))
						v230 = *(*int32)(unsafe.Add(mBase, _consts[496]))
						F_AddWaitEventToSet(m, v226, int32(1), int32(-1), v230)
						mBase = m.M
						v232 = m.ExcPending
						if v232 != 0 {
							return int32(0)
						} else {
							v234 = *(*int32)(unsafe.Add(mBase, _consts[587]))
							F_AddWaitEventToSet(m, v234, int32(16), int32(-1), int32(0))
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 - int32(-64)
								return v9
							}
						}
					}
				}
			}
		}
	}
}
func F_pq_parse_errornotice(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l1&int32(3) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(21)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v42
	v44 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v15 = l1 + int32(100)
	if base.Ui32(v15) <= base.Ui32(l1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v36 = F__emscripten_memset_bulkmem(m, l1+int32(4), base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L10
L5:
	;
	v21 = l1 + int32(4)
	if base.Ui32(v21) < base.Ui32(v15) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = v15
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	v30 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), (l1^int32(-1)+v23)&int32(-4)+int32(4))
	mBase = m.M
	goto L9
L9:
	;
	goto L1
L10:
	;
	goto L1
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L12
	} else {
		goto L169
	}
L12:
	;
	return
L13:
	;
	v47 = v44 << (uint(int32(24)) % 32)
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v52 = v47
	goto L17
L15:
	;
	goto L16
L16:
	;
	F_pq_getmsgend(m, l0)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L12
	} else {
		goto L168
	}
L17:
	;
	v53 = F_pq_getmsgrawstring(m, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v56 = v52 >> (uint(int32(24)) % 32)
	switch v56 - int32(67) {
	case 0:
		goto L37
	case 1:
		goto L36
	default:
		goto L22
	case 3:
		goto L25
	case 5:
		goto L35
	case 9:
		goto L24
	case 10:
		goto L21
	case 13:
		goto L34
	case 15:
		goto L23
	case 16:
		goto L20
	case 19:
		goto L38
	case 20:
		goto L31
	case 32:
		goto L28
	case 33:
		goto L27
	case 43:
		goto L26
	case 45:
		goto L33
	case 46:
		goto L32
	case 48:
		goto L30
	case 49:
		goto L29
	}
L20:
	;
	v469 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L166
	}
L21:
	;
	v466 = F_pstrdup(m, v53)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L12
	} else {
		goto L165
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L12
	} else {
		goto L162
	}
L23:
	;
	v450 = F_pstrdup(m, v53)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L12
	} else {
		goto L161
	}
L24:
	;
	v447 = F_pg_strtoint32(m, v53)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L12
	} else {
		goto L160
	}
L25:
	;
	v444 = F_pstrdup(m, v53)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L12
	} else {
		goto L159
	}
L26:
	;
	v441 = F_pstrdup(m, v53)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L158
	}
L27:
	;
	v438 = F_pstrdup(m, v53)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L12
	} else {
		goto L157
	}
L28:
	;
	v435 = F_pstrdup(m, v53)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L12
	} else {
		goto L156
	}
L29:
	;
	v432 = F_pstrdup(m, v53)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L12
	} else {
		goto L155
	}
L30:
	;
	v429 = F_pstrdup(m, v53)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L12
	} else {
		goto L154
	}
L31:
	;
	v426 = F_pstrdup(m, v53)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L12
	} else {
		goto L153
	}
L32:
	;
	v423 = F_pstrdup(m, v53)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L12
	} else {
		goto L152
	}
L33:
	;
	v420 = F_pg_strtoint32(m, v53)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L12
	} else {
		goto L151
	}
L34:
	;
	v417 = F_pg_strtoint32(m, v53)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L12
	} else {
		goto L150
	}
L35:
	;
	v414 = F_pstrdup(m, v53)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L12
	} else {
		goto L149
	}
L36:
	;
	v411 = F_pstrdup(m, v53)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L12
	} else {
		goto L148
	}
L37:
	;
	if v53&int32(3) == int32(0) {
		v337 = v53
		goto L132
	} else {
		goto L133
	}
L38:
	;
	v59 = int32(527861)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _consts[591])))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v63 == int32(0) {
		v82 = v62
		v83 = v63
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v83-v82 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	goto L39
L41:
	;
	if v62 != v63 {
		v82 = v62
		v83 = v63
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v67 = v53
	v68 = v59
	goto L43
L43:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v82 = v71
		v83 = v72
		goto L40
	} else {
		goto L45
	}
L44:
	;
	v82 = v71
	v83 = v72
	goto L40
L45:
	;
	v75 = int32(1)
	if v71 == v72 {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(14)
	goto L20
L48:
	;
	goto L49
L49:
	;
	v89 = int32(527977)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[592])))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v93 == int32(0) {
		v112 = v92
		v113 = v93
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v113-v112 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	goto L50
L52:
	;
	if v92 != v93 {
		v112 = v92
		v113 = v93
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v97 = v53
	v98 = v89
	goto L54
L54:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v101
		v113 = v102
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v112 = v101
	v113 = v102
	goto L51
L56:
	;
	v105 = int32(1)
	if v101 == v102 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(15)
	goto L20
L59:
	;
	goto L60
L60:
	;
	v119 = int32(519486)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, _consts[593])))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v123 == int32(0) {
		v142 = v122
		v143 = v123
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v143-v142 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L62:
	;
	goto L61
L63:
	;
	if v122 != v123 {
		v142 = v122
		v143 = v123
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v127 = v53
	v128 = v119
	goto L65
L65:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	if v132 == int32(0) {
		v142 = v131
		v143 = v132
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v142 = v131
	v143 = v132
	goto L62
L67:
	;
	v135 = int32(1)
	if v131 == v132 {
		v127 = v127 + v135
		v128 = v128 + v135
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(17)
	goto L20
L70:
	;
	goto L71
L71:
	;
	v149 = int32(534014)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _consts[594])))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v153 == int32(0) {
		v172 = v152
		v173 = v153
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v173-v172 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L73:
	;
	goto L72
L74:
	;
	if v152 != v153 {
		v172 = v152
		v173 = v153
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v157 = v53
	v158 = v149
	goto L76
L76:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v162 == int32(0) {
		v172 = v161
		v173 = v162
		goto L73
	} else {
		goto L78
	}
L77:
	;
	v172 = v161
	v173 = v162
	goto L73
L78:
	;
	v165 = int32(1)
	if v161 == v162 {
		v157 = v157 + v165
		v158 = v158 + v165
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	goto L20
L81:
	;
	goto L82
L82:
	;
	v179 = int32(528317)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _consts[595])))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v183 == int32(0) {
		v202 = v182
		v203 = v183
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v203-v202 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	goto L83
L85:
	;
	if v182 != v183 {
		v202 = v182
		v203 = v183
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v187 = v53
	v188 = v179
	goto L87
L87:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v192 == int32(0) {
		v202 = v191
		v203 = v192
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v202 = v191
	v203 = v192
	goto L84
L89:
	;
	v195 = int32(1)
	if v191 == v192 {
		v187 = v187 + v195
		v188 = v188 + v195
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(19)
	goto L20
L92:
	;
	goto L93
L93:
	;
	v209 = int32(516995)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _consts[596])))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v213 == int32(0) {
		v232 = v212
		v233 = v213
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v233-v232 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	goto L94
L96:
	;
	if v212 != v213 {
		v232 = v212
		v233 = v213
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v217 = v53
	v218 = v209
	goto L98
L98:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v222 == int32(0) {
		v232 = v221
		v233 = v222
		goto L95
	} else {
		goto L100
	}
L99:
	;
	v232 = v221
	v233 = v222
	goto L95
L100:
	;
	v225 = int32(1)
	if v221 == v222 {
		v217 = v217 + v225
		v218 = v218 + v225
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(21)
	goto L20
L103:
	;
	goto L104
L104:
	;
	v239 = int32(525789)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _consts[597])))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v243 == int32(0) {
		v262 = v242
		v263 = v243
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v263-v262 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	goto L105
L107:
	;
	if v242 != v243 {
		v262 = v242
		v263 = v243
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v247 = v53
	v248 = v239
	goto L109
L109:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	if v252 == int32(0) {
		v262 = v251
		v263 = v252
		goto L106
	} else {
		goto L111
	}
L110:
	;
	v262 = v251
	v263 = v252
	goto L106
L111:
	;
	v255 = int32(1)
	if v251 == v252 {
		v247 = v247 + v255
		v248 = v248 + v255
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(22)
	goto L20
L114:
	;
	goto L115
L115:
	;
	v269 = int32(536119)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _consts[598])))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v273 == int32(0) {
		v292 = v272
		v293 = v273
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v293-v292 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	goto L116
L118:
	;
	if v272 != v273 {
		v292 = v272
		v293 = v273
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v277 = v53
	v278 = v269
	goto L120
L120:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	if v282 == int32(0) {
		v292 = v281
		v293 = v282
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v292 = v281
	v293 = v282
	goto L117
L122:
	;
	v285 = int32(1)
	if v281 == v282 {
		v277 = v277 + v285
		v278 = v278 + v285
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(23)
	goto L20
L125:
	;
	goto L126
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v53
	F_errmsg_internal(m, int32(700689), v8+int32(16))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(488009), int32(272), int32(411798))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	if v370 != int32(5) {
		goto L11
	} else {
		goto L147
	}
L131:
	;
	v370 = v362 - v53
	goto L130
L132:
	;
	v341 = v337
	goto L141
L133:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v321 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v370 = int32(0)
	goto L130
L135:
	;
	goto L136
L136:
	;
	v326 = v53
	goto L137
L137:
	;
	v330 = v326 + int32(1)
	if v330&int32(3) == int32(0) {
		v337 = v330
		goto L132
	} else {
		goto L139
	}
L138:
	;
	v362 = v330
	goto L131
L139:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if v335 != 0 {
		v326 = v330
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v350 = int32(-2139062144)
	if (int32(16843008)-v347|v347)&v350 == v350 {
		v341 = v341 + int32(4)
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v356 = v341
	goto L144
L143:
	;
	goto L142
L144:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if v360 != 0 {
		v356 = v356 + int32(1)
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v362 = v356
	goto L131
L146:
	;
	goto L145
L147:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v374 = int32(16)
	v376 = int32(63)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+3)))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = (v373+v374)&v376 | (v378+v374)&v376<<(uint(int32(6))%32) | (v386+v374)&v376<<(uint(int32(12))%32) | (v394+v374)&v376<<(uint(int32(18))%32) | (v402+v374)&v376<<(uint(int32(24))%32)
	goto L20
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v411
	goto L20
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v414
	goto L20
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v417
	goto L20
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v420
	goto L20
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v423
	goto L20
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v426
	goto L20
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v429
	goto L20
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v432
	goto L20
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v435
	goto L20
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v438
	goto L20
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v441
	goto L20
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v444
	goto L20
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v447
	goto L20
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v450
	goto L20
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v56
	F_errmsg_internal(m, int32(479425), v8)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L12
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(488009), int32(326), int32(411798))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L12
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v466
	goto L20
L166:
	;
	v472 = v469 << (uint(int32(24)) % 32)
	if v472 != 0 {
		v52 = v472
		goto L17
	} else {
		goto L167
	}
L167:
	;
	goto L18
L168:
	;
	m.G0 = v8 + int32(48)
	return
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v53
	F_errmsg_internal(m, int32(703736), v8+int32(32))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(488009), int32(276), int32(411798))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L12
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	*(*int32)(unsafe.Add(mBase, _consts[589])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[314])) = int32(1594392)
	*(*int32)(unsafe.Add(mBase, _consts[216])) = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[590])) = int32(196610)
	F_on_dsm_detach(m, l0, int32(807), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		return
	}
}
func F_pq_send_ascii_string(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = l1
	v8 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L16
	}
L4:
	;
	if base.I32_extend8_s(v8) < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v14 = int32(63)
	goto L8
L7:
	;
	v14 = v8
	goto L8
L8:
	;
	v15 = base.I32_extend8_s(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= v17+int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v36 = v7 + int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v37 != 0 {
		v7 = v36
		v8 = v37
		goto L4
	} else {
		goto L15
	}
L10:
	;
	F_appendStringInfoChar(m, l0, v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v17))) = uint8(v15)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v26 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v28))) = uint8(v32)
	goto L9
L13:
	;
	return
L14:
	;
	goto L9
L15:
	;
	goto L5
L16:
	;
	return
}
