package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_recordSharedDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v94 int32
	_ = v94
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_recordSharedDependencyOn[0]))
	if v14 != 0 {
		v17 = F_table_open(m, int32(1214), int32(3))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.B2i32(base.B2i32(v19 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_recordSharedDependencyOn_0)) < base.Ui32(v20)) == int32(0))&((base.B2i32(v19 != int32(2615))|base.B2i32(v20 != int32(2200)))&base.B2i32(v19 != int32(1262))) == int32(0) {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_shdepLockAndCheckObject(m, v42, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v46 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+11)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
					v52 = int32(1)
					if v41 <= int32(3591) {
						if v41 <= int32(2670) {
							switch v41 - int32(1213) {
							case 0, 1, 19, 20, 47, 48, 49:
								v123 = v52
							case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
								v123 = int32(0)
							default:
								if base.Ui32(int32(2)) <= base.Ui32(v41-int32(2396)) {
									v123 = int32(0)
								} else {
									v123 = v52
								}
							}
						} else {
							v64 = v41 - int32(2671)
							if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v64))|base.B2i32(int32(1)<<(uint(v64)%32)&int32(226492515) == int32(0)) != 0 {
								if base.B2i32(base.Ui32(v41-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v41-int32(2846)) < base.Ui32(int32(2))) != 0 {
									v123 = v52
								} else {
									v123 = int32(0)
								}
							} else {
								v123 = v52
							}
						}
					} else {
						if v41 <= int32(_a_F_recordSharedDependencyOn_1) {
							v77 = v41 - int32(_a_F_recordSharedDependencyOn_2)
							if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v77))|base.B2i32(int32(1)<<(uint(v77)%32)&int32(963) == int32(0)) != 0 {
								if base.Ui32(v41-int32(3592)) < base.Ui32(int32(2)) {
									v123 = v52
								} else {
									if base.Ui32(int32(2)) <= base.Ui32(v41-int32(4060)) {
										v123 = int32(0)
									} else {
										v123 = v52
									}
								}
							} else {
								v123 = v52
							}
						} else {
							switch v41 - int32(_a_F_recordSharedDependencyOn_3) {
							case 0, 1, 2, 3, 4, 59, 60:
								v123 = v52
							case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
								v123 = int32(0)
							default:
								if base.Ui32(v41-int32(_a_F_recordSharedDependencyOn_4)) < base.Ui32(int32(3)) {
									v123 = v52
								} else {
									v94 = v41 - int32(_a_F_recordSharedDependencyOn_5)
									if base.Ui32(int32(15)) < base.Ui32(v94) {
										v123 = int32(0)
									} else {
										if int32(1)<<(uint(v94)%32)&int32(_a_F_recordSharedDependencyOn_6) != 0 {
											v123 = v52
										} else {
											v123 = int32(0)
										}
									}
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = base.I32_extend8_s(l2)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v41
					v132 = *(*int32)(unsafe.Add(mBase, _c_F_recordSharedDependencyOn[1]))
					if v123 != 0 {
						v133 = int32(0)
					} else {
						v133 = v132
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v133
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
					v140 = F_heap_form_tuple(m, v135, v11+int32(16), v11+int32(8))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v17, v140)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							F_pfree(m, v140)
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								F_relation_close(m, v17, int32(3))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									m.G0 = v11 + int32(48)
									return
								}
							}
						}
					}
				}
			} else {
				F_relation_close(m, v17, int32(3))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return
				} else {
					m.G0 = v11 + int32(48)
					return
				}
			}
		}
	} else {
		m.G0 = v11 + int32(48)
		return
	}
}
func F_shared_buffer_readv_stage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(23)))) = uint8(v14)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	goto L1
L1:
	;
	v24 = v10 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = (l0 - v27) >> (uint(int32(7)) % 32)
	v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+4)) = uint32(v32)
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+8)) = uint32(v34)
	goto L2
L2:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+23)))
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v41 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v10 + int32(48)
	return
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19<<(uint(int32(3))%32)+v41<<(uint(int32(3))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(_a_F_shared_buffer_readv_stage_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(_a_F_shared_buffer_readv_stage_1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(_a_F_shared_buffer_readv_stage_2)
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	v62 = v49 + v47<<(uint(int32(6))%32)
	v64 = v62 - int32(40)
	v65 = int32(_a_F_shared_buffer_readv_stage_3)
	v67 = base.AtomicRmwOr32(m, v64, v56, v65)
	if v67&v65 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	goto L11
L9:
	;
	v91 = v67
	goto L10
L10:
	;
	v96 = int32(_a_F_shared_buffer_readv_stage_4)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[2]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(24))+8))
	if v99 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_perform_spin_delay(m, v10+int32(24))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v91 = v83
	goto L10
L13:
	;
	return
L14:
	;
	v81 = int32(_a_F_shared_buffer_readv_stage_3)
	v83 = base.AtomicRmwOr32(m, v64, int32(0), v81)
	if v83&v81 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v117 = v62 - int32(28)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = (v91 + int32(_a_F_shared_buffer_readv_stage_5)) & int32(-4194305)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[3]))
	F_ResourceOwnerForget(m, v128, v47, int32(_a_F_shared_buffer_readv_stage_6))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L27
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[2])) = v114
	goto L17
L19:
	;
	if int32(999) < v97 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v97 < int32(11) {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v104 = int32(900)
	if v104 <= v97 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v107 = v104
	goto L25
L24:
	;
	v107 = v97
	goto L25
L25:
	;
	v114 = v107 + int32(100)
	goto L18
L26:
	;
	v114 = v97 - int32(1)
	goto L18
L27:
	;
	v133 = v41 + int32(1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+23)))
	if base.Ui32(v133) < base.Ui32(v134) {
		v41 = v133
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L7
}
