package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckSetNamespace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v5 = F_get_namespace_name(m, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L77
	}
L3:
	;
	F_pfree(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L76
	}
L4:
	;
	return
L5:
	;
	if v5 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v7 = int32(_a_F_CheckSetNamespace_0)
	goto L11
L7:
	;
	goto L8
L8:
	;
	v107 = F_get_namespace_name(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L38
	}
L9:
	;
	if v45-v46 == int32(0) {
		v230 = v5
		goto L3
	} else {
		goto L22
	}
L11:
	;
	goto L12
L12:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v14 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v15 = v5
	v16 = v7
	v17 = int32(8)
	v18 = v14
	goto L17
L14:
	;
	v41 = v7
	v45 = int32(0)
	goto L15
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L9
L16:
	;
	v41 = v36
	v45 = v38
	goto L15
L17:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.B2i32(v18 != v20)|base.B2i32(v20 == int32(0)) != 0 {
		v36 = v16
		v38 = v18
		goto L16
	} else {
		goto L19
	}
L18:
	;
	v36 = v30
	v38 = int32(0)
	goto L16
L19:
	;
	v26 = v17 - int32(1)
	if v26 == int32(0) {
		v36 = v16
		v38 = v18
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v29 = int32(1)
	v30 = v16 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v31 != 0 {
		v15 = v15 + v29
		v16 = v30
		v17 = v26
		v18 = v31
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v56 = int32(_a_F_CheckSetNamespace_1)
	goto L25
L23:
	;
	F_pfree(m, v5)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L36
	}
L25:
	;
	goto L26
L26:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v63 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v64 = v5
	v65 = v56
	v66 = int32(14)
	v67 = v63
	goto L31
L28:
	;
	v90 = v56
	v94 = int32(0)
	goto L29
L29:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	goto L23
L30:
	;
	v90 = v85
	v94 = v87
	goto L29
L31:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if base.B2i32(v67 != v69)|base.B2i32(v69 == int32(0)) != 0 {
		v85 = v65
		v87 = v67
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v85 = v79
	v87 = int32(0)
	goto L30
L33:
	;
	v75 = v66 - int32(1)
	if v75 == int32(0) {
		v85 = v65
		v87 = v67
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v78 = int32(1)
	v79 = v65 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v80 != 0 {
		v64 = v64 + v78
		v65 = v79
		v66 = v75
		v67 = v80
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v94-v95 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	if v107 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v109 = int32(_a_F_CheckSetNamespace_0)
	goto L44
L40:
	;
	goto L41
L41:
	;
	v209 = int32(99)
	if base.B2i32(l0 != v209)&base.B2i32(l1 != v209) != 0 {
		goto L1
	} else {
		goto L71
	}
L42:
	;
	if v147-v148 == int32(0) {
		v230 = v107
		goto L3
	} else {
		goto L55
	}
L44:
	;
	goto L45
L45:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v116 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v117 = v107
	v118 = v109
	v119 = int32(8)
	v120 = v116
	goto L50
L47:
	;
	v143 = v109
	v147 = int32(0)
	goto L48
L48:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	goto L42
L49:
	;
	v143 = v138
	v147 = v140
	goto L48
L50:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if base.B2i32(v120 != v122)|base.B2i32(v122 == int32(0)) != 0 {
		v138 = v118
		v140 = v120
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v138 = v132
	v140 = int32(0)
	goto L49
L52:
	;
	v128 = v119 - int32(1)
	if v128 == int32(0) {
		v138 = v118
		v140 = v120
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v131 = int32(1)
	v132 = v118 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v133 != 0 {
		v117 = v117 + v131
		v118 = v132
		v119 = v128
		v120 = v133
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v158 = int32(_a_F_CheckSetNamespace_1)
	goto L58
L56:
	;
	F_pfree(m, v107)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L69
	}
L58:
	;
	goto L59
L59:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v165 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v166 = v107
	v167 = v158
	v168 = int32(14)
	v169 = v165
	goto L64
L61:
	;
	v192 = v158
	v196 = int32(0)
	goto L62
L62:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	goto L56
L63:
	;
	v192 = v187
	v196 = v189
	goto L62
L64:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if base.B2i32(v169 != v171)|base.B2i32(v171 == int32(0)) != 0 {
		v187 = v167
		v189 = v169
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v187 = v181
	v189 = int32(0)
	goto L63
L66:
	;
	v177 = v168 - int32(1)
	if v177 == int32(0) {
		v187 = v167
		v189 = v169
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v180 = int32(1)
	v181 = v167 + v180
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	if v182 != 0 {
		v166 = v166 + v180
		v167 = v181
		v168 = v177
		v169 = v182
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	if v196-v197 == int32(0) {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	goto L41
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_CheckSetNamespace_2), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_CheckSetNamespace_3), int32(3474), int32(_a_F_CheckSetNamespace_4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
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
	goto L2
L77:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_CheckSetNamespace_5), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_CheckSetNamespace_3), int32(3468), int32(_a_F_CheckSetNamespace_4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_namespace_name_or_temp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_get_namespace_name_or_temp[0]))
	if base.B2i32(v5 != int32(0))&base.B2i32(l0 == v5) != 0 {
		v11 = F_pstrdup(m, int32(_a_F_get_namespace_name_or_temp_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	} else {
		v17 = F_SearchSysCache1(m, int32(38), l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v28 = F_pstrdup(m, v23+v24+int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v17)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						return v28
					}
				}
			}
		}
	}
}
func F_get_namespace_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13901(m, l0, l1, int32(_a_F_get_namespace_oid_0), int32(3547), int32(_a_F_get_namespace_oid_1), int32(_a_F_get_namespace_oid_2), int32(1411), int32(37))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_report_namespace_conflict(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 <= int32(3599) {
		if l0 != int32(2607) {
			if l0 != int32(3381) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_report_namespace_conflict_0), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(144), int32(_a_F_report_namespace_conflict_2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = int32(_a_F_report_namespace_conflict_3)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_report_namespace_conflict_4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = F_get_namespace_name(m, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
							F_errmsg(m, v25, v7+int32(16))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(150), int32(_a_F_report_namespace_conflict_2))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
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
		} else {
			v25 = int32(_a_F_report_namespace_conflict_5)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_report_namespace_conflict_4))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(150), int32(_a_F_report_namespace_conflict_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
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
	} else {
		switch l0 - int32(3600) {
		case 0:
			v25 = int32(_a_F_report_namespace_conflict_6)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_report_namespace_conflict_4))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(150), int32(_a_F_report_namespace_conflict_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 1:
			v25 = int32(_a_F_report_namespace_conflict_7)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_report_namespace_conflict_4))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(150), int32(_a_F_report_namespace_conflict_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 2:
			v25 = int32(_a_F_report_namespace_conflict_8)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_report_namespace_conflict_4))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = F_get_namespace_name(m, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, v25, v7+int32(16))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(150), int32(_a_F_report_namespace_conflict_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
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
			if l0 != int32(3764) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_report_namespace_conflict_0), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(144), int32(_a_F_report_namespace_conflict_2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = int32(_a_F_report_namespace_conflict_9)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_report_namespace_conflict_4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = F_get_namespace_name(m, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
							F_errmsg(m, v25, v7+int32(16))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_report_namespace_conflict_1), int32(150), int32(_a_F_report_namespace_conflict_2))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
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
