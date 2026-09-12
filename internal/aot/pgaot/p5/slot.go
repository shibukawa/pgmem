package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateSlotSyncParams(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v11 <= int32(1) {
		v15 = F_errstart(m, l0, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v96 = v2
				m.G0 = v7 + int32(48)
				return v96
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(703027), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v90 = v2
						v92 = int32(1070)
						F_errfinish(m, int32(491349), v92, int32(148545))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v90
							m.G0 = v7 + int32(48)
							return v96
						}
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _consts[522]))
		if v30 != 0 {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
			if v31 != 0 {
				v47 = int32(0)
				v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[523])))
				if v49 == v47 {
					v53 = F_errstart(m, l0, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						if v53 == int32(0) {
							v96 = v47
							m.G0 = v7 + int32(48)
							return v96
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(314128)
								F_errmsg(m, int32(448302), v7+int32(32))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v90 = v47
									v92 = int32(1100)
									F_errfinish(m, int32(491349), v92, int32(148545))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v96 = v90
										m.G0 = v7 + int32(48)
										return v96
									}
								}
							}
						}
					}
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, _consts[524]))
					if v69 != 0 {
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
						if v71 != 0 {
							v96 = int32(1)
							m.G0 = v7 + int32(48)
							return v96
						} else {
							v73 = int32(0)
							v75 = F_errstart(m, l0, v73)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								if v75 == int32(0) {
									v96 = v73
									m.G0 = v7 + int32(48)
									return v96
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(237423)
										F_errmsg(m, int32(104893), v7+int32(16))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											v90 = v73
											v92 = int32(1114)
											F_errfinish(m, int32(491349), v92, int32(148545))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												v96 = v90
												m.G0 = v7 + int32(48)
												return v96
											}
										}
									}
								}
							}
						}
					} else {
						v73 = int32(0)
						v75 = F_errstart(m, l0, v73)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							if v75 == int32(0) {
								v96 = v73
								m.G0 = v7 + int32(48)
								return v96
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(237423)
									F_errmsg(m, int32(104893), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										v90 = v73
										v92 = int32(1114)
										F_errfinish(m, int32(491349), v92, int32(148545))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											v96 = v90
											m.G0 = v7 + int32(48)
											return v96
										}
									}
								}
							}
						}
					}
				}
			} else {
				v32 = int32(0)
				v34 = F_errstart(m, l0, v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 == int32(0) {
						v96 = v32
						m.G0 = v7 + int32(48)
						return v96
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(371995)
							F_errmsg(m, int32(104893), v7)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v90 = v32
								v92 = int32(1085)
								F_errfinish(m, int32(491349), v92, int32(148545))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									v96 = v90
									m.G0 = v7 + int32(48)
									return v96
								}
							}
						}
					}
				}
			}
		} else {
			v32 = int32(0)
			v34 = F_errstart(m, l0, v32)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				if v34 == int32(0) {
					v96 = v32
					m.G0 = v7 + int32(48)
					return v96
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(371995)
						F_errmsg(m, int32(104893), v7)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v90 = v32
							v92 = int32(1085)
							F_errfinish(m, int32(491349), v92, int32(148545))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v96 = v90
								m.G0 = v7 + int32(48)
								return v96
							}
						}
					}
				}
			}
		}
	}
}
func F_slot_store_data(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	m.T0[v17].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L23
	}
L4:
	;
	v25 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v156 = v154 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v156)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v159)
	goto L22
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v38 = v31 + v32<<(uint(int32(4))%32) + v25*int32(100)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+111)))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v143 = v25 + int32(1)
	if v143 != v15 {
		v25 = v143
		goto L7
	} else {
		goto L21
	}
L10:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v25<<(uint(int32(2))%32)))) = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v25))) = uint8(v137)
	goto L9
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v25<<(uint(int32(1))%32)))))
	if v45 < int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = v38 + int32(20)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, _consts[535])) = v45
	v55 = v50 + v45<<(uint(int32(4))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v45))))
	switch v58 - int32(98) {
	case 0:
		goto L14
	default:
		goto L13
	case 18:
		goto L15
	}
L13:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v115+v25<<(uint(int32(2))%32)))) = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v121+v25))) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	goto L9
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	F_getTypeBinaryInputInfo(m, v88, v12+int32(12), v12+int32(8))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	F_getTypeInputInfo(m, v61, v12+int32(12), v12+int32(8))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v72 = F_OidInputFunctionCall(m, v68, v69, v70, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v25<<(uint(int32(2))%32)))) = v72
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79+v25))) = uint8(v81)
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	goto L9
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v98 = F_OidReceiveFunctionCall(m, v95, v55, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v100+v25<<(uint(int32(2))%32)))) = v98
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v105 != v106 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+v25))) = uint8(v110)
	*(*int32)(unsafe.Add(mBase, _consts[535])) = int32(-1)
	goto L9
L21:
	;
	goto L8
L22:
	;
	m.G0 = v12 + int32(16)
	return
L23:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v45 + int32(1)
	F_errmsg(m, int32(465206), v12)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(486965), int32(847), int32(495475))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
