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
	var v34 float64
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
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
			v34 = *(*float64)(unsafe.Add(mBase, uint32(v30)))
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
			v46 = v34
			v48 = F_palloc(m, int32(40))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v48)+32)) = v46
				if l0 != 0 {
					v54 = l1 - l0 + int32(1)
				} else {
					v54 = l1
				}
				*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v54
				m.G0 = v11 + int32(16)
				return v48
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+80)))
		F_HnswLoadElementImpl(m, v37, v38, v11, l2, l3, l4, l5, int32(0), v11+int32(12))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			v46 = v44
			v48 = F_palloc(m, int32(40))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v48)+32)) = v46
				if l0 != 0 {
					v54 = l1 - l0 + int32(1)
				} else {
					v54 = l1
				}
				*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v54
				m.G0 = v11 + int32(16)
				return v48
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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	if l1&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l1, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(1572864)
	v35 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v35)
	v41 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v45 = l1 + v44
	v46 = int32(65424)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v46)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(-1)
	return
}
func F_HnswUpdateConnection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	v9 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*float32)(unsafe.Add(mBase, uint32(v14)+8)) = l3
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = l2 - l0 + int32(1)
	goto L3
L2:
	;
	v21 = v9
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
	v22 = v21
	goto L6
L5:
	;
	v22 = l2
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v24 < l4 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v14 + int32(16)
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24 + int32(1)
	v31 = l1 + v24*int32(12)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v34
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
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v40
	if v40 < v24 {
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
	v49 = int32(0)
	v55 = v9
	goto L15
L13:
	;
	v75 = v9
	goto L14
L14:
	;
	v80 = F_lappend(m, v75, v14+int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L20
	}
L15:
	;
	v61 = F_lappend(m, v55, l1+int32(8)+v49*int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v75 = v61
	goto L14
L17:
	;
	return
L18:
	;
	v64 = v49 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v64 < v65 {
		v49 = v64
		v55 = v61
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v82 = int32(4)
	v87 = F_SelectNeighbors(m, l0, v80, l4, l7, l1+v82, v14+v82, v14, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v89 == int32(0) {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v92 <= int32(0) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v96 = l1 + int32(8)
	v100 = int32(0)
	goto L24
L24:
	;
	v111 = v96 + v100*int32(12)
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
	v131 = v100 + int32(1)
	if v131 != v92 {
		v100 = v131
		goto L24
	} else {
		goto L34
	}
L27:
	;
	v122 = v96 + v100*int32(12)
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v122))) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v125
	if l5 == int32(0) {
		goto L7
	} else {
		goto L33
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v114 == v115 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v117 != v118 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v100
	goto L7
L34:
	;
	goto L25
}
