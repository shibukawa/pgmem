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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_islower[0]))
	switch v3 - int32(1) {
	case 0:
		v15 = Fn13990(m, l0, int32(97), int32(_a_F_pg_wc_islower_0), int32(_a_F_pg_wc_islower_1), int32(689))
		mBase = m.M
		return v15
	case 1:
		v20 = F_towupper(m, l0)
		mBase = m.M
		return base.B2i32(v20 != l0)
	case 2:
		if base.Ui32(l0) <= base.Ui32(int32(255)) {
			v35 = base.B2i32(base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26))) != int32(0))
		} else {
			v35 = int32(0)
		}
		return v35
	default:
		return base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
	}
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
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_toupper[0]))
	switch v4 - int32(1) {
	case 0:
		if base.Ui32(l0) <= base.Ui32(int32(127)) {
			v164 = l0<<(uint(int32(2))%32) + int32(_a_F_pg_wc_toupper_0)
		} else {
			v27 = int32(0)
			if base.Ui32(l0) <= base.Ui32(int32(1415)) {
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[1]))))
				v159 = v32
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_1)) {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_2)) {
						if base.Ui32(l0-int32(_a_F_pg_wc_toupper_3)) <= base.Ui32(int32(95)) {
							v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[2]))))
							v159 = v45
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_4)) {
								v159 = v27
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_5)) {
									v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[3]))))
									v159 = v54
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_6)) {
										v159 = v27
									} else {
										v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[4]))))
										v159 = v61
									}
								}
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_7)) {
							v159 = v27
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_8)) {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_9)) {
									v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[5]))))
									v159 = v72
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_10)) {
										v159 = v27
									} else {
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[6]))))
										v159 = v79
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_11)) {
									v159 = v27
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_12)) {
										v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[7]))))
										v159 = v88
									} else {
										if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_13)) {
											v159 = v27
										} else {
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[8]))))
											v159 = v95
										}
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_14)) {
						v159 = v27
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_15)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_16)) {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_17)) {
									v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[9]))))
									v159 = v108
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_18)) {
										v159 = v27
									} else {
										v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[10]))))
										v159 = v115
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_19)) {
									v159 = v27
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_20)) {
										v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[11]))))
										v159 = v124
									} else {
										if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_21)) {
											v159 = v27
										} else {
											v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[12]))))
											v159 = v131
										}
									}
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_22)) {
								v159 = v27
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_23)) {
									if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_toupper_24)) {
										v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[13]))))
										v159 = v142
									} else {
										if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_toupper_25)) {
											v159 = v27
										} else {
											v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[14]))))
											v159 = v149
										}
									}
								} else {
									if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_pg_wc_toupper_26)) {
										v159 = v27
									} else {
										v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_toupper[15]))))
										v159 = v158
									}
								}
							}
						}
					}
				}
			}
			v164 = v159<<(uint(int32(2))%32) + int32(_a_F_pg_wc_toupper_27)
		}
		v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
		if v165 != 0 {
			v166 = v165
		} else {
			v166 = l0
		}
		return v166
	case 1:
		v169 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_toupper[16]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v191 = F_casemap(m, l0, int32(1))
			mBase = m.M
			return v191
		} else {
			v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+4)))
			if v172&int32(1) == int32(0) {
				v191 = F_casemap(m, l0, int32(1))
				mBase = m.M
				return v191
			} else {
				if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
					v185 = l0 - int32(32)
				} else {
					v185 = l0
				}
				return v185 & int32(255)
			}
		}
	case 2:
		v194 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_toupper[16]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			if base.Ui32(int32(255)) < base.Ui32(l0) {
				v224 = l0
			} else {
				if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
					v223 = l0 & int32(95)
				} else {
					v223 = l0
				}
				v224 = v223
			}
			return v224
		} else {
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+4)))
			if v197&int32(1) == int32(0) {
				if base.Ui32(int32(255)) < base.Ui32(l0) {
					v224 = l0
				} else {
					if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
						v223 = l0 & int32(95)
					} else {
						v223 = l0
					}
					v224 = v223
				}
				return v224
			} else {
				if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
					v210 = l0 - int32(32)
				} else {
					v210 = l0
				}
				return v210 & int32(255)
			}
		}
	default:
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v224 = l0
			return v224
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
