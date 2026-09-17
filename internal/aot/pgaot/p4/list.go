package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_list_append_unique_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v34
L2:
	;
	v30 = F_lappend(m, l0, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = v3
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11+v14<<(uint(int32(2))%32))))
	if v20 == l1 {
		v34 = l0
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v23 = v14 + int32(1)
	if v8 != v23 {
		v14 = v23
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v34 = v30
	goto L1
}
func F_list_concat_unique_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v3 = int32(0)
	if l1 == v3 {
		v69 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v69
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v69 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = l0
	v16 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v16<<(uint(int32(2))%32))))
	if v13 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v69 = v58
	goto L1
L6:
	;
	v66 = v16 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v66 < v67 {
		v13 = v58
		v16 = v66
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v54 = F_lappend_oid(m, v13, v24)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = int32(0)
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30+v34<<(uint(int32(2))%32))))
	if v42 == v24 {
		v58 = v13
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v45 = v34 + int32(1)
	if v27 != v45 {
		v34 = v45
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	return int32(0)
L15:
	;
	v58 = v54
	goto L6
L16:
	;
	goto L5
}
func F_list_deduplicate_oid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	v2 = int32(0)
	if l0 == v2 {
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v12 < int32(2) {
		} else {
			v15 = int32(1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v12 != int32(2) {
				v19 = int32(1)
				v20 = v12 - v19
				v25 = int32(0)
				v28 = v25
				v30 = v15
				v31 = v25
				for {
					v36 = int32(2)
					v38 = v16 + v30<<(uint(v36)%32)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v16+v28<<(uint(v36)%32))))
					if v39 != v43 {
						v46 = v28 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16+v46<<(uint(int32(2))%32)))) = v39
						v51 = v46
					} else {
						v51 = v28
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v16+v51<<(uint(int32(2))%32))))
					if v52 != v56 {
						v59 = v51 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16+v59<<(uint(int32(2))%32)))) = v52
						v64 = v59
					} else {
						v64 = v51
					}
					v65 = int32(2)
					v66 = v30 + v65
					v68 = v31 + v65
					if v68 != v20&int32(-2) {
						v28 = v64
						v30 = v66
						v31 = v68
						continue
					} else {
						break
					}
					break
				}
				if v20&v19 == int32(0) {
					v97 = v64
				} else {
					v73 = v64
					v75 = v66
					v81 = int32(2)
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v16+v75<<(uint(v81)%32))))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v16+v73<<(uint(v81)%32))))
					if v84 == v88 {
						v97 = v73
					} else {
						v91 = v73 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16+v91<<(uint(int32(2))%32)))) = v84
						v97 = v91
					}
				}
			} else {
				v73 = v2
				v75 = v15
				v81 = int32(2)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v16+v75<<(uint(v81)%32))))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v16+v73<<(uint(v81)%32))))
				if v84 == v88 {
					v97 = v73
				} else {
					v91 = v73 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v16+v91<<(uint(int32(2))%32)))) = v84
					v97 = v91
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97 + int32(1)
		}
	}
	return
}
func F_list_delete_nth_cell(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(1) {
		if l0+int32(16) != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v24 = (v6 + (l1 ^ int32(-1))) << (uint(int32(2)) % 32)
		if v24 != 0 {
			v27 = v5 + l1<<(uint(int32(2))%32)
			base.MemoryCopy(m, v27, v27+int32(4), v24)
		} else {
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 - int32(1)
		return l0
	}
}
func F_list_difference(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v3 = int32(0)
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v118
L2:
	;
	if l0 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L7
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= int32(0) {
		v118 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v18 = v3
	v19 = v3
	goto L9
L9:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = v23 + v19<<(uint(int32(2))%32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 < v28 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v118 = v65
	goto L1
L11:
	;
	v70 = v19 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v70 < v71 {
		v18 = v65
		v19 = v70
		goto L9
	} else {
		goto L22
	}
L12:
	;
	v33 = v22
	goto L15
L13:
	;
	v59 = v27
	goto L14
L14:
	;
	v60 = F_lappend(m, v18, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L21
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v33<<(uint(int32(2))%32))))
	v43 = F_equal(m, v42, v27)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v59 = v51
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	if v43 != 0 {
		v65 = v18
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v48 = v33 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v48 < v49 {
		v33 = v48
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v65 = v60
	goto L11
L22:
	;
	goto L10
L23:
	;
	return int32(0)
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v83 = v81 + int32(4)
	if v83 <= v80 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v86 = v80
	goto L28
L27:
	;
	v86 = v83
	goto L28
L28:
	;
	if v86&(v86-int32(1)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = int32(1) << (uint(int32(32)-base.I32_clz(v86)) % 32)
	goto L31
L30:
	;
	v93 = v86
	goto L31
L31:
	;
	v95 = v93 - int32(4)
	v100 = F_palloc(m, v95<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v77
	v106 = v100 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v106
	v109 = v81 << (uint(int32(2)) % 32)
	if v109 == int32(0) {
		v118 = v100
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v106, v112, v109)
	return v100
}
func F_list_free_deep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v3 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v21 != l0+int32(16) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+v7<<(uint(int32(2))%32))))
	F_pfree(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v16 = v7 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 < v17 {
		v7 = v16
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_pfree(m, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_pfree(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	goto L3
}
func F_list_member_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		v36 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v36
L5:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v16 = v10
	goto L8
L7:
	;
	v16 = v13
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = int32(0)
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17+v19<<(uint(int32(2))%32))))
	v28 = base.B2i32(v27 == l1)
	if v27 == l1 {
		v36 = v28
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v36 = v28
	goto L4
L11:
	;
	v30 = v19 + int32(1)
	if v30 != v16 {
		v19 = v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_list_union_int(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v3 = int32(0)
	if l0 == v3 {
		v49 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		v114 = v49
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = v14 + int32(4)
	if v16 <= v13 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = v13
	goto L5
L4:
	;
	v19 = v16
	goto L5
L5:
	;
	if v19&(v19-int32(1)) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = int32(1) << (uint(int32(32)-base.I32_clz(v19)) % 32)
	goto L8
L7:
	;
	v26 = v19
	goto L8
L8:
	;
	v28 = v26 - int32(4)
	v33 = F_palloc(m, v28<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v10
	v41 = v33 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v41
	v44 = v14 << (uint(int32(2)) % 32)
	if v44 == int32(0) {
		v49 = v33
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v41, v47, v44)
	v49 = v33
	goto L1
L12:
	;
	return v114
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v55 <= int32(0) {
		v114 = v49
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v60 = v49
	v64 = v3
	goto L15
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v64<<(uint(int32(2))%32))))
	if v60 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v114 = v103
	goto L12
L17:
	;
	v109 = v64 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v109 < v110 {
		v60 = v103
		v64 = v109
		goto L15
	} else {
		goto L26
	}
L18:
	;
	v99 = F_lappend_int(m, v60, v69)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L25
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v72 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v77 = int32(0)
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v75+v77<<(uint(int32(2))%32))))
	if v87 == v69 {
		v103 = v60
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v90 = v77 + int32(1)
	if v72 != v90 {
		v77 = v90
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v103 = v99
	goto L17
L26:
	;
	goto L16
}
