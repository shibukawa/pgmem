package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswEntryCandidate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l3 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if l0 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
			v29 = v20
		} else {
			v21 = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
			if v22 == v21 {
				v29 = v21
			} else {
				v29 = l0 + v22 - int32(1)
			}
		}
		v30 = F_FunctionCall2Coll(m, v15, v16, v17, v29)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v43 = v30
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v43)))
			v46 = F_palloc(m, int32(40))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v46)+32)) = v44
				if l0 != 0 {
					v52 = l1 - l0 + int32(1)
				} else {
					v52 = l1
				}
				*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v52
				m.G0 = v11 + int32(16)
				return v46
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+80)))
		F_HnswLoadElementImpl(m, v35, v36, v11, l2, l3, l4, l5, int32(0), v11+int32(12))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			v43 = v11
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v43)))
			v46 = F_palloc(m, int32(40))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v46)+32)) = v44
				if l0 != 0 {
					v52 = l1 - l0 + int32(1)
				} else {
					v52 = l1
				}
				*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v52
				m.G0 = v11 + int32(16)
				return v46
			}
		}
	}
}
func F_HnswFormIndexValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		if v11 != 0 {
			m.T0[v11].(func(*base.Module, int32))(m, v7)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
				if v14 != 0 {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					v17 = F_FunctionCall1Coll(m, v14, v16, v7)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
						if base.F64_gt(v19, float64(0)) == int32(0) {
							v32 = int32(0)
							return v32
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
							v26 = F_DirectFunctionCall1Coll(m, v24, v25, v7)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								v29 = v26
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
								v32 = int32(1)
								return v32
							}
						}
					}
				} else {
					v29 = v7
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
					v32 = int32(1)
					return v32
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
			if v14 != 0 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
				v17 = F_FunctionCall1Coll(m, v14, v16, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
					if base.F64_gt(v19, float64(0)) == int32(0) {
						v32 = int32(0)
						return v32
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						v26 = F_DirectFunctionCall1Coll(m, v24, v25, v7)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v29 = v26
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
							v32 = int32(1)
							return v32
						}
					}
				}
			} else {
				v29 = v7
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
				v32 = int32(1)
				return v32
			}
		}
	}
}
func F_HnswGetM(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v2 == int32(0) {
		return int32(16)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		return v7
	}
}
func F_HnswInitPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v3 = int32(_a_F_HnswInitPage_0)
	v5 = int32(0)
	if v5|(l1&int32(3)|int32(1)) == v5 {
		v21 = l1 + v3
		v23 = l1 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l1^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l1, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l1, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(_a_F_HnswInitPage_1)
	v44 = int32(_a_F_HnswInitPage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v44)
	v50 = int32(_a_F_HnswInitPage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v54 = l1 + v53
	v55 = int32(_a_F_HnswInitPage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
	return
}
func F_HnswUpdateConnection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v125 int32
	_ = v125
	v9 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*float32)(unsafe.Add(mBase, uint32(v13)+8)) = l3
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = l2 - l0 + int32(1)
	goto L3
L2:
	;
	v20 = v9
	goto L3
L3:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v20
	goto L6
L5:
	;
	v21 = l2
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v23 < l4 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v13 + int32(16)
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23 + int32(1)
	v30 = l1 + v23*int32(12)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v33
	if l5 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v39
	if v39 < v23 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(-2)
	goto L7
L12:
	;
	v48 = int32(0)
	v54 = v9
	goto L15
L13:
	;
	v73 = v9
	goto L14
L14:
	;
	v76 = v13 + int32(4)
	v77 = F_lappend(m, v73, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L17
	} else {
		goto L20
	}
L15:
	;
	v59 = F_lappend(m, v54, l1+int32(8)+v48*int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v73 = v59
	goto L14
L17:
	;
	return
L18:
	;
	v62 = v48 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v62 < v63 {
		v48 = v62
		v54 = v59
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v82 = F_SelectNeighbors(m, l0, v77, l4, l7, l1+int32(4), v76, v13, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v84 == int32(0) {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v87 <= int32(0) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v91 = l1 + int32(8)
	v95 = int32(0)
	goto L24
L24:
	;
	v105 = v91 + v95*int32(12)
	if l0 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L7
L26:
	;
	v125 = v95 + int32(1)
	if v125 != v87 {
		v95 = v125
		goto L24
	} else {
		goto L34
	}
L27:
	;
	v116 = v91 + v95*int32(12)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v119
	if l5 == int32(0) {
		goto L7
	} else {
		goto L33
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v108 == v109 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v111 != v112 {
		goto L26
	} else {
		goto L32
	}
L31:
	;
	goto L26
L32:
	;
	goto L27
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v95
	goto L7
L34:
	;
	goto L25
}
