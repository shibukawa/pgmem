package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wc_iscased_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v110
L2:
	;
	v51 = int32(691)
	v52 = int32(0)
	goto L16
L3:
	;
	v11 = int32(3408)
	v12 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_iscased_builtin[0]))))
	v110 = int32(base.Ui32(v41&int32(8)) >> (uint(int32(3)) % 32))
	goto L1
L6:
	;
	v17 = base.I32_div_s(v11+v12, int32(2))
	v19 = v17 * int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_wc_iscased_builtin[1])))
	if base.Ui32(v22) < base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_wc_iscased_builtin[2]))))
	if v35 != int32(3) {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v33 <= v32 {
		v11 = v32
		v12 = v33
		goto L6
	} else {
		goto L14
	}
L10:
	;
	v32 = v11
	v33 = v17 + int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_wc_iscased_builtin[3])))
	if base.Ui32(v28) <= base.Ui32(l0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v32 = v17 - int32(1)
	v33 = v12
	goto L9
L14:
	;
	goto L2
L15:
	;
	v110 = int32(1)
	goto L1
L16:
	;
	v57 = base.I32_div_s(v51+v52, int32(2))
	v59 = v57 << (uint(int32(3)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+uint32(_c_F_wc_iscased_builtin[4])))
	if base.Ui32(v62) < base.Ui32(l0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v79 = int32(659)
	v80 = int32(0)
	goto L26
L18:
	;
	if v74 <= v73 {
		v51 = v73
		v52 = v74
		goto L16
	} else {
		goto L25
	}
L19:
	;
	v73 = v51
	v74 = v57 + int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+uint32(_c_F_wc_iscased_builtin[5])))
	if base.Ui32(v68) <= base.Ui32(l0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v110 = int32(1)
	goto L1
L23:
	;
	goto L24
L24:
	;
	v73 = v57 - int32(1)
	v74 = v52
	goto L18
L25:
	;
	goto L17
L26:
	;
	v85 = base.I32_div_s(v79+v80, int32(2))
	v87 = v85 << (uint(int32(3)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_wc_iscased_builtin[6])))
	if base.Ui32(v90) < base.Ui32(l0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v110 = int32(0)
	goto L1
L28:
	;
	if v102 <= v101 {
		v79 = v101
		v80 = v102
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v101 = v79
	v102 = v85 + int32(1)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_wc_iscased_builtin[7])))
	if base.Ui32(v96) <= base.Ui32(l0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v110 = int32(1)
	goto L1
L33:
	;
	goto L34
L34:
	;
	v101 = v85 - int32(1)
	v102 = v80
	goto L28
L35:
	;
	goto L27
}
func F_wc_iscased_libc_other_mb(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	v4 = int32(1)
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v19 = v4
	} else {
		if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
			v19 = v4
		} else {
			v19 = base.B2i32(base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26))) != int32(0))
		}
	}
	return v19
}
func F_wc_islower_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	v4 = F_towupper(m, l0)
	return base.B2i32(base.B2i32(v4 != l0) != int32(0))
}
func F_wc_islower_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v13 int32
	_ = v13
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v13 = base.B2i32(base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26))) != int32(0))
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_wc_isprint_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v11 = base.B2i32(base.Ui32(l0-int32(32)) < base.Ui32(int32(95)))
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_wc_isspace_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v48
L2:
	;
	v10 = int32(10)
	v11 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_isspace_builtin[0]))))
	v48 = int32(base.Ui32(v38&int32(32)) >> (uint(int32(5)) % 32))
	goto L1
L5:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isspace_builtin[1])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v48 = int32(0)
	goto L1
L7:
	;
	if v33 <= v32 {
		v10 = v32
		v11 = v33
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v32 = v10
	v33 = v16 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isspace_builtin[2])))
	if base.Ui32(v28) <= base.Ui32(l0) {
		v48 = int32(1)
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v32 = v16 - int32(1)
	v33 = v11
	goto L7
L12:
	;
	goto L6
}
func F_wc_tolower_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	if base.Ui32(l0) <= base.Ui32(int32(127)) {
		v146 = l0<<(uint(int32(2))%32) + int32(_a_F_wc_tolower_builtin_0)
	} else {
		v9 = int32(0)
		if base.Ui32(l0) <= base.Ui32(int32(1415)) {
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[0]))))
			v141 = v14
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_1)) {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_2)) {
					if base.Ui32(l0-int32(_a_F_wc_tolower_builtin_3)) <= base.Ui32(int32(95)) {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[1]))))
						v141 = v27
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_4)) {
							v141 = v9
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_5)) {
								v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[2]))))
								v141 = v36
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_6)) {
									v141 = v9
								} else {
									v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[3]))))
									v141 = v43
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_7)) {
						v141 = v9
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_8)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_9)) {
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[4]))))
								v141 = v54
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_10)) {
									v141 = v9
								} else {
									v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[5]))))
									v141 = v61
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_11)) {
								v141 = v9
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_12)) {
									v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[6]))))
									v141 = v70
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_13)) {
										v141 = v9
									} else {
										v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[7]))))
										v141 = v77
									}
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_14)) {
					v141 = v9
				} else {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_15)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_16)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_17)) {
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[8]))))
								v141 = v90
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_18)) {
									v141 = v9
								} else {
									v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[9]))))
									v141 = v97
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_19)) {
								v141 = v9
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_20)) {
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[10]))))
									v141 = v106
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_21)) {
										v141 = v9
									} else {
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[11]))))
										v141 = v113
									}
								}
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_22)) {
							v141 = v9
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_23)) {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_tolower_builtin_24)) {
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[12]))))
									v141 = v124
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_tolower_builtin_25)) {
										v141 = v9
									} else {
										v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[13]))))
										v141 = v131
									}
								}
							} else {
								if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_wc_tolower_builtin_26)) {
									v141 = v9
								} else {
									v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_tolower_builtin[14]))))
									v141 = v140
								}
							}
						}
					}
				}
			}
		}
		v146 = v141<<(uint(int32(2))%32) + int32(_a_F_wc_tolower_builtin_27)
	}
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v147 != 0 {
		v148 = v147
	} else {
		v148 = l0
	}
	return v148
}
