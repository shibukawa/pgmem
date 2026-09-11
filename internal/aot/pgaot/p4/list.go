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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
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
	v15 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v15<<(uint(int32(2))%32))))
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
	v66 = v15 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v66 < v67 {
		v13 = v58
		v15 = v66
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
	v35 = int32(0)
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30+v35<<(uint(int32(2))%32))))
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
	v45 = v35 + int32(1)
	if v27 != v45 {
		v35 = v45
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	if l0 == int32(0) {
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v13 < int32(2) {
		} else {
			v16 = int32(1)
			v18 = v13 - v16
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v13 == int32(2) {
				v77 = int32(0)
				v79 = v16
			} else {
				v29 = int32(0)
				v32 = v29
				v34 = v16
				v35 = v29
				for {
					v41 = int32(2)
					v42 = v34 << (uint(v41) % 32)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v21+v42)))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v21+v32<<(uint(v41)%32))))
					if v44 != v48 {
						v51 = v32 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v21+v51<<(uint(int32(2))%32)))) = v44
						v56 = v51
					} else {
						v56 = v32
					}
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v21+int32(4)))))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v21+v56<<(uint(int32(2))%32))))
					if v58 != v62 {
						v65 = v56 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v21+v65<<(uint(int32(2))%32)))) = v58
						v70 = v65
					} else {
						v70 = v56
					}
					v71 = int32(2)
					v72 = v34 + v71
					v74 = v35 + v71
					if v74 != v18&int32(-2) {
						v32 = v70
						v34 = v72
						v35 = v74
						continue
					} else {
						break
					}
					break
				}
				v77 = v70
				v79 = v72
			}
			if v18&v16 == int32(0) {
				v103 = v77
			} else {
				v88 = int32(2)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v21+v79<<(uint(v88)%32))))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v21+v77<<(uint(v88)%32))))
				if v91 == v95 {
					v103 = v77
				} else {
					v98 = v77 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v21+v98<<(uint(int32(2))%32)))) = v91
					v103 = v98
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(1)
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0+int32(16) != v5 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = int32(2)
	v22 = v5 + l1<<(uint(v20)%32)
	v24 = v22 + int32(4)
	v29 = (v6 + (l1 ^ int32(-1))) << (uint(v20) % 32)
	if v22 == v24 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	F_pfree(m, v5)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_pfree(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v174 - int32(1)
	return l0
L11:
	;
	goto L10
L12:
	;
	v33 = v22 + v29
	if base.Ui32(v24-v33) <= base.Ui32(int32(0)-v29<<(uint(int32(1))%32)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = F___memcpy(m, v22, v24, v29)
	mBase = m.M
	goto L10
L14:
	;
	goto L15
L15:
	;
	v43 = (v22 ^ v24) & int32(3)
	if base.Ui32(v22) < base.Ui32(v24) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v145 == int32(0) {
		goto L11
	} else {
		goto L52
	}
L17:
	;
	if base.Ui32(v123) <= base.Ui32(int32(3)) {
		v144 = v122
		v145 = v123
		v146 = v124
		goto L16
	} else {
		goto L48
	}
L18:
	;
	if v43 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v43 != 0 {
		v105 = v29
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v144 = v24
	v145 = v29
	v146 = v22
	goto L16
L22:
	;
	goto L23
L23:
	;
	if v22&int32(3) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v122 = v24
	v123 = v29
	v124 = v22
	goto L17
L25:
	;
	goto L26
L26:
	;
	v50 = v24
	v51 = v29
	v52 = v22
	goto L27
L27:
	;
	if v51 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v122 = v59
	v123 = v61
	v124 = v63
	goto L17
L29:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v56)
	v58 = int32(1)
	v59 = v50 + v58
	v61 = v51 - v58
	v63 = v52 + v58
	if v63&int32(3) != 0 {
		v50 = v59
		v51 = v61
		v52 = v63
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v105 == int32(0) {
		goto L11
	} else {
		goto L44
	}
L32:
	;
	if v33&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v70 = v29
	goto L36
L34:
	;
	v85 = v29
	goto L35
L35:
	;
	if base.Ui32(v85) <= base.Ui32(int32(3)) {
		v105 = v85
		goto L31
	} else {
		goto L40
	}
L36:
	;
	if v70 == int32(0) {
		goto L11
	} else {
		goto L38
	}
L37:
	;
	v85 = v76
	goto L35
L38:
	;
	v76 = v70 - int32(1)
	v77 = v22 + v76
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v76))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v79)
	if v77&int32(3) != 0 {
		v70 = v76
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v92 = v85
	goto L41
L41:
	;
	v96 = v92 - int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v24+v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v96))) = v99
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v92 = v96
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v105 = v96
	goto L31
