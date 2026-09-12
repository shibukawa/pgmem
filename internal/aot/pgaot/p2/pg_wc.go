package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_islower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	v3 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v3 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L2
	case 2:
		goto L1
	default:
		goto L4
	}
L1:
	;
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v59 = F_towupper(m, l0)
	mBase = m.M
	goto L19
L3:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
L5:
	;
	return v54
L6:
	;
	v19 = int32(0)
	v20 = int32(689)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v54 = base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
	goto L5
L9:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 << (uint(int32(3)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[600])))
	if base.Ui32(v30) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v54 = int32(0)
	goto L5
L11:
	;
	if v41 <= v42 {
		v19 = v41
		v20 = v42
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v41 = v25 + int32(1)
	v42 = v20
	goto L11
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[601])))
	if base.Ui32(v36) <= base.Ui32(l0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = int32(1)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v41 = v19
	v42 = v25 - int32(1)
	goto L11
L18:
	;
	goto L10
L19:
	;
	return base.B2i32(v59 != l0)
L20:
	;
	goto L23
L21:
	;
	v74 = int32(0)
	goto L22
L22:
	;
	return v74
L23:
	;
	v74 = base.B2i32(base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L22
}
func F_pg_wc_toupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	v4 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	switch v4 - int32(1) {
	case 0:
		if base.Ui32(l0) <= base.Ui32(int32(127)) {
			v126 = l0<<(uint(int32(2))%32) + int32(1859716)
		} else {
			v27 = int32(0)
			if base.Ui32(l0) < base.Ui32(int32(1416)) {
				v114 = l0
				v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
				v121 = v119
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(43967)) {
					if base.Ui32(l0) <= base.Ui32(int32(8580)) {
						if base.Ui32(l0-int32(4256)) <= base.Ui32(int32(95)) {
							v114 = l0 - int32(2840)
							v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
							v121 = v119
						} else {
							if base.Ui32(l0) < base.Ui32(int32(5024)) {
								v121 = v27
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(5117)) {
									v114 = l0 - int32(3512)
									v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
									v121 = v119
								} else {
									if base.Ui32(l0) < base.Ui32(int32(7296)) {
										v121 = v27
									} else {
										v114 = l0 - int32(5690)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									}
								}
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(9398)) {
							v121 = v27
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(11565)) {
								if base.Ui32(l0) <= base.Ui32(int32(9449)) {
									v114 = l0 - int32(6507)
									v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
									v121 = v119
								} else {
									if base.Ui32(l0) < base.Ui32(int32(11264)) {
										v121 = v27
									} else {
										v114 = l0 - int32(8321)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(42560)) {
									v121 = v27
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(42998)) {
										v114 = l0 - int32(39315)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									} else {
										if base.Ui32(l0) < base.Ui32(int32(43859)) {
											v121 = v27
										} else {
											v114 = l0 - int32(40175)
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
											v121 = v119
										}
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(64256)) {
						v121 = v27
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(68997)) {
							if base.Ui32(l0) <= base.Ui32(int32(65370)) {
								if base.Ui32(l0) <= base.Ui32(int32(64279)) {
									v114 = l0 - int32(60463)
									v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
									v121 = v119
								} else {
									if base.Ui32(l0) < base.Ui32(int32(65313)) {
										v121 = v27
									} else {
										v114 = l0 - int32(61496)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(66560)) {
									v121 = v27
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(67004)) {
										v114 = l0 - int32(62685)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									} else {
										if base.Ui32(l0) < base.Ui32(int32(68736)) {
											v121 = v27
										} else {
											v114 = l0 - int32(64416)
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
											v121 = v119
										}
									}
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(71840)) {
								v121 = v27
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(93823)) {
									if base.Ui32(l0) <= base.Ui32(int32(71903)) {
										v114 = l0 - int32(67258)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									} else {
										if base.Ui32(l0) < base.Ui32(int32(93760)) {
											v121 = v27
										} else {
											v114 = l0 - int32(89114)
											v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
											v121 = v119
										}
									}
								} else {
									if base.Ui32(int32(67)) < base.Ui32(l0-int32(125184)) {
										v121 = v27
									} else {
										v114 = l0 - int32(120474)
										v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114<<(uint(int32(1))%32))+uint32(_consts[604]))))
										v121 = v119
									}
								}
							}
						}
					}
				}
			}
			v126 = v121<<(uint(int32(2))%32) + int32(1859712)
		}
		v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
		if v127 != 0 {
			v128 = v127
		} else {
			v128 = l0
		}
		return v128
	case 1:
		v131 = *(*int32)(unsafe.Add(mBase, _consts[587]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v153 = F_casemap(m, l0, int32(1))
			mBase = m.M
			return v153
		} else {
			v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+4)))
			if v134&int32(1) == int32(0) {
				v153 = F_casemap(m, l0, int32(1))
				mBase = m.M
				return v153
			} else {
				if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
					v147 = l0 - int32(32)
				} else {
					v147 = l0
				}
				return v147 & int32(255)
			}
		}
	case 2:
		v156 = *(*int32)(unsafe.Add(mBase, _consts[587]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			if base.Ui32(int32(255)) < base.Ui32(l0) {
				v186 = l0
			} else {
				if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
					v185 = l0 & int32(95)
				} else {
					v185 = l0
				}
				v186 = v185
			}
			return v186
		} else {
			v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)))
			if v159&int32(1) == int32(0) {
				if base.Ui32(int32(255)) < base.Ui32(l0) {
					v186 = l0
				} else {
					if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
						v185 = l0 & int32(95)
					} else {
						v185 = l0
					}
					v186 = v185
				}
				return v186
			} else {
				if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
					v172 = l0 - int32(32)
				} else {
					v172 = l0
				}
				return v172 & int32(255)
			}
		}
	default:
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v186 = l0
			return v186
		} else {
			if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
				v17 = l0 - int32(32)
			} else {
				v17 = l0
			}
			return v17 & int32(255)
		}
	}
}
