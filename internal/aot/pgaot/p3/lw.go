package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockQueueSelf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[0]))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L45
	}
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+344)))
	if v12 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L16
	} else {
		goto L42
	}
L5:
	;
	v13 = int32(536870912)
	v15 = base.AtomicRmwOr32(m, l0, int32(4), v13)
	if v15&v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v20 = v15
	goto L9
L7:
	;
	goto L8
L8:
	;
	v87 = base.AtomicRmwOr32(m, l0, int32(4), int32(-2147483648))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+345)) = uint8(v2)
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+344)) = uint8(v91)
	if v2 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_LWLockQueueSelf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(860)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_LWLockQueueSelf_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	if v20&int32(536870912) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	goto L14
L12:
	;
	goto L13
L13:
	;
	v55 = int32(_a_F_LWLockQueueSelf_2)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[1]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v58 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v44&int32(536870912) != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v75 = int32(536870912)
	v77 = base.AtomicRmwOr32(m, l0, int32(4), v75)
	if v77&v75 != 0 {
		v20 = v77
		goto L9
	} else {
		goto L30
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[1])) = v73
	goto L20
L22:
	;
	if int32(999) < v56 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v56 < int32(11) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v63 = int32(900)
	if v63 <= v56 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v66 = v63
	goto L28
L27:
	;
	v66 = v56
	goto L28
L28:
	;
	v73 = v66 + int32(100)
	goto L21
L29:
	;
	v73 = v56 - int32(1)
	goto L21
L30:
	;
	goto L10
L31:
	;
	v151 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	m.G0 = v8 + int32(32)
	return
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[3]))
	v102 = v97 + v99*int32(768)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v103 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[3]))
	v128 = v123 + v125*int32(768)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v129 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v102)+348)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v99
	goto L31
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+348)) = v103
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v113+v103*int32(768))+352)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v102)+352)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v99
	goto L31
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v125
	goto L31
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128)+348)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v125
	goto L38
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+352)) = v129
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockQueueSelf[2]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v129*int32(768))+348)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v128)+348)) = int32(-1)
	goto L38
L42:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockQueueSelf_3), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_LWLockQueueSelf_1), int32(1026), int32(_a_F_LWLockQueueSelf_4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockQueueSelf_5), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_LWLockQueueSelf_1), int32(1029), int32(_a_F_LWLockQueueSelf_4))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LWLockShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v134 int32
	_ = v134
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(_a_F_LWLockShmemRequest_0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+20)) = int64(17416)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_LWLockShmemRequest_1)
	F_ShmemRequestStructWithOpts(m, v12+int32(16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LWLockShmemRequest[0])))
		if v26 != 0 {
			v117 = int32(-1)
		} else {
			v27 = int32(_a_F_LWLockShmemRequest_2)
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemRequest[1]))
			if v29 == int32(0) {
				v117 = v27
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				if v32 <= int32(0) {
					v117 = v27
				} else {
					v36 = v32 & int32(3)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
					v38 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v32) {
						v44 = v38
						v46 = v38
						v51 = v2
						for {
							v55 = v37 + v44<<(uint(int32(2))%32)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+64))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+64))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+64))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+64))
							v67 = v57 + (v59 + (v61 + (v63 + v46)))
							v68 = int32(4)
							v69 = v44 + v68
							v71 = v51 + v68
							if v71 != v32&int32(2147483644) {
								v44 = v69
								v46 = v67
								v51 = v71
								continue
							} else {
								break
							}
							break
						}
						if v36 == int32(0) {
							v106 = v67
						} else {
							v75 = v69
							v77 = v67
							v84 = v75
							v86 = v77
							v92 = v2
							for {
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v37+v84<<(uint(int32(2))%32))))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
								v98 = v97 + v86
								v99 = int32(1)
								v102 = v92 + v99
								if v102 != v36 {
									v84 = v84 + v99
									v86 = v98
									v92 = v102
									continue
								} else {
									break
								}
								break
							}
							v106 = v98
						}
					} else {
						v75 = v38
						v77 = v38
						v84 = v75
						v86 = v77
						v92 = v2
						for {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v37+v84<<(uint(int32(2))%32))))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
							v98 = v97 + v86
							v99 = int32(1)
							v102 = v92 + v99
							if v102 != v36 {
								v84 = v84 + v99
								v86 = v98
								v92 = v102
								continue
							} else {
								break
							}
							break
						}
						v106 = v98
					}
					v117 = v106<<(uint(int32(7))%32) + int32(_a_F_LWLockShmemRequest_2)
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(_a_F_LWLockShmemRequest_3)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v117
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_LWLockShmemRequest_4)
		F_ShmemRequestStructWithOpts(m, v12)
		mBase = m.M
		v134 = m.ExcPending
		if v134 != 0 {
			return
		} else {
			m.G0 = v12 + int32(32)
			return
		}
	}
}
