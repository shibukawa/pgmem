package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageBtreeSearch(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v15
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = int32(1)
	v21 = l0 - v14 + v20
	v24 = v21 + v13 - v20
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 == int32(430584521) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v175)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v175
	return
L4:
	;
	v33 = v24
	v34 = v15
	goto L7
L5:
	;
	v114 = int32(2)
	v115 = v24
	goto L6
L6:
	;
	v122 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if base.Ui32(int32(509)) < base.Ui32(v124) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v41 = v33 + int32(12)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v44 = int32(0)
	v48 = v42
	goto L9
L8:
	;
	v114 = v88 + int32(1)
	v115 = v104
	goto L6
L9:
	;
	if base.Ui32(v48) <= base.Ui32(v44) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.Ui32(v71) < base.Ui32(v42) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L10
L12:
	;
	v71 = v44
	goto L11
L13:
	;
	goto L14
L14:
	;
	v58 = int32(1)
	v59 = int32(base.Ui32(v44+v48) >> (uint(v58) % 32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v41+v59<<(uint(int32(3))%32))))
	v66 = base.B2i32(base.Ui32(l1) < base.Ui32(v65))
	if base.Ui32(l1) < base.Ui32(v65) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = v44
	goto L17
L16:
	;
	v67 = v59 + v58
	goto L17
L17:
	;
	if base.Ui32(l1) < base.Ui32(v65) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v59
	goto L20
L19:
	;
	v68 = v48
	goto L20
L20:
	;
	if l1 != v65 {
		v44 = v67
		v48 = v68
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v71 = v59
	goto L11
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41+v71<<(uint(int32(3))%32))))
	v82 = base.B2i32(v80 != l1)
	goto L24
L23:
	;
	v82 = int32(1)
	goto L24
L24:
	;
	if base.Ui32(int32(509)) < base.Ui32(v42) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v88 = v34 + int32(1)
	goto L27
L26:
	;
	v88 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v88
	v90 = int32(3)
	v93 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v33+v71<<(uint(v90)%32)-base.B2i32(v71 != v93)&v82<<(uint(v90)%32))+16))
	v102 = v21 + v99 - int32(1)
	if v99 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v104 = v102
	goto L30
L29:
	;
	v104 = v93
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v105 == int32(430584521) {
		v33 = v104
		v34 = v88
		goto L7
	} else {
		goto L31
	}
L31:
	;
	goto L8
L32:
	;
	v127 = v114
	goto L34
L33:
	;
	v127 = v122
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v127
	v130 = v115 + int32(12)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v132 = v122
	v136 = v131
	goto L35
L35:
	;
	if base.Ui32(v136) <= base.Ui32(v132) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v115
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if base.Ui32(v159) < base.Ui32(v165) {
		goto L48
	} else {
		goto L49
	}
L37:
	;
	goto L36
L38:
	;
	v159 = v132
	goto L37
L39:
	;
	goto L40
L40:
	;
	v146 = int32(1)
	v147 = int32(base.Ui32(v132+v136) >> (uint(v146) % 32))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130+v147<<(uint(int32(3))%32))))
	v154 = base.B2i32(base.Ui32(l1) < base.Ui32(v153))
	if base.Ui32(l1) < base.Ui32(v153) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v155 = v132
	goto L43
L42:
	;
	v155 = v147 + v146
	goto L43
L43:
	;
	if base.Ui32(l1) < base.Ui32(v153) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v156 = v147
	goto L46
L45:
	;
	v156 = v136
	goto L46
L46:
	;
	if l1 != v153 {
		v132 = v155
		v136 = v156
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v159 = v147
	goto L37
L48:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v130+v159<<(uint(int32(3))%32))))
	v173 = base.B2i32(l1 == v170)
	goto L50
L49:
	;
	v173 = int32(0)
	goto L50
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v173)
	return
}
func F_FreePageManagerPut(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v7 = F_FreePageManagerPutInternal(m, l0, l1, l2, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(l2) < base.Ui32(v7) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = F_FreePageBtreeCleanup(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v14 = v7
	goto L5
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v16) < base.Ui32(v14) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	if base.Ui32(v7) < base.Ui32(v10) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v13 = v10
	goto L9
L8:
	;
	v13 = v7
	goto L9
L9:
	;
	v14 = v13
	goto L5
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
	goto L12
L11:
	;
	goto L12
L12:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v19 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+548))
	if v22 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	return
