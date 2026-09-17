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
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[0]))
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
					F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_0), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v90 = v2
						v92 = int32(1070)
						F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_1), v92, int32(_a_F_ValidateSlotSyncParams_2))
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
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[1]))
		if v30 != 0 {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
			if v31 != 0 {
				v47 = int32(0)
				v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[2])))
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
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_ValidateSlotSyncParams_3)
								F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_4), v7+int32(32))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v90 = v47
									v92 = int32(1100)
									F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_1), v92, int32(_a_F_ValidateSlotSyncParams_2))
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
					v69 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateSlotSyncParams[3]))
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
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_ValidateSlotSyncParams_5)
										F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_6), v7+int32(16))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											v90 = v73
											v92 = int32(1114)
											F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_1), v92, int32(_a_F_ValidateSlotSyncParams_2))
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
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_ValidateSlotSyncParams_5)
									F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_6), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										v90 = v73
										v92 = int32(1114)
										F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_1), v92, int32(_a_F_ValidateSlotSyncParams_2))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_ValidateSlotSyncParams_7)
							F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_6), v7)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v90 = v32
								v92 = int32(1085)
								F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_1), v92, int32(_a_F_ValidateSlotSyncParams_2))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_ValidateSlotSyncParams_7)
						F_errmsg(m, int32(_a_F_ValidateSlotSyncParams_6), v7)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v90 = v32
							v92 = int32(1085)
							F_errfinish(m, int32(_a_F_ValidateSlotSyncParams_1), v92, int32(_a_F_ValidateSlotSyncParams_2))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	m.T0[v18].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L25
	}
L4:
	;
	v28 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v162 = v160 & int32(_a_F_slot_store_data_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v162)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v165)
	goto L24
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v40 = v33 + v34<<(uint(int32(4))%32) + v28*int32(100)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+111)))
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v148 = v28 + int32(1)
	if v148 != v16 {
		v28 = v148
		goto L7
	} else {
		goto L23
	}
L10:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v133+v28<<(uint(int32(2))%32)))) = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v139+v28))) = uint8(v141)
	goto L9
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43+v28<<(uint(int32(1))%32)))))
	if v47 < int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v51 = v40 + int32(20)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = v47
	v57 = v52 + v47<<(uint(int32(4))%32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v47))))
	v62 = v60 - int32(98)
	if v62 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v119+v28<<(uint(int32(2))%32)))) = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v28))) = uint8(v127)
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = int32(-1)
	goto L9
L14:
	;
	if v62 != int32(18) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v51)+68))
	F_getTypeBinaryInputInfo(m, v92, v13+int32(12), v13+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51)+68))
	F_getTypeInputInfo(m, v65, v13+int32(12), v13+int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
	v76 = F_OidInputFunctionCall(m, v72, v73, v74, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v78+v28<<(uint(int32(2))%32)))) = v76
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v85 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v28))) = uint8(v85)
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = int32(-1)
	goto L9
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
	v102 = F_OidReceiveFunctionCall(m, v99, v57, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v28<<(uint(int32(2))%32)))) = v102
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v109 != v110 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112+v28))) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, _c_F_slot_store_data[0])) = int32(-1)
	goto L9
L23:
	;
	goto L8
L24:
	;
	m.G0 = v13 + int32(16)
	return
L25:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47 + int32(1)
	F_errmsg(m, int32(_a_F_slot_store_data_1), v13)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_slot_store_data_2), int32(847), int32(_a_F_slot_store_data_3))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