L43:
	;
	goto L42
L44:
	;
	v112 = v105
	goto L45
L45:
	;
	v116 = v112 - int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22+v116))) = uint8(v119)
	if v116 != 0 {
		v112 = v116
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L11
L47:
	;
	goto L46
L48:
	;
	v129 = v122
	v130 = v123
	v131 = v124
	goto L49
L49:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133
	v135 = int32(4)
	v136 = v129 + v135
	v138 = v131 + v135
	v140 = v130 - v135
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		v129 = v136
		v130 = v140
		v131 = v138
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v144 = v136
	v145 = v140
	v146 = v138
	goto L16
L51:
	;
	goto L50
L52:
	;
	v151 = v144
	v152 = v145
	v153 = v146
	goto L53
L53:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v155)
	v157 = int32(1)
	v162 = v152 - v157
	if v162 != 0 {
		v151 = v151 + v157
		v152 = v162
		v153 = v153 + v157
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L11
L55:
	;
	goto L54
}
func F_list_difference(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l0 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v12 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = v3
	v20 = v3
	goto L10
L8:
	;
	v78 = v3
	goto L9
L9:
	;
	return v78
L10:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = v23 + v19<<(uint(int32(2))%32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 < v28 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v78 = v67
	goto L9
L12:
	;
	v70 = v19 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v70 < v71 {
		v19 = v70
		v20 = v67
		goto L10
	} else {
		goto L23
	}
L13:
	;
	v33 = v22
	goto L16
L14:
	;
	v59 = v27
	goto L15
L15:
	;
	v60 = F_lappend(m, v20, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L18
	} else {
		goto L22
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v33<<(uint(int32(2))%32))))
	v43 = F_equal(m, v42, v27)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v59 = v51
	goto L15
L18:
	;
	return int32(0)
L19:
	;
	if v43 != 0 {
		v67 = v20
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v48 = v33 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v48 < v49 {
		v33 = v48
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v67 = v60
	goto L12
L23:
	;
	goto L11
L24:
	;
	return int32(0)
L25:
	;
	goto L26
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(8)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v91 = v89 + int32(4)
	if v91 <= v88 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v94 = v88
	goto L29
L28:
	;
	v94 = v91
	goto L29
L29:
	;
	if v94&(v94-int32(1)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v101 = int32(1) << (uint(int32(32)-base.I32_clz(v94)) % 32)
	goto L32
L31:
	;
	v101 = v94
	goto L32
L32:
	;
	v103 = v101 - int32(4)
	v108 = F_palloc(m, v103<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v85
	v114 = v108 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v118 = v89 << (uint(int32(2)) % 32)
	if v118 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	return v108
L35:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v114, v116, v118)
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L34
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
	var v35 int32
	_ = v35
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
		v35 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v35
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
		v35 = v28
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v35 = v28
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		v114 = v49
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v49 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = v14 + int32(4)
	if v16 <= v13 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = v13
	goto L7
L6:
	;
	v19 = v16
	goto L7
L7:
	;
	if v19&(v19-int32(1)) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = int32(1) << (uint(int32(32)-base.I32_clz(v19)) % 32)
	goto L10
L9:
	;
	v26 = v19
	goto L10
L10:
	;
	v28 = v26 - int32(4)
	v33 = F_palloc(m, v28<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v10
	v41 = v33 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = v14 << (uint(int32(2)) % 32)
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v49 = v33
	goto L1
L14:
	;
	v46 = F__emscripten_memcpy_bulkmem(m, v41, v43, v45)
	mBase = m.M
	goto L16
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	return v114
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v54 <= int32(0) {
		v114 = v49
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v60 = v49
	v63 = v3
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v63<<(uint(int32(2))%32))))
	if v60 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v114 = v103
	goto L17
L22:
	;
	v108 = v63 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v108 < v109 {
		v60 = v103
		v63 = v108
		goto L20
	} else {
		goto L31
	}
L23:
	;
	v98 = F_lappend_int(m, v60, v68)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L30
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v71 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v76 = int32(0)
	goto L26
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v74+v76<<(uint(int32(2))%32))))
	if v86 == v68 {
		v103 = v60
		goto L22
	} else {
		goto L28
	}
L27:
	;
	goto L23
L28:
	;
	v89 = v76 + int32(1)
	if v71 != v89 {
		v76 = v89
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v103 = v98
	goto L22
L31:
	;
	goto L21
}
