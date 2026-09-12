package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MarkPortalActive(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v7 != int32(2) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v17
				F_errmsg(m, int32(244375), v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(497107), int32(401), int32(343021))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(3)
		v30 = *(*int32)(unsafe.Add(mBase, _consts[72]))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v31
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PortalDefineQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l3
	return
}
func F_PortalDrop(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v10 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L53
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v13 == int32(3) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L49
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	m.T0[v16].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_hash_search(m, v22, v23, int32(2), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L8
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	if v26 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if v30 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	F_errmsg_internal(m, int32(69662), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(497107), int32(515), int32(234368))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_ReleaseCachedPlan(m, v43, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v49 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(0)
	goto L20
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	F_UnregisterSnapshotFromOwner(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	goto L24
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v87 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v60 != int32(5) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v63 = l1
	goto L33
L32:
	;
	v63 = int32(0)
	goto L33
L33:
	;
	if v63 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v66 = base.B2i32(v60 != int32(5))
	F_ResourceOwnerRelease(m, v56, int32(1), v66, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ResourceOwnerRelease(m, v70, int32(2), v66, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ResourceOwnerRelease(m, v75, int32(3), v66, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ResourceOwnerDelete(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L29
L39:
	;
	v88 = int32(4515392)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v91
	F_tuplestore_end(m, v87)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v100 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	goto L41
L43:
	;
	F_MemoryContextDelete(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_MemoryContextDelete(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, l0)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v8 + int32(32)
	return
L49:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v118
	F_errmsg(m, int32(712398), v8)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(497107), int32(479), int32(234368))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
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
	F_errcode(m, int32(258))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v135
	F_errmsg(m, int32(712367), v8+int32(16))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(497107), int32(487), int32(234368))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PortalRunSelect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l3
	} else {
	}
	if l1 != 0 {
		v10 = int32(0)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
		v14 = v11 | base.B2i32(l2 <= v10)
		if v14&int32(1) != 0 {
			v17 = v10
		} else {
			v17 = l2
		}
		if v17 != int32(2147483647) {
			v21 = v17
		} else {
			v21 = int32(0)
		}
		v25 = (v14 ^ int32(-1)) & int32(1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		if v26 != 0 {
			v28 = F_RunFromStore(m, l0, v25, base.I64_extend_i32_u(v21), l3)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int64(0)
			} else {
				v42 = v28
				if v14&int32(1) != 0 {
					v122 = v42
					return v122
				} else {
					if v42 != int64(0) {
						v47 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v47)
					} else {
					}
					if base.Ui64(base.I64_extend_i32_u(v21)) <= base.Ui64(v42) {
						v52 = v21
					} else {
						v52 = int32(0)
					}
					if v52 == int32(0) {
						v55 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v55)
					} else {
					}
					v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v57 + v42
					return v42
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			F_PushActiveSnapshot(m, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				F_ExecutorRun(m, v8, v25, base.I64_extend_i32_u(v21))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+112))
					F_PopActiveSnapshot(m)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						v42 = v39
						if v14&int32(1) != 0 {
							v122 = v42
							return v122
						} else {
							if v42 != int64(0) {
								v47 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v47)
							} else {
							}
							if base.Ui64(base.I64_extend_i32_u(v21)) <= base.Ui64(v42) {
								v52 = v21
							} else {
								v52 = int32(0)
							}
							if v52 == int32(0) {
								v55 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v55)
							} else {
							}
							v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v57 + v42
							return v42
						}
					}
				}
			}
		}
	} else {
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
		if v61&int32(4) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(421772), int32(0))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int64(0)
					} else {
						F_errhint(m, int32(620285), int32(0))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(492310), int32(941), int32(110121))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
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
		} else {
			v64 = int32(0)
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
			v68 = v65 | base.B2i32(l2 <= v64)
			v70 = v68 & int32(1)
			if v70 != 0 {
				v71 = v64
			} else {
				v71 = l2
			}
			if v71 != int32(2147483647) {
				v75 = v71
			} else {
				v75 = int32(0)
			}
			v77 = v70 - int32(1)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v78 != 0 {
				v80 = F_RunFromStore(m, l0, v77, base.I64_extend_i32_u(v75), l3)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int64(0)
				} else {
					v92 = v80
					if v68&int32(1) != 0 {
						v122 = v92
						return v122
					} else {
						if v92 == int64(0) {
						} else {
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
							if v97 != int32(1) {
							} else {
								v100 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v100)
								v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v102 + int64(1)
							}
						}
						if base.Ui64(base.I64_extend_i32_u(v75)) <= base.Ui64(v92) {
							v109 = v75
						} else {
							v109 = int32(0)
						}
						if v109 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
							v114 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v114)
							return v92
						} else {
							v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v117 - v92
							v122 = v92
							return v122
						}
					}
				}
			} else {
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				F_PushActiveSnapshot(m, v82)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					F_ExecutorRun(m, v8, v77, base.I64_extend_i32_u(v75))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)+112))
						F_PopActiveSnapshot(m)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							v92 = v89
							if v68&int32(1) != 0 {
								v122 = v92
								return v122
							} else {
								if v92 == int64(0) {
								} else {
									v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
									if v97 != int32(1) {
									} else {
										v100 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v100)
										v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v102 + int64(1)
									}
								}
								if base.Ui64(base.I64_extend_i32_u(v75)) <= base.Ui64(v92) {
									v109 = v75
								} else {
									v109 = int32(0)
								}
								if v109 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
									v114 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v114)
									return v92
								} else {
									v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v117 - v92
									v122 = v92
									return v122
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_PortalRunUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	v7 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v10 == v7 {
		v18 = int32(1)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		switch v14 - int32(158) {
		case 0, 1, 45, 64, 65, 66, 67, 86, 88, 89:
			v18 = int32(0)
		default:
			v18 = int32(1)
		}
	}
	if v18 != 0 {
		v19 = F_GetTransactionSnapshot(m)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if l3 != 0 {
				v21 = F_RegisterSnapshot(m, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v21
					v24 = v21
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					F_PushActiveSnapshotWithLevel(m, v24, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, _consts[113]))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						v32 = v30
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v32
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						F_ProcessUtility(m, l1, v34, base.B2i32(v35 != int32(0)), l2^int32(1), v40, v41, l4, l5)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v45
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
							if v47 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _consts[113]))
								if base.B2i32(v51 != int32(0)) == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
									return
								} else {
									F_PopActiveSnapshot(m)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v24 = v19
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				F_PushActiveSnapshotWithLevel(m, v24, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[113]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v32 = v30
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v32
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					F_ProcessUtility(m, l1, v34, base.B2i32(v35 != int32(0)), l2^int32(1), v40, v41, l4, l5)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v45
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						if v47 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _consts[113]))
							if base.B2i32(v51 != int32(0)) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
								return
							} else {
								F_PopActiveSnapshot(m)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v32 = v7
		*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v32
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		F_ProcessUtility(m, l1, v34, base.B2i32(v35 != int32(0)), l2^int32(1), v40, v41, l4, l5)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v45
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v47 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, _consts[113]))
				if base.B2i32(v51 != int32(0)) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
					return
				} else {
					F_PopActiveSnapshot(m)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
						return
					}
				}
			}
		}
	}
}
