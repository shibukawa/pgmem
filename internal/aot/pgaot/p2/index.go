package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(2), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			if l1 != 0 {
				v109 = int32(0)
				m.G0 = v8 - int32(-64)
				return v109
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(53160), v8)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497548), int32(69), int32(465299))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v33 = v31 + v32
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+72)))
			if v34 != int32(105) {
				if l1 != 0 {
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v109 = int32(0)
						m.G0 = v8 - int32(-64)
						return v109
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(509883)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v33 + int32(4)
							F_errmsg(m, int32(191779), v6+int32(-16))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(497548), int32(84), int32(465299))
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
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
				if v59 == int32(0) {
					if l1 != 0 {
						F_ReleaseCatCache(m, v11)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v109 = int32(0)
							m.G0 = v8 - int32(-64)
							return v109
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v33 + int32(4)
								F_errmsg(m, int32(219747), v6+int32(-48))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497548), int32(100), int32(465299))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v84 = F_OidFunctionCall0Coll(m, v59)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							if v84 != 0 {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
								if v86 == int32(438) {
									v109 = v84
									m.G0 = v8 - int32(-64)
									return v109
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
										F_errmsg_internal(m, int32(108891), v6+int32(-32))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(497548), int32(43), int32(373127))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
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
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
									F_errmsg_internal(m, int32(108891), v6+int32(-32))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(497548), int32(43), int32(373127))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
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
				}
			}
		}
	}
}
func F_get_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v6)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+100)))
	v20 = F_build_index_paths(m, l0, l1, l2, l3, v16, int32(2), v12+int32(15))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	return
L3:
	;
	if v20 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = v6
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v32<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+109)))
	if v41 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_add_path(m, l1, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+110)))
	if v46 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v60 = v32 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v60 < v61 {
		v32 = v60
		goto L6
	} else {
		goto L19
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+64))
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v40)+104))
	if base.F64_lt(v50, float64(1)) == int32(0) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v56 = F_lappend(m, v55, v40)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56
	goto L12
L19:
	;
	goto L7
L20:
	;
	v73 = int32(0)
	v76 = F_build_index_paths(m, l0, l1, l2, l3, v73, int32(1), v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	m.G0 = v12 + int32(16)
	return
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v79 = F_list_concat(m, v78, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v79
	goto L22
}
func F_index_can_return(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v13 != v9 {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[123]))
		v17 = F_list_member_ptr(m, v16, v9)
		mBase = m.M
		v18 = v17
	} else {
		v18 = int32(1)
	}
	if v18 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
		if v22 != 0 {
			v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v28 = v23
				m.G0 = v7 + int32(16)
				return v28
			}
		} else {
			v28 = int32(0)
			m.G0 = v7 + int32(16)
			return v28
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v40 + int32(4)
				F_errmsg(m, int32(438436), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497018), int32(837), int32(244710))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
func F_index_check_primary_key(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L12
	} else {
		goto L57
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L53
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L49
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L46
	}
L6:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
	if v94 != 0 {
		goto L4
	} else {
		goto L31
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+131)))
	if v15 != int32(1) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v18 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	F_list_free(m, v18)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L30
	}
L12:
	;
	return
L13:
	;
	if v18 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v23 <= v22 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v28 = v22
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v28<<(uint(int32(2))%32))))
	v39 = F_SearchSysCache1(m, int32(34), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	F_list_free(m, v18)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L25
	}
L18:
	;
	if v39 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v44)+14)))
	F_ReleaseCatCache(m, v39)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v46 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = v28 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v53 <= v52 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L17
L24:
	;
	v28 = v52
	goto L16
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v64 + int32(4)
	F_errmsg(m, int32(439574), v8+int32(-16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(492295), int32(221), int32(21163))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	goto L6
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(0) < v95 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v103 = int32(0)
	v104 = v95
	goto L35
L33:
	;
	goto L34
L34:
	;
	m.G0 = v10 - int32(-64)
	return
L35:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1+int32(12)+v103<<(uint(int32(1))%32)))))
	if v111 == int32(0) {
		goto L3
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	if int32(0) <= v111 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v118 = F_SearchSysCache2(m, int32(7), v117, v111)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	v131 = v104
	goto L40
L40:
	;
	v134 = v103 + int32(1)
	if v134 < v131 {
		v103 = v134
		v104 = v131
		goto L35
	} else {
		goto L45
	}
L41:
	;
	if v118 == int32(0) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+86)))
	if v125 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_ReleaseCatCache(m, v118)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v131 = v130
	goto L40
