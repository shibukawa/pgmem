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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
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
	var v135 int32
	_ = v135
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
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(1)
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v177)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v177
	return
L2:
	;
	v20 = int32(1)
	v21 = l0 - v14 + v20
	v24 = v21 + v13 - v20
	if v24 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 == int32(430584521) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = int32(1)
	v39 = v24
	v41 = v31
	goto L7
L5:
	;
	v113 = int32(2)
	v115 = v24
	goto L6
L6:
	;
	v122 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if base.Ui32(int32(509)) < base.Ui32(v124) {
		goto L29
	} else {
		goto L30
	}
L7:
	;
	v47 = v39 + int32(12)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v50 = int32(0)
	v53 = v48
	goto L9
L8:
	;
	v113 = v94 + int32(1)
	v115 = v104
	goto L6
L9:
	;
	if base.Ui32(v53) <= base.Ui32(v50) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.Ui32(v78) < base.Ui32(v48) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L10
L12:
	;
	v78 = v50
	goto L11
L13:
	;
	goto L14
L14:
	;
	v64 = int32(1)
	v65 = int32(base.Ui32(v50+v53) >> (uint(v64) % 32))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v47+v65<<(uint(int32(3))%32))))
	v72 = base.B2i32(base.Ui32(l1) < base.Ui32(v71))
	if base.Ui32(l1) < base.Ui32(v71) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v73 = v50
	goto L17
L16:
	;
	v73 = v65 + v64
	goto L17
L17:
	;
	if base.Ui32(l1) < base.Ui32(v71) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v74 = v65
	goto L20
L19:
	;
	v74 = v53
	goto L20
L20:
	;
	if l1 != v71 {
		v50 = v73
		v53 = v74
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v78 = v65
	goto L11
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v47+v78<<(uint(int32(3))%32))))
	v88 = base.B2i32(v86 != l1)
	goto L24
L23:
	;
	v88 = int32(1)
	goto L24
L24:
	;
	if base.Ui32(int32(509)) < base.Ui32(v48) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = v41 + int32(1)
	goto L27
L26:
	;
	v94 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v94
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v78-base.B2i32(v78 != int32(0))&v88)<<(uint(int32(3))%32))+16))
	v104 = v21 - v31 + v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105 == int32(430584521) {
		v39 = v104
		v41 = v94
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	v127 = v113
	goto L31
L30:
	;
	v127 = v122
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v127
	v130 = v115 + int32(12)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v132 = v122
	v135 = v131
	goto L32
L32:
	;
	if base.Ui32(v135) <= base.Ui32(v132) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v115
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if base.Ui32(v160) < base.Ui32(v166) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	goto L33
L35:
	;
	v160 = v132
	goto L34
L36:
	;
	goto L37
L37:
	;
	v146 = int32(1)
	v147 = int32(base.Ui32(v132+v135) >> (uint(v146) % 32))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130+v147<<(uint(int32(3))%32))))
	v154 = base.B2i32(base.Ui32(l1) < base.Ui32(v153))
	if base.Ui32(l1) < base.Ui32(v153) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v155 = v132
	goto L40
L39:
	;
	v155 = v147 + v146
	goto L40
L40:
	;
	if base.Ui32(l1) < base.Ui32(v153) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v156 = v147
	goto L43
L42:
	;
	v156 = v135
	goto L43
L43:
	;
	if l1 != v153 {
		v132 = v155
		v135 = v156
		goto L32
	} else {
		goto L44
	}
L44:
	;
	v160 = v147
	goto L34
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v130+v160<<(uint(int32(3))%32))))
	v173 = base.B2i32(l1 == v171)
	goto L47
L46:
	;
	v173 = int32(0)
	goto L47
L47:
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
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
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v75
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
	v59 = l0 - v58
	v63 = int32(0)
	v64 = v59 + v22
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
	v75 = int32(0)
	goto L16
L22:
	;
	v75 = v30
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
		v75 = v34
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v46 = v30 - int32(3)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26+v46<<(uint(int32(2))%32))))
	if v50 != 0 {
		v75 = v40
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v52 = v30 - int32(4)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+v52<<(uint(int32(2))%32))))
	if v56 != 0 {
		v75 = v46
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
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if base.Ui32(v63) < base.Ui32(v67) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v75 = v69
	goto L16
L31:
	;
	v69 = v67
	goto L33
L32:
	;
	v69 = v63
	goto L33
L33:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if v70 == int32(0) {
		v75 = v69
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v73 = v70 + v59
	if v73 != 0 {
		v63 = v69
		v64 = v73
		goto L29
	} else {
		goto L35
	}
L35:
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	v25 = int32(base.Ui32(v7+int32(262120)) >> (uint(int32(2)) % 32))
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32(v25&int32(65535)) < base.Ui32(int32(291)) {
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
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40&int32(65535)<<(uint(int32(2))%32)+(l0+int32(24))-int32(3)))))
	if v51&int32(384) == int32(0) {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v57 = v40 + int32(1)
	v58 = int32(65535)
	if base.Ui32(v57&v58) <= base.Ui32(v25&v58) {
		v40 = v57
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v15 = l3<<(uint(int32(2))%32) + l2 + int32(20)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if base.Ui32(v16&int32(32767)+int32(base.Ui32(v16)>>(uint(int32(17))%32))) < base.Ui32(int32(8185)) {
		if base.Ui32(v16) < base.Ui32(int32(131072)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v86 + int32(4)
					F_errmsg(m, int32(730303), v9+int32(80))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
						*(*int32)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = int32(base.Ui32(v99)>>(uint(int32(15))%32)) & int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = int32(base.Ui32(v99) >> (uint(int32(17)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v99 & int32(32767)
						F_errdetail_internal(m, int32(606767), v9+int32(48))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(522395), int32(791), int32(315254))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
			if v16&int32(98304) != int32(32768) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v86 + int32(4)
						F_errmsg(m, int32(730303), v9+int32(80))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							*(*int32)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = int32(base.Ui32(v99)>>(uint(int32(15))%32)) & int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = int32(base.Ui32(v99) >> (uint(int32(17)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v99 & int32(32767)
							F_errdetail_internal(m, int32(606767), v9+int32(48))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(522395), int32(791), int32(315254))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
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
				m.G0 = v9 + int32(96)
				return v15
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v43 + int32(4)
				F_errmsg(m, int32(730409), v9+int32(32))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(base.Ui32(v54)>>(uint(int32(15))%32)) & int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(base.Ui32(v54) >> (uint(int32(17)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v54 & int32(32767)
					F_errdetail_internal(m, int32(606767), v9)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(522395), int32(775), int32(315254))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v5 = v3 << (uint(int32(8)) % 32)
	v6 = F_palloc(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v10 = F__emscripten_memcpy_bulkmem(m, v6, l0, v5)
			mBase = m.M
			v11 = v10
		} else {
			v11 = v6
		}
		return v11
	}
}
