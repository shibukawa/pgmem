package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_recordSharedDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v96 int32
	_ = v96
	var v125 int32
	_ = v125
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_recordSharedDependencyOn[0]))
	if v15 != 0 {
		v18 = F_table_open(m, int32(1214), int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.B2i32(base.B2i32(v20 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_recordSharedDependencyOn_0)) < base.Ui32(v21)) == int32(0))&((base.B2i32(v20 != int32(2615))|base.B2i32(v21 != int32(2200)))&base.B2i32(v20 != int32(1262))) == int32(0) {
				v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
				v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_shdepLockAndCheckObject(m, v44, v45)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+11)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v48
					v54 = int32(1)
					if v43 <= int32(3591) {
						if v43 <= int32(2670) {
							switch v43 - int32(1213) {
							case 0, 1, 19, 20, 47, 48, 49:
								v125 = v54
							case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
								v125 = int32(0)
							default:
								if base.Ui32(int32(2)) <= base.Ui32(v43-int32(2396)) {
									v125 = int32(0)
								} else {
									v125 = v54
								}
							}
						} else {
							v66 = v43 - int32(2671)
							if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v66))|base.B2i32(int32(1)<<(uint(v66)%32)&int32(226492515) == int32(0)) != 0 {
								if base.B2i32(base.Ui32(v43-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v43-int32(2846)) < base.Ui32(int32(2))) != 0 {
									v125 = v54
								} else {
									v125 = int32(0)
								}
							} else {
								v125 = v54
							}
						}
					} else {
						if v43 <= int32(_a_F_recordSharedDependencyOn_1) {
							v79 = v43 - int32(_a_F_recordSharedDependencyOn_2)
							if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v79))|base.B2i32(int32(1)<<(uint(v79)%32)&int32(963) == int32(0)) != 0 {
								if base.Ui32(v43-int32(3592)) < base.Ui32(int32(2)) {
									v125 = v54
								} else {
									if base.Ui32(int32(2)) <= base.Ui32(v43-int32(4060)) {
										v125 = int32(0)
									} else {
										v125 = v54
									}
								}
							} else {
								v125 = v54
							}
						} else {
							switch v43 - int32(_a_F_recordSharedDependencyOn_3) {
							case 0, 1, 2, 3, 4, 59, 60:
								v125 = v54
							case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
								v125 = int32(0)
							default:
								if base.Ui32(v43-int32(_a_F_recordSharedDependencyOn_4)) < base.Ui32(int32(3)) {
									v125 = v54
								} else {
									v96 = v43 - int32(_a_F_recordSharedDependencyOn_5)
									if base.Ui32(int32(15)) < base.Ui32(v96) {
										v125 = int32(0)
									} else {
										if int32(1)<<(uint(v96)%32)&int32(_a_F_recordSharedDependencyOn_6) != 0 {
											v125 = v54
										} else {
											v125 = int32(0)
										}
									}
								}
							}
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = base.I64_extend8_s(base.I64_extend_i32_u(l2))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = base.I64_extend_i32_u(v45)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = base.I64_extend_i32_u(v44)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = base.I64_extend_i32_u(v43)
					v139 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_recordSharedDependencyOn[1])))
					if v125 != 0 {
						v140 = int64(0)
					} else {
						v140 = v139
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v140
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
					v147 = F_heap_form_tuple(m, v142, v12+int32(16), v12+int32(8))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v18, v147)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							F_pfree(m, v147)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								F_relation_close(m, v18, int32(3))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							}
						}
					}
				}
			} else {
				F_relation_close(m, v18, int32(3))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					m.G0 = v12 + int32(80)
					return
				}
			}
		}
	} else {
		m.G0 = v12 + int32(80)
		return
	}
}
func F_shared_buffer_readv_stage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v67 int64
	_ = v67
	var v81 int64
	_ = v81
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v147 int64
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v160 int64
	_ = v160
	var v166 int64
	_ = v166
	var v170 int64
	_ = v170
	var v181 int64
	_ = v181
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(23)))) = uint8(v15)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	goto L1
L1:
	;
	v25 = v11 + int32(8)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = (l0 - v28) >> (uint(int32(7)) % 32)
	v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+4)) = uint32(v33)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+8)) = uint32(v35)
	goto L2
L2:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+23)))
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v44 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v11 + int32(48)
	return
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[1]))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(3))%32)+v44<<(uint(int32(3))%32))))
	v55 = v47 + base.I32_wrap_i64(v51)*int32(56)
	v57 = v55 - int32(32)
	v58 = int64(4194304)
	v60 = base.AtomicRmwOr64(m, v57, int32(0), v58)
	if v60&v58 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v67 = v60
	goto L11
L9:
	;
	v147 = v60
	goto L10
L10:
	;
	v154 = v55 - int32(20)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v154))) = v157
	v160 = v147 | int64(4194304)
	v166 = base.AtomicRmwCmpxchg64(m, v57, int32(0), v160, v147&int64(-4194305)+int64(1))
	if v166 != v160 {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(_a_F_shared_buffer_readv_stage_0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(_a_F_shared_buffer_readv_stage_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(_a_F_shared_buffer_readv_stage_2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(0)
	v81 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v81
	if v67&int64(4194304) != v81 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v147 = v140
	goto L10
L13:
	;
	goto L16
L14:
	;
	goto L15
L15:
	;
	v118 = int32(_a_F_shared_buffer_readv_stage_3)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[2]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(24))+8))
	if v121 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	F_perform_spin_delay(m, v11+int32(24))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	return
L19:
	;
	v99 = int64(0)
	v102 = base.AtomicRmwCmpxchg64(m, v57, int32(0), v99, v99)
	if v102&int64(4194304) != v99 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v138 = int64(4194304)
	v140 = base.AtomicRmwOr64(m, v57, int32(0), v138)
	if v140&v138 != int64(0) {
		v67 = v140
		goto L11
	} else {
		goto L32
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[2])) = v136
	goto L22
L24:
	;
	if int32(999) < v119 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v119 < int32(11) {
		goto L22
	} else {
		goto L31
	}
L27:
	;
	v126 = int32(900)
	if v126 <= v119 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v129 = v126
	goto L30
L29:
	;
	v129 = v119
	goto L30
L30:
	;
	v136 = v129 + int32(100)
	goto L23
L31:
	;
	v136 = v119 - int32(1)
	goto L23
L32:
	;
	goto L12
L33:
	;
	v170 = v166
	goto L36
L34:
	;
	goto L35
L35:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_stage[3]))
	F_ResourceOwnerForget(m, v192, base.I64_extend32_s(v51), int32(_a_F_shared_buffer_readv_stage_4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L18
	} else {
		goto L39
	}
L36:
	;
	v181 = base.AtomicRmwCmpxchg64(m, v57, int32(0), v170, v170&int64(-4194305)+int64(1))
	if v170 != v181 {
		v170 = v181
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	goto L37
L39:
	;
	v198 = v44 + int32(1)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+23)))
	if base.Ui32(v198) < base.Ui32(v199) {
		v44 = v198
		goto L6
	} else {
		goto L40
	}
L40:
	;
	goto L7
}
