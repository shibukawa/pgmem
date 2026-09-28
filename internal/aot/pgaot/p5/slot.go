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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = F_IsLogicalDecodingEnabled(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v17 = F_errstart(m, l0, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v99 = v2
					m.G0 = v7 + int32(48)
					return v99
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_0), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_ValidateSlotSyncParams_1), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v93 = v2
								v94 = int32(1234)
								F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_2), v94, int32(_a_F_ValidateSlotSyncParams_3))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v99 = v93
									m.G0 = v7 + int32(48)
									return v99
								}
							}
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[0]))
			if v34 != 0 {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
				if v35 != 0 {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[1])))
					if v51 == int32(0) {
						v55 = F_errstart(m, l0, int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 == int32(0) {
								v99 = v2
								m.G0 = v7 + int32(48)
								return v99
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_ValidateSlotSyncParams_4)
									F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_5), v7+int32(32))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v93 = v2
										v94 = int32(1265)
										F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_2), v94, int32(_a_F_ValidateSlotSyncParams_3))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v99 = v93
											m.G0 = v7 + int32(48)
											return v99
										}
									}
								}
							}
						}
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[2]))
						if v71 != 0 {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
							if v73 != 0 {
								v99 = int32(1)
								m.G0 = v7 + int32(48)
								return v99
							} else {
								v75 = int32(0)
								v77 = F_errstart(m, l0, v75)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									if v77 == int32(0) {
										v99 = v75
										m.G0 = v7 + int32(48)
										return v99
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_ValidateSlotSyncParams_6)
											F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_7), v7+int32(16))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												v93 = v75
												v94 = int32(1279)
												F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_2), v94, int32(_a_F_ValidateSlotSyncParams_3))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v99 = v93
													m.G0 = v7 + int32(48)
													return v99
												}
											}
										}
									}
								}
							}
						} else {
							v75 = int32(0)
							v77 = F_errstart(m, l0, v75)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								if v77 == int32(0) {
									v99 = v75
									m.G0 = v7 + int32(48)
									return v99
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_ValidateSlotSyncParams_6)
										F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_7), v7+int32(16))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											v93 = v75
											v94 = int32(1279)
											F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_2), v94, int32(_a_F_ValidateSlotSyncParams_3))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v99 = v93
												m.G0 = v7 + int32(48)
												return v99
											}
										}
									}
								}
							}
						}
					}
				} else {
					v37 = F_errstart(m, l0, int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v37 == int32(0) {
							v99 = v2
							m.G0 = v7 + int32(48)
							return v99
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_ValidateSlotSyncParams_8)
								F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_7), v7)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v93 = v2
									v94 = int32(1250)
									F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_2), v94, int32(_a_F_ValidateSlotSyncParams_3))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v99 = v93
										m.G0 = v7 + int32(48)
										return v99
									}
								}
							}
						}
					}
				}
			} else {
				v37 = F_errstart(m, l0, int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v37 == int32(0) {
						v99 = v2
						m.G0 = v7 + int32(48)
						return v99
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_ValidateSlotSyncParams_8)
							F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_7), v7)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v93 = v2
								v94 = int32(1250)
								F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_2), v94, int32(_a_F_ValidateSlotSyncParams_3))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v99 = v93
									m.G0 = v7 + int32(48)
									return v99
								}
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	m.T0[v19].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < v17 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L27
	}
L5:
	;
	v28 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v168 = v166 & int32(_a_F_slot_store_data_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v171)
	goto L26
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v42 = v35 + v36<<(uint(int32(3))%32) + v28*int32(100)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v153 = v28 + int32(1)
	if v153 != v17 {
		v28 = v153
		goto L8
	} else {
		goto L25
	}
L11:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v137+v28<<(uint(int32(3))%32)))) = int64(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v143+v28))) = uint8(v145)
	goto L10
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45+v28<<(uint(int32(1))%32)))))
	if v49 < int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v52 <= v49 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v55 = v42 + int32(28)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = v49
	v61 = v56 + v49<<(uint(int32(4))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v49))))
	v66 = v64 - int32(98)
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v123+v28<<(uint(int32(3))%32)))) = int64(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v129+v28))) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = int32(-1)
	goto L10
L16:
	;
	if v66 != int32(18) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	F_getTypeBinaryInputInfo(m, v96, v14+int32(28), v14+int32(24))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	F_getTypeInputInfo(m, v69, v14+int32(28), v14+int32(24))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v55)+76))
	v80 = F_OidInputFunctionCall(m, v76, v77, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v82+v28<<(uint(int32(3))%32)))) = v80
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v28))) = uint8(v89)
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = int32(-1)
	goto L10
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v55)+76))
	v106 = F_OidReceiveFunctionCall(m, v103, v61, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v108+v28<<(uint(int32(3))%32)))) = v106
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v113 != v114 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116+v28))) = uint8(v118)
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = int32(-1)
	goto L10
L25:
	;
	goto L9
L26:
	;
	m.G0 = v14 + int32(32)
	return
L27:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v49 + int32(1)
	F_errmsg_plural(m, int32(_a_F_slot_store_data_1), int32(_a_F_slot_store_data_2), v183, v14+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_slot_store_data_3), int32(1054), int32(_a_F_slot_store_data_4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v49 + int32(1)
	F_errmsg(m, int32(_a_F_slot_store_data_5), v14)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_slot_store_data_3), int32(1093), int32(_a_F_slot_store_data_4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
