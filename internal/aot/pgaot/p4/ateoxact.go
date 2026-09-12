package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_RelationCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, _consts[1371])) = v2
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1372])))
	if v14 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[1373]))
	if int32(0) < v144 {
		goto L45
	} else {
		goto L46
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1374]))
	if v18 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[1375]))
	F_hash_seq_init(m, v8+int32(28), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L25
	}
L5:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = int32(44)
	goto L8
L7:
	;
	v23 = int32(32)
	goto L8
L8:
	;
	v25 = v2
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[1375]))
	v35 = int32(0)
	v37 = F_hash_search(m, v30, v25<<(uint(int32(2))%32)+int32(4431008), v35, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L1
L11:
	;
	v78 = v25 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[1374]))
	if v78 < v80 {
		v25 = v78
		goto L9
	} else {
		goto L24
	}
L12:
	;
	return
L13:
	;
	if v37 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41+v23)))
	v44 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+40)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v44
	if v43 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v50 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_RelationClearRelation(m, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v57 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L11
L20:
	;
	if v57 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v61 + int32(4)
	F_errmsg_internal(m, int32(82939), v8+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(476107), int32(3365), int32(221142))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	goto L10
L25:
	;
	v90 = F_hash_seq_search(m, v8+int32(28))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	if v90 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if l0 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v96 = int32(44)
	goto L30
L29:
	;
	v96 = int32(32)
	goto L30
L30:
	;
	v98 = v90
	goto L31
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102+v96)))
	v105 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v102)+40)) = v105
	*(*int64)(unsafe.Add(mBase, uint32(v102)+32)) = v105
	if v104 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L1
L33:
	;
	v136 = F_hash_seq_search(m, v8+int32(28))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L43
	}
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v111 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_RelationClearRelation(m, v102)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v118 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L39
	}
L38:
	;
	goto L33
L39:
	;
	if v118 == int32(0) {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v102)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v122 + int32(4)
	F_errmsg_internal(m, int32(82939), v8)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(476107), int32(3365), int32(221142))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	goto L33
L43:
	;
	if v136 != 0 {
		v98 = v136
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[1376]))
	if int32(0) < v148 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1372])) = uint8(v188)
	*(*int32)(unsafe.Add(mBase, _consts[1374])) = v188
	*(*int32)(unsafe.Add(mBase, _consts[1376])) = v188
	*(*int32)(unsafe.Add(mBase, _consts[1373])) = v188
	m.G0 = v8 + int32(48)
	return
L48:
	;
	v153 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[1377]))
	F_pfree(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L12
	} else {
		goto L55
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[1377]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158+v153<<(uint(int32(2))%32))))
	F_FreeTupleDesc(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v166 = v153 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, _consts[1376]))
	if v166 < v168 {
		v153 = v166
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1377])) = int32(0)
	goto L47
}
func F_AtEOXact_Snapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1448]))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pairingheap_remove(m, int32(4114840), v12+int32(52))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1448])) = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1449]))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v23 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[1447]))
	if v86 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v29 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1449])) = int32(0)
	goto L8
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v29<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v39 = F_unlink(m, v38)
	mBase = m.M
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	F_pairingheap_remove(m, int32(4114840), v61+int32(52))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L20
	}
L15:
	;
	v44 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v48
	F_errmsg_internal(m, int32(284163), v9+int32(16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(472259), int32(1057), int32(82832))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	v67 = v29 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v67 < v68 {
		v29 = v67
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	if l0 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	F_pairingheap_remove(m, int32(4114840), v86+int32(52))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1447])) = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	if v98 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[255]))
	if v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+40))
	v106 = *(*int32)(unsafe.Add(mBase, _consts[255]))
	v108 = v106 - int32(48)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v109))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v104)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v125 = int32(0)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v125
	v129 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+40)) = v125
	goto L22
L29:
	;
	if v121 == int32(0) {
		goto L22
	} else {
		goto L33
	}
L30:
	;
	v121 = base.B2i32(base.Ui32(v104) < base.Ui32(v109))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v121 = int32(base.Ui32(v104-v109) >> (uint(int32(31)) % 32))
	goto L29
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v125 = v124
	goto L28
L34:
	;
	v184 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[255])) = v184
	*(*int32)(unsafe.Add(mBase, _consts[254])) = v184
	*(*int32)(unsafe.Add(mBase, _consts[1450])) = v184
	*(*int32)(unsafe.Add(mBase, _consts[1451])) = v184
	*(*uint8)(unsafe.Add(mBase, _consts[481])) = uint8(v184)
	if l1 != 0 {
		goto L52
	} else {
		goto L53
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[255]))
	if v135 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	if v154 == int32(0) {
		goto L34
	} else {
		goto L42
	}
L37:
	;
	v140 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v140 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errmsg_internal(m, int32(221258), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(472259), int32(1075), int32(82832))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	v159 = v154
	goto L43
L43:
	;
	v165 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	goto L34
L45:
	;
	if v165 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v159
	F_errmsg_internal(m, int32(326243), v9)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v176 != 0 {
		v159 = v176
		goto L43
	} else {
		goto L51
	}
L49:
	;
	F_errfinish(m, int32(472259), int32(1079), int32(82832))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L44
L52:
	;
	v199 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v199
	v202 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+40)) = v199
	goto L54
L53:
	;
	goto L54
L54:
	;
	m.G0 = v9 + int32(32)
	return
}
func F_AtEOXact_TypeCache(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1394]))
	if v1 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v1
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1394])) = int32(0)
	m.G0 = v7 + int32(16)
	return
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1395]))
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1396]))
	v24 = int32(0)
	v26 = F_hash_search(m, v18, v20+v14<<(uint(int32(2))%32), v24, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v56 = v14 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1394]))
	if v56 < v58 {
		v14 = v56
		goto L4
	} else {
		goto L16
	}
L7:
	;
	return
L8:
	;
	if v26 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)))
	if v30 != int32(99) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+312))
	if v33&int32(-1572865) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1397]))
	v48 = F_hash_search(m, v42, v26+int32(16), int32(1), v7+int32(15))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v52
	goto L6
L16:
	;
	goto L5
}
