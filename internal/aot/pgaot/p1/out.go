package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortOutOfAnyTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[0]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[1]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v7
	goto L3
L2:
	;
	v10 = v9
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[2])) = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[3]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v15 = v13
	v16 = v14
	goto L4
L4:
	;
	switch v16 {
	case 0:
		goto L13
	case 1, 2, 3, 4, 5, 6, 9, 10:
		goto L12
	case 7, 8:
		goto L11
	case 11, 12, 13, 14, 17, 18:
		goto L10
	case 15, 16, 19:
		goto L9
	default:
		v52 = v15
		goto L6
	}
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[2])) = v59
	return
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	if v56 != 0 {
		v15 = v52
		v16 = v56
		goto L4
	} else {
		goto L26
	}
L7:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L25
	}
L8:
	;
	F_CleanupSubTransaction(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L24
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v32 == int32(0) {
		goto L8
	} else {
		goto L22
	}
L10:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L16
	} else {
		goto L21
	}
L11:
	;
	F_AtAbort_Portals(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L16
	} else {
		goto L20
	}
L12:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L16
	} else {
		goto L19
	}
L13:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	switch v19 {
	case 0:
		v52 = v15
		goto L6
	case 1:
		goto L15
	default:
		goto L14
	}
L14:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(2)
	goto L14
L16:
	;
	return
L17:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v52 = v15
	goto L6
L19:
	;
	goto L7
L20:
	;
	goto L7
L21:
	;
	goto L8
L22:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_AtSubAbort_Portals(m, v38, v36, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_AbortOutOfAnyTransaction[3]))
	v52 = v47
	goto L6
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	v52 = v15
	goto L6
L26:
	;
	goto L5
}
func F_outBitmapset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_appendStringInfoChar(m, l0, int32(98))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if int32(0) <= v71 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v71 = base.I32_ctz(v57) | v58<<(uint(int32(5))%32)
	goto L4
L6:
	;
	v71 = int32(-2)
	goto L4
L7:
	;
	v24 = base.I32_div_s(int32(0), int32(32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 <= v24 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = l1 + int32(8)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v35 = v32 & int32(-1)
	if v35 != 0 {
		v57 = v35
		v58 = v24
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v37 = v24 + int32(1)
	if v37 == v25 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v40 = v37
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28+v40<<(uint(int32(2))%32))))
	if v47 != 0 {
		v57 = v47
		v58 = v40
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L6
L13:
	;
	v49 = v40 + int32(1)
	if v49 != v25 {
		v40 = v49
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v77 = v71
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L33
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v77
	F_appendStringInfo(m, l0, int32(_a_F_outBitmapset_0), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	if l1 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if int32(0) <= v137 {
		v77 = v137
		goto L18
	} else {
		goto L32
	}
L22:
	;
	v137 = base.I32_ctz(v123) | v124<<(uint(int32(5))%32)
	goto L21
L23:
	;
	v137 = int32(-2)
	goto L21
L24:
	;
	v88 = v77 + int32(1)
	v90 = base.I32_div_s(v88, int32(32))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v91 <= v90 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v94 = l1 + int32(8)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v90<<(uint(int32(2))%32))))
	v101 = v98 & (int32(-1) << (uint(v88) % 32))
	if v101 != 0 {
		v123 = v101
		v124 = v90
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v103 = v90 + int32(1)
	if v103 == v91 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v106 = v103
	goto L28
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v94+v106<<(uint(int32(2))%32))))
	if v113 != 0 {
		v123 = v113
		v124 = v106
		goto L22
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v115 = v106 + int32(1)
	if v115 != v91 {
		v106 = v115
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L19
L33:
	;
	m.G0 = v7 + int32(16)
	return
}
func F_outToken(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	F_appendStringInfoString(m, l0, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L29
	}
L3:
	;
	v65 = int32(_a_F_outToken_0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v8 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v65 = int32(_a_F_outToken_1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	if base.B2i32(v8 == int32(34))|base.B2i32(v8 == int32(60))|base.B2i32(base.Ui32((v8-int32(48))&int32(255)) < base.Ui32(int32(10))) != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v43 = l1
	v44 = v41
	v45 = int32(92)
	goto L16
L10:
	;
	v41 = int32(1)
	goto L9
L11:
	;
	v41 = int32(0)
	goto L9
L12:
	;
	switch v8 - int32(43) {
	case 0, 2:
		goto L13
	default:
		goto L10
	}
L13:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v27 == int32(46) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(int32(9)) < base.Ui32((v27-int32(48))&int32(255)) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	if v44 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_appendStringInfoChar(m, l0, base.I32_extend8_s(v45))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	switch v52 {
	case 0:
		goto L1
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39:
		v61 = v52
		goto L23
	case 9, 10, 32, 40, 41:
		goto L24
	default:
		goto L25
	}
L21:
	;
	return
L22:
	;
	v44 = int32(1)
	goto L16
L23:
	;
	v43 = v43 + int32(1)
	v44 = int32(0)
	v45 = v61
	goto L16
L24:
	;
	F_appendStringInfoChar(m, l0, int32(92))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L21
	} else {
		goto L28
	}
L25:
	;
	switch v52 - int32(123) {
	case 0, 2:
		goto L24
	case 1:
		v61 = v52
		goto L23
	default:
		goto L26
	}
L26:
	;
	if v52 != int32(92) {
		v61 = v52
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v61 = v60
	goto L23
L29:
	;
	goto L1
}
