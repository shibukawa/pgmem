package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SendSharedInvalidMessages(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	if l1 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	v18 = v14 + int32(12)
	v21 = l0
	v22 = l1
	goto L3
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v36 = F_LWLockAcquire(m, v32+int32(768), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v38 = int32(64)
	if base.Ui32(v38) <= base.Ui32(v22) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = v38
	goto L9
L8:
	;
	v41 = v22
	goto L9
L9:
	;
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v54 = v52 - v53
	if int32(4096) < v54+v41 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_SICleanupQueue(m, int32(1), v41)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L30
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v58 <= v54 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v60 = v21
	v63 = v52
	v64 = v41
	goto L15
L15:
	;
	v71 = base.I32_rem_s(v63, int32(4096))
	v74 = v14 + int32(16) + v71<<(uint(int32(4))%32)
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = v77
	v79 = int32(1)
	v80 = v63 + v79
	v82 = v60 + int32(16)
	if v79 < v64 {
		v60 = v82
		v63 = v80
		v64 = v64 - v79
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1)
	if v87 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_s_lock(m, v18, int32(491621), int32(422), int32(167501))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v95 = v22 - v41
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v80
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[791])))
	if v96 < v99 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[792])))
	v107 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v138+int32(768))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L28
	}
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v102+v107<<(uint(int32(2))%32))))
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(65570)+v117<<(uint(int32(4))%32)))) = uint8(v121)
	v124 = v107 + v121
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[791])))
	if v124 < v125 {
		v107 = v124
		goto L25
	} else {
		goto L27
	}
L26:
	;
	goto L24
L27:
	;
	goto L26
L28:
	;
	if int32(0) < v95 {
		v21 = v82
		v22 = v95
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	goto L10
}
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
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	if v14 != 0 {
		v17 = F_table_open(m, int32(1214), int32(3))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v21 = int32(0)
			if v19 == int32(2613) {
				v34 = v21
			} else {
				if base.Ui32(int32(11999)) < base.Ui32(v20) {
					v34 = v21
				} else {
					v34 = (base.B2i32(v19 != int32(2615)) | base.B2i32(v20 != int32(2200))) & base.B2i32(v19 != int32(1262))
				}
			}
			if v34 == int32(0) {
				v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_shdepLockAndCheckObject(m, v39, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+11)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v43
					v49 = int32(1)
					if v38 <= int32(3591) {
						if v38 <= int32(2670) {
							switch v38 - int32(1213) {
							case 0, 1, 19, 20, 47, 48, 49:
								v117 = v49
							case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
								v117 = int32(0)
							default:
								if base.Ui32(int32(2)) <= base.Ui32(v38-int32(2396)) {
									v117 = int32(0)
								} else {
									v117 = v49
								}
							}
						} else {
							v61 = v38 - int32(2671)
							if base.Ui32(int32(27)) < base.Ui32(v61) {
								if base.Ui32(v38-int32(2964)) < base.Ui32(int32(4)) {
									v117 = v49
								} else {
									if base.Ui32(v38-int32(2846)) < base.Ui32(int32(2)) {
										v117 = v49
									} else {
										v117 = int32(0)
									}
								}
							} else {
								if int32(1)<<(uint(v61)%32)&int32(226492515) == int32(0) {
									if base.Ui32(v38-int32(2964)) < base.Ui32(int32(4)) {
										v117 = v49
									} else {
										if base.Ui32(v38-int32(2846)) < base.Ui32(int32(2)) {
											v117 = v49
										} else {
											v117 = int32(0)
										}
									}
								} else {
									v117 = v49
								}
							}
						}
					} else {
						if v38 <= int32(5999) {
							v73 = v38 - int32(4177)
							if base.Ui32(int32(9)) < base.Ui32(v73) {
								if base.Ui32(v38-int32(3592)) < base.Ui32(int32(2)) {
									v117 = v49
								} else {
									if base.Ui32(int32(2)) <= base.Ui32(v38-int32(4060)) {
										v117 = int32(0)
									} else {
										v117 = v49
									}
								}
							} else {
								if int32(1)<<(uint(v73)%32)&int32(963) == int32(0) {
									if base.Ui32(v38-int32(3592)) < base.Ui32(int32(2)) {
										v117 = v49
									} else {
										if base.Ui32(int32(2)) <= base.Ui32(v38-int32(4060)) {
											v117 = int32(0)
										} else {
											v117 = v49
										}
									}
								} else {
									v117 = v49
								}
							}
						} else {
							switch v38 - int32(6243) {
							case 0, 1, 2, 3, 4, 59, 60:
								v117 = v49
							case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
								v117 = int32(0)
							default:
								if base.Ui32(v38-int32(6000)) < base.Ui32(int32(3)) {
									v117 = v49
								} else {
									v89 = v38 - int32(6100)
									if base.Ui32(int32(15)) < base.Ui32(v89) {
										v117 = int32(0)
									} else {
										if int32(1)<<(uint(v89)%32)&int32(49153) != 0 {
											v117 = v49
										} else {
											v117 = int32(0)
										}
									}
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = base.I32_extend8_s(l2)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v37
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v38
					v126 = *(*int32)(unsafe.Add(mBase, _consts[4]))
					if v117 != 0 {
						v127 = int32(0)
					} else {
						v127 = v126
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v127
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
					v134 = F_heap_form_tuple(m, v129, v11+int32(16), v11+int32(8))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v17, v134)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return
						} else {
							F_pfree(m, v134)
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return
							} else {
								F_sequence_close(m, v17, int32(3))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
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
				F_sequence_close(m, v17, int32(3))
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(23)))) = uint8(v14)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	goto L1
L1:
	;
	v24 = v10 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[742]))
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
	v40 = int32(0)
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19<<(uint(int32(3))%32)+v40<<(uint(int32(3))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(228784)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(493364)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	v62 = v49 + v47<<(uint(int32(6))%32)
	v64 = v62 - int32(40)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66
	if v65&v66 != 0 {
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
	v90 = v65
	goto L10
L10:
	;
	v98 = int32(4104476)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[736]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(24))+8))
	if v101 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_perform_spin_delay(m, v10+int32(24))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v90 = v82
	goto L10
L13:
	;
	return
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v83 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v82 | v83
	if v82&v83 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v119 = v62 - int32(28)
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = (v90 + int32(4194305)) & int32(-4194305)
	v130 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	F_ResourceOwnerForget(m, v130, v47, int32(1611036))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L27
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[736])) = v116
	goto L17
L19:
	;
	if int32(999) < v99 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v99 < int32(11) {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v106 = int32(900)
	if v106 <= v99 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v109 = v106
	goto L25
L24:
	;
	v109 = v99
	goto L25
L25:
	;
	v116 = v109 + int32(100)
	goto L18
L26:
	;
	v116 = v99 - int32(1)
	goto L18
L27:
	;
	v135 = v40 + int32(1)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+23)))
	if base.Ui32(v135) < base.Ui32(v136) {
		v40 = v135
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L7
}
