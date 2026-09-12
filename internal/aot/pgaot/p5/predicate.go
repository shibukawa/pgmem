package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PredicateLockPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[789]))
	if v11 == int32(0) {
		m.G0 = v8 + int32(16)
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		switch v14 {
		case 0, 5:
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+108)))
			if v15&int32(128) != 0 {
				F_ReleasePredicateLocks(m, int32(0), int32(1))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v22) < base.Ui32(int32(12000)) {
					m.G0 = v8 + int32(16)
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+118)))
					if v26 == int32(116) {
						m.G0 = v8 + int32(16)
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
						F_PredicateLockAcquire(m, v8)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		default:
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_PredicateLockPageSplit(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == int32(0) {
		m.G0 = v9 + int32(96)
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui32(v16) < base.Ui32(int32(12000)) {
			m.G0 = v9 + int32(96)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
			if v20 == int32(116) {
				m.G0 = v9 + int32(96)
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v25 = v9 + int32(88)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = l1
				v28 = v9 + int32(72)
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = l2
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v9)+76)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v23
				v39 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				v43 = F_LWLockAcquire(m, v39+int32(3840), v30)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v45
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v47
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v9)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v49
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v9)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v51
					v58 = F_TransferPredicateLocksToNewTarget(m, v9+int32(48), v9+int32(32), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						if v58 == int32(0) {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
							if v62 == int32(0) {
								v65 = int32(-1)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
								if v66 == v65 {
								} else {
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v9)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v69
									v74 = v65
									*(*int32)(unsafe.Add(mBase, uint32(v9)+76)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v74
								}
							} else {
								v71 = *(*int64)(unsafe.Add(mBase, uint32(v9)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v71
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
								v74 = v73
								*(*int32)(unsafe.Add(mBase, uint32(v9)+76)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v74
							}
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+88))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v79
							v81 = *(*int64)(unsafe.Add(mBase, uint32(v9)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v81
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v9)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v83
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v9)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v85
							v90 = F_TransferPredicateLocksToNewTarget(m, v9+int32(16), v9, int32(1))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, _consts[7]))
								F_LWLockRelease(m, v94+int32(3840))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									m.G0 = v9 + int32(96)
									return
								}
							}
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							F_LWLockRelease(m, v94+int32(3840))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								m.G0 = v9 + int32(96)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_PredicateLockTID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[789]))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v15 {
	case 0, 5:
		goto L3
	default:
		goto L1
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
	if v16&int32(128) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v23) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	goto L1
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
	if v27 == int32(116) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.Ui32(l3) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v154 = v23
	goto L13
L13:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v155
	v161 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v162 = int32(0)
	v164 = F_hash_search(m, v161, v9, v162, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L55
	}
L14:
	;
	if v152 != 0 {
		goto L1
	} else {
		goto L54
	}
L15:
	;
	v152 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v43 == l3 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v152 = int32(1)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v47 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v152 = v144
	goto L14
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v51 == int32(0) {
		v144 = int32(0)
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v115 = int32(0)
	v117 = v47 - int32(1)
	goto L44
L25:
	;
	v56 = v51
	goto L26
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v61 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v144 = int32(0)
	goto L21
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v56)+80))
	if v108 != 0 {
		v56 = v108
		goto L26
	} else {
		goto L43
	}
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v64 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v67 = int32(1)
	if l3 == v64 {
		v144 = v67
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
	v71 = v69 - int32(1)
	if v71 < int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v76 = int32(0)
	v78 = v71
	goto L33
L33:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v84 = int32(2)
	v85 = base.I32_div_s(v78-v76, v84)
	v86 = v85 + v76
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82+v86<<(uint(v84)%32))))
	if v90 == l3 {
		v144 = v67
		goto L21
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v94 = F_TransactionIdPrecedes(m, v90, l3)
	mBase = m.M
	if v94 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v95 = v86 + int32(1)
	goto L38
L37:
	;
	v95 = v76
	goto L38
L38:
	;
	if v94 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v98 = v78
	goto L41
L40:
	;
	v98 = v86 - int32(1)
	goto L41
L41:
	;
	if v95 <= v98 {
		v76 = v95
		v78 = v98
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	goto L27
L44:
	;
	v122 = int32(2)
	v123 = base.I32_div_s(v117-v115, v122)
	v124 = v123 + v115
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v113+v124<<(uint(v122)%32))))
	v129 = base.B2i32(v128 == l3)
	if v128 == l3 {
		v144 = v129
		goto L21
	} else {
		goto L46
	}
L45:
	;
	v144 = v129
	goto L21
L46:
	;
	v132 = base.B2i32(base.Ui32(v128) < base.Ui32(l3))
	if base.Ui32(v128) < base.Ui32(l3) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v133 = v124 + int32(1)
	goto L49
L48:
	;
	v133 = v115
	goto L49
L49:
	;
	if base.Ui32(v128) < base.Ui32(l3) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v136 = v117
	goto L52
L51:
	;
	v136 = v124 - int32(1)
	goto L52
L52:
	;
	if v133 <= v136 {
		v115 = v133
		v117 = v136
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v154 = v153
	goto L13
L55:
	;
	if v164 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+16)))
	if v166 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v169
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v171 | v172<<(uint(int32(16))%32)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v177
	F_PredicateLockAcquire(m, v9)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L1
}