L16:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v74
	goto L15
L17:
	;
	v26 = l0 + int32(36)
	v30 = int32(128)
	goto L20
L18:
	;
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = int32(0)
	v65 = v22
	goto L29
L20:
	;
	v34 = v30 - int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26+v34<<(uint(int32(2))%32))))
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v74 = int32(0)
	goto L16
L22:
	;
	v74 = v30
	goto L16
L23:
	;
	goto L24
L24:
	;
	v39 = int32(2)
	v40 = v30 - v39
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26+v40<<(uint(v39)%32))))
	if v44 != 0 {
		v74 = v34
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v46 = v30 - int32(3)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26+v46<<(uint(int32(2))%32))))
	if v50 != 0 {
		v74 = v40
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v52 = v30 - int32(4)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+v52<<(uint(int32(2))%32))))
	if v56 != 0 {
		v74 = v46
		goto L16
	} else {
		goto L27
	}
L27:
	;
	if v52 != 0 {
		v30 = v52
		goto L20
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	v68 = v65 + (l0 - v58 + int32(1))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+3))
	if base.Ui32(v64) < base.Ui32(v69) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v74 = v71
	goto L16
L31:
	;
	v71 = v69
	goto L33
L32:
	;
	v71 = v64
	goto L33
L33:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+11))
	if v72 != 0 {
		v64 = v71
		v65 = v72
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
}
func F_PageGetHeapFreeSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = int32(4)
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v8 = v6 - v7
	if v8 <= v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v5
	goto L3
L2:
	;
	v11 = v8
	goto L3
L3:
	;
	v13 = v11 - int32(4)
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v7) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v13
L8:
	;
	v25 = int32(base.Ui32(v7+int32(_a_F_PageGetHeapFreeSpace_0)) >> (uint(int32(2)) % 32))
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32(v25&int32(_a_F_PageGetHeapFreeSpace_1)) < base.Ui32(int32(291)) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v30&int32(1) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	goto L14
L14:
	;
	v40 = int32(1)
	goto L15
L15:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(20)+v40&int32(_a_F_PageGetHeapFreeSpace_1)<<(uint(int32(2))%32))+1)))
	if v49&int32(384) == int32(0) {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v55 = v40 + int32(1)
	v56 = int32(_a_F_PageGetHeapFreeSpace_1)
	if base.Ui32(v55&v56) <= base.Ui32(v25&v56) {
		v40 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_PageGetItemIdCareful_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v16 = l2 + l3<<(uint(int32(2))%32) + int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v21 = int32(base.Ui32(v17) >> (uint(int32(17)) % 32))
	if base.Ui32(v17&int32(_a_F_PageGetItemIdCareful_1_0)+v21) < base.Ui32(int32(_a_F_PageGetItemIdCareful_1_1)) {
		if base.B2i32(v21 == int32(0))|base.B2i32(v17&int32(_a_F_PageGetItemIdCareful_1_2) != int32(_a_F_PageGetItemIdCareful_1_3)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v83 + int32(4)
					F_errmsg(m, int32(_a_F_PageGetItemIdCareful_1_4), v10+int32(80))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						*(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = int32(base.Ui32(v94)>>(uint(int32(15))%32)) & int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(base.Ui32(v94) >> (uint(int32(17)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v94 & int32(_a_F_PageGetItemIdCareful_1_0)
						F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_1_5), v10+int32(48))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_PageGetItemIdCareful_1_6), int32(791), int32(_a_F_PageGetItemIdCareful_1_7))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
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
			m.G0 = v10 + int32(96)
			return v16
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v45 + int32(4)
				F_errmsg(m, int32(_a_F_PageGetItemIdCareful_1_8), v10+int32(32))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(base.Ui32(v54)>>(uint(int32(15))%32)) & int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(base.Ui32(v54) >> (uint(int32(17)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v54 & int32(_a_F_PageGetItemIdCareful_1_0)
					F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_1_5), v10)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_PageGetItemIdCareful_1_6), int32(775), int32(_a_F_PageGetItemIdCareful_1_7))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
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
func F_PageGetTempPageCopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v6 = v4 << (uint(int32(8)) % 32)
	v7 = F_palloc(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			base.MemoryCopy(m, v7, l0, v6)
		} else {
		}
		return v7
	}
}