L45:
	;
	goto L36
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v38
	F_errmsg_internal(m, int32(40101), v8+int32(-32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(492295), int32(168), int32(22532))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(157510), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(492295), int32(234), int32(21163))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(146431), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(492295), int32(251), int32(21163))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v111
	F_errmsg_internal(m, int32(46437), v10)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(492295), int32(262), int32(21163))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v124 + int32(4)
	F_errmsg(m, int32(532557), v8+int32(-48))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(492295), int32(269), int32(21163))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_compute_xid_horizon_for_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if l2 < int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l2^int32(-1))<<(uint(int32(2))%32))))
		v35 = v27
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		v35 = v29 + l2<<(uint(int32(13))%32) + int32(-8192)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l0
	if l2 < int32(0) {
		v40 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(l2^int32(-1))<<(uint(int32(6))%32))+16))
		v55 = v46
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+l2<<(uint(int32(6))%32)+int32(-64))+16))
		v55 = v54
	}
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(0)
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v55
	v64 = F_palloc(m, l4<<(uint(int32(3))%32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v64
		v71 = F_palloc(m, l4*int32(6))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v71
			if int32(0) < l4 {
				v79 = v58
				v81 = int32(0)
				for {
					v94 = v64 + v81<<(uint(int32(3))%32)
					v95 = int32(1)
					v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v81<<(uint(v95)%32)))))
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v98<<(uint(int32(2))%32)+(v35+int32(24))-int32(4))))
					v107 = v35 + v104&int32(32767)
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v108
					v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)) = uint16(v79)
					*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)) = uint16(v110)
					v115 = v71 + v81*int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v115)+2)) = v95
					*(*uint16)(unsafe.Add(mBase, uint32(v115))) = uint16(v98)
					v120 = v79 + v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v120
					v123 = v81 + v95
					if v123 != l4 {
						v79 = v120
						v81 = v123
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+76))
			v142 = m.T0[v141].(func(*base.Module, int32, int32) int32)(m, l1, v16+int32(4))
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				F_pfree(m, v144)
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
					F_pfree(m, v147)
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(32)
						return v142
					}
				}
			}
		}
	}
}
func F_index_fetch_heap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v3)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v13 != 0 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[126])))
		if v15&int32(1) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(336328), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(326492), int32(1218), int32(383551))
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
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+188))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
			v31 = m.T0[v30].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20, l0+int32(60), v23, l1, l0+int32(66), v8+int32(15))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v31 == int32(0) {
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v55 == int32(0) {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
					} else {
					}
					m.G0 = v8 + int32(16)
					return v31
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
					if v38 == int32(0) {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
						if v41 != int32(1) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v55 == int32(0) {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
							} else {
							}
							m.G0 = v8 + int32(16)
							return v31
						} else {
							F_pgstat_assoc_relation(m, v37)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
								v48 = v47
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
								if v55 == int32(0) {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
								} else {
								}
								m.G0 = v8 + int32(16)
								return v31
							}
						}
					} else {
						v48 = v38
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v55 == int32(0) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
						} else {
						}
						m.G0 = v8 + int32(16)
						return v31
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+188))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
		v31 = m.T0[v30].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20, l0+int32(60), v23, l1, l0+int32(66), v8+int32(15))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			if v31 == int32(0) {
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v55 == int32(0) {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
				} else {
				}
				m.G0 = v8 + int32(16)
				return v31
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
				if v38 == int32(0) {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
					if v41 != int32(1) {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v55 == int32(0) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
						} else {
						}
						m.G0 = v8 + int32(16)
						return v31
					} else {
						F_pgstat_assoc_relation(m, v37)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
							v48 = v47
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v55 == int32(0) {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
							} else {
							}
							m.G0 = v8 + int32(16)
							return v31
						}
					}
				} else {
					v48 = v38
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v55 == int32(0) {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
					} else {
					}
					m.G0 = v8 + int32(16)
					return v31
				}
			}
		}
	}
}
func F_index_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = l1 - int32(1)
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v15 {
		v22 = l2 + v14<<(uint(int32(4))%32) + int32(20)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int32(0) {
			v64 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = v64
				m.G0 = v9 + int32(16)
				return v70
			}
		} else {
			v28 = l0 + v23 + int32(8)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
			if v29 != int32(1) {
				v70 = v28
				m.G0 = v9 + int32(16)
				return v70
			} else {
				v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
				switch v32&int32(65535) - int32(1) {
				case 0:
					v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
					v70 = v37
					m.G0 = v9 + int32(16)
					return v70
				case 1:
					v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28))))
					v70 = v38
					m.G0 = v9 + int32(16)
					return v70
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
						F_errmsg_internal(m, int32(483136), v9)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(326451), int32(70), int32(67779))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v70 = v39
					m.G0 = v9 + int32(16)
					return v70
				}
			}
		}
	} else {
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if int32(base.Ui32(v55)>>(uint(v14)%32))&int32(1) != 0 {
			v64 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = v64
				m.G0 = v9 + int32(16)
				return v70
			}
		} else {
			v59 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v59)
			v70 = int32(0)
			m.G0 = v9 + int32(16)
			return v70
		}
	}
}
func F_index_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v19 != v15 {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[123]))
		v23 = F_list_member_ptr(m, v22, v15)
		mBase = m.M
		v24 = v23
	} else {
		v24 = int32(1)
	}
	if v24 == int32(0) {
		v28 = F_EstimateSnapshotSpace(m, l2)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = F_add_size(m, int32(32), v28)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l7)+8)) = v32
				v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(l7))) = v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(l7)+12)) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l7)+20)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(l7)+24)) = int64(0)
				v43 = l7 + int32(32)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v52 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v43)+16)) = uint8(v55)
				*(*uint8)(unsafe.Add(mBase, uint32(v43)+17)) = uint8(v54)
				*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v52
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v51
				if v54 != 0 {
					v62 = v50
				} else {
					v62 = int32(0)
				}
				if v55 != 0 {
					v63 = v62
				} else {
					v63 = v50
				}
				*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v65 != 0 {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v71 = F___memcpy(m, l7+int32(56), v68, v65<<(uint(int32(2))%32))
					mBase = m.M
				} else {
				}
				if int32(0) < v63 {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v75 = int32(2)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
					v84 = F___memcpy(m, v43+v74<<(uint(v75)%32)+int32(24), v80, v81<<(uint(v75)%32))
					mBase = m.M
				} else {
				}
				v88 = (v30 + int32(7)) & int32(-8)
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l7)+24)) = v88
					v93 = l5<<(uint(int32(3))%32) + int32(8)
					v94 = F_add_size(m, v88, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l7)+24))
						v97 = l7 + v96
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v97
						v101 = F__emscripten_memset_bulkmem(m, v97, base.I32_extend8_s(int32(0)), v93)
						mBase = m.M
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
						*(*int32)(unsafe.Add(mBase, uint32(v102))) = l5
						v109 = (v94 + int32(7)) & int32(-8)
						if l4 == int32(0) {
							m.G0 = v13 + int32(16)
							return
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+124))
							if v114 == int32(0) {
								m.G0 = v13 + int32(16)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l7)+28)) = v109
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+124))
								m.T0[v120].(func(*base.Module, int32))(m, v109+l7)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v13 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v109 = v88
					if l4 == int32(0) {
						m.G0 = v13 + int32(16)
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+124))
						if v114 == int32(0) {
							m.G0 = v13 + int32(16)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l7)+28)) = v109
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+124))
							m.T0[v120].(func(*base.Module, int32))(m, v109+l7)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return
							} else {
								m.G0 = v13 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v129 = m.ExcPending
		if v129 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v132 = m.ExcPending
			if v132 != 0 {
				return
			} else {
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v133 + int32(4)
				F_errmsg(m, int32(438436), v13)
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return
				} else {
					F_errfinish(m, int32(497018), int32(520), int32(341479))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
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
