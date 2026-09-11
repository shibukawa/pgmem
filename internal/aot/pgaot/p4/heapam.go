package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_index_fetch_end(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		F_ReleaseBuffer(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
			F_pfree(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_heapam_index_fetch_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		F_ReleaseBuffer(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_heapam_relation_copy_data(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	F_FlushRelationBuffers(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+118)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v17
	v22 = F_RelationCreateStorage(m, v7+int32(-16), v14, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v50 = v24
	goto L6
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v28
	v32 = F_smgropen(m, v7+int32(-32), v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52)+118)))
	F_RelationCopyStorage(m, v50, v22, int32(0), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = v48
	goto L6
L9:
	;
	v44 = v36
	goto L11
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	v44 = v42
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v44 + int32(1)
	goto L8
L12:
	;
	v59 = int32(1)
	goto L13
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_RelationDropStorage(m, l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L15:
	;
	v89 = v63
	goto L17
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v65
	v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v67
	v71 = F_smgropen(m, v7+int32(-48), v64)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v90 = F_smgrexists(m, v89, v59)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v89 = v87
	goto L17
L20:
	;
	v83 = v75
	goto L22
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	v83 = v81
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v83 + int32(1)
	goto L19
L23:
	;
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_smgrcreate(m, v22, v59, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v136 = v59 + int32(1)
	if v136 != int32(4) {
		v59 = v136
		goto L13
	} else {
		goto L44
	}
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+118)))
	if v96 != int32(112) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v105 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v59 != int32(3) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_log_smgrcreate(m, l1, v59)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	if v96 != int32(117) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L28
L35:
	;
	v129 = v105
	goto L37
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v109
	v111 = F_smgropen(m, v9, v106)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v130)+118)))
	F_RelationCopyStorage(m, v129, v22, v59, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L43
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v111
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+72))
	if v115 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v129 = v127
	goto L37
L40:
	;
	v123 = v115
	goto L42
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+76))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v111)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)+72))
	v123 = v121
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+72)) = v123 + int32(1)
	goto L39
L43:
	;
	goto L26
L44:
	;
	goto L14
L45:
	;
	F_smgrclose(m, v22)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v9 - int32(-64)
	return
}
func F_heapam_relation_set_new_filelocator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_set_new_filelocator[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v11
	v13 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v13
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v18
		v21 = F_RelationCreateStorage(m, v8, l2, int32(1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if l2 == int32(117) {
				F_smgrcreate(m, v21, int32(3), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_log_smgrcreate(m, l1, int32(3))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_smgrclose(m, v21)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				F_smgrclose(m, v21)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_heapam_relation_toast_am(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+84))
	return v3
}
