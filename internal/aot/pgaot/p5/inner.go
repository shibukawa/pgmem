package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WalkInnerWith(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v78 int32
	_ = v78
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v4 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v7 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v12 = F_lcons(m, v10, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v54 = F_lcons(m, int32(0), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L15
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v24 = v4
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v24<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v33 = F_makeDependencyGraphWalker(m, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v36 = v24 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 < v37 {
		v24 = v36
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v49 = F_list_delete_first(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v49
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v96 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l2)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L27
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v66 = v4
	goto L19
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v66<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v75 = F_makeDependencyGraphWalker(m, v74, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v80 = v78
	goto L24
L23:
	;
	v80 = int32(0)
	goto L24
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = F_lappend(m, v81, v73)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v82
	v86 = v66 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v86 < v87 {
		v66 = v86
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v99 = F_list_delete_first(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v99
	return
}
func F_inner_subltree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	v4 = int32(0)
	if l1|l2 < v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(139460), int32(0))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497293), int32(273), int32(411177))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
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
		if l2 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(139460), int32(0))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497293), int32(273), int32(411177))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
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
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v17 <= l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(139460), int32(0))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(497293), int32(273), int32(411177))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
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
				v20 = l0 + int32(8)
				if l2 < v17 {
					v22 = l2
				} else {
					v22 = v17
				}
				if v22 <= int32(0) {
					v177 = v20
					v186 = v20
				} else {
					v26 = v22 - int32(1)
					if v26 == int32(0) {
						v156 = v20
						v159 = v20
					} else {
						v29 = int32(3)
						v30 = v26 & v29
						if base.Ui32(v22-int32(2)) < base.Ui32(v29) {
							v117 = v20
							v119 = int32(0)
							v120 = v20
						} else {
							v39 = v20
							v41 = int32(0)
							v42 = v20
							v47 = v4
							for {
								v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
								v52 = int32(9)
								v54 = int32(131064)
								v56 = v39 + (v51+v52)&v54
								v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56))))
								v62 = v56 + (v57+v52)&v54
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
								v68 = v62 + (v63+v52)&v54
								v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68))))
								v74 = v68 + (v69+v52)&v54
								if v41|int32(1) == l1 {
									v78 = v56
								} else {
									v78 = v42
								}
								if v41|int32(2) == l1 {
									v82 = v62
								} else {
									v82 = v78
								}
								if v41|int32(3) == l1 {
									v86 = v68
								} else {
									v86 = v82
								}
								v88 = v41 + int32(4)
								if v88 == l1 {
									v90 = v74
								} else {
									v90 = v86
								}
								v92 = v47 + int32(4)
								if v92 != v26&int32(-4) {
									v39 = v74
									v41 = v88
									v42 = v90
									v47 = v92
									continue
								} else {
									break
								}
								break
							}
							v117 = v74
							v119 = v88
							v120 = v90
						}
						if v30 == int32(0) {
							v156 = v117
							v159 = v120
						} else {
							v131 = v117
							v133 = v119
							v134 = v120
							v138 = v4
							for {
								v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131))))
								v148 = v131 + (v143+int32(9))&int32(131064)
								v150 = v133 + int32(1)
								if v150 == l1 {
									v152 = v148
								} else {
									v152 = v134
								}
								v154 = v138 + int32(1)
								if v154 != v30 {
									v131 = v148
									v133 = v150
									v134 = v152
									v138 = v154
									continue
								} else {
									break
								}
								break
							}
							v156 = v148
							v159 = v152
						}
					}
					v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156))))
					v177 = v159
					v186 = v156 + (v168+int32(9))&int32(131064)
				}
				v187 = v186 - v177
				v189 = v187 + int32(8)
				v190 = F_palloc0(m, v189)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					v192 = v22 - l1
					*(*uint16)(unsafe.Add(mBase, uint32(v190)+4)) = uint16(v192)
					*(*int32)(unsafe.Add(mBase, uint32(v190))) = v189 << (uint(int32(2)) % 32)
					if v187 != 0 {
						v199 = F__emscripten_memcpy_bulkmem(m, v190+int32(8), v177, v187)
						mBase = m.M
					} else {
					}
					return v190
				}
			}
		}
	}
}
