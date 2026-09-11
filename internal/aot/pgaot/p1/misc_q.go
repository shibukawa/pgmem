package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_queue_listen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	goto L1
L1:
	;
	v14 = int32(_a_F_queue_listen_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[1]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_queue_listen[1])) = v18
	if l1&int32(3) == int32(0) {
		v43 = l1
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v79 = F_palloc(m, v76+int32(5))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v76 = v68 - l1
	goto L2
L4:
	;
	v47 = v43
	goto L13
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v76 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v32 = l1
	goto L9
L9:
	;
	v36 = v32 + int32(1)
	if v36&int32(3) == int32(0) {
		v43 = v36
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v68 = v36
	goto L3
L11:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v41 != 0 {
		v32 = v36
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v56 = int32(-2139062144)
	if (int32(16843008)-v53|v53)&v56 == v56 {
		v47 = v47 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v62 = v47
	goto L16
L15:
	;
	goto L14
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 != 0 {
		v62 = v62 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v68 = v62
	goto L3
L18:
	;
	goto L17
L19:
	;
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = l0
	v83 = v79 + int32(4)
	if (l1^v83)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[3]))
	if v159 != 0 {
		goto L44
	} else {
		goto L45
	}
L22:
	;
	goto L21
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v137)
	if v137&int32(255) == int32(0) {
		goto L22
	} else {
		goto L38
	}
L24:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v136 = l1
	v137 = v89
	v138 = v83
	goto L23
L25:
	;
	goto L26
L26:
	;
	if l1&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v93 = l1
	v95 = v83
	goto L30
L28:
	;
	v107 = l1
	v109 = v83
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v114 = int32(-2139062144)
	if (int32(16843008)-v111|v111)&v114 != v114 {
		v136 = v107
		v137 = v111
		v138 = v109
		goto L23
	} else {
		goto L34
	}
L30:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v96)
	if v96 == int32(0) {
		goto L22
	} else {
		goto L32
	}
L31:
	;
	v107 = v103
	v109 = v101
	goto L29
L32:
	;
	v100 = int32(1)
	v101 = v95 + v100
	v103 = v93 + v100
	if v103&int32(3) != 0 {
		v93 = v103
		v95 = v101
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v119 = v107
	v120 = v111
	v121 = v109
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v120
	v123 = int32(4)
	v124 = v121 + v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v127 = v119 + v123
	v131 = int32(-2139062144)
	if (v125|(int32(16843008)-v125))&v131 == v131 {
		v119 = v127
		v120 = v125
		v121 = v124
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v136 = v127
	v137 = v125
	v138 = v124
	goto L23
L37:
	;
	goto L36
L38:
	;
	v145 = v136
	v147 = v138
	goto L39
L39:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)) = uint8(v148)
	v150 = int32(1)
	if v148 != 0 {
		v145 = v145 + v150
		v147 = v147 + v150
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L22
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_queue_listen[1])) = v15
	m.G0 = v9 + int32(16)
	return
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v182 = F_lappend(m, v181, v79)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L50
	}
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v13 <= v160 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[4]))
	v165 = F_MemoryContextAlloc(m, v163, int32(12))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L19
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v79
	v173 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v173
	v176 = int32(_a_F_queue_listen_1)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v177
	*(*int32)(unsafe.Add(mBase, _c_F_queue_listen[3])) = v165
	goto L42
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_queue_listen[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v182
	goto L42
}
func F_quickdie(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v24 = int32(_a_F_quickdie_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_quickdie[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_quickdie[0])) = v25 | int32(4)
	F_sigprocmask(m, int32(_a_F_quickdie_0), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		v35 = int32(_a_F_quickdie_1)
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_quickdie[1]))
		v38 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_quickdie[1])) = v37 + v38
		v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_quickdie[2])))
		if v42 != v38 {
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, _c_F_quickdie[3]))
			if v46 != int32(2) {
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_quickdie[3])) = int32(0)
			}
		}
		v53 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_quickdie[4])) = v53
		v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_quickdie[5])))
		if v58 != int32(1) {
			v66 = v53
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, _c_F_quickdie[6]))
			if v62 == int32(0) {
				v66 = v53
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
				v66 = v65
			}
		}
		switch v66 {
		case 0:
			v70 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				if v70 == int32(0) {
					F__Exit(m, int32(2))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					F_errcode(m, int32(16908741))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_quickdie_2), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v118 = int32(3089)
							F_errfinish(m, int32(_a_F_quickdie_3), v118, int32(_a_F_quickdie_4))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								F__Exit(m, int32(2))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 1:
			v84 = F_errstart(m, int32(20), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				if v84 == int32(0) {
					F__Exit(m, int32(2))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					F_errcode(m, int32(33685957))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_quickdie_5), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errdetail(m, int32(_a_F_quickdie_6), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_errhint(m, int32(_a_F_quickdie_7), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									v118 = int32(3101)
									F_errfinish(m, int32(_a_F_quickdie_3), v118, int32(_a_F_quickdie_4))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										F__Exit(m, int32(2))
										mBase = m.M
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
		case 2:
			v106 = F_errstart(m, int32(20), int32(0))
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return
			} else {
				if v106 == int32(0) {
					F__Exit(m, int32(2))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					F_errcode(m, int32(16908741))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_quickdie_8), int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v118 = int32(3107)
							F_errfinish(m, int32(_a_F_quickdie_3), v118, int32(_a_F_quickdie_4))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								F__Exit(m, int32(2))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		default:
			F__Exit(m, int32(2))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	F_initStringInfo(m, v6+int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v14 = F_quote_identifier(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
				F_appendStringInfo(m, v6+int32(16), int32(_a_F_quote_qualified_identifier_0), v6)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v24 = F_quote_identifier(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_appendStringInfoString(m, v6+int32(16), v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
							m.G0 = v6 + int32(32)
							return v28
						}
					}
				}
			}
		} else {
			v24 = F_quote_identifier(m, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_appendStringInfoString(m, v6+int32(16), v24)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
					m.G0 = v6 + int32(32)
					return v28
				}
			}
		}
	}
}
