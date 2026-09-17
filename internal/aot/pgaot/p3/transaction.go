package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TransactionIdFollowsOrEquals(m *base.Module, l0 int32, l1 int32) int32 {
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		return base.B2i32(base.Ui32(l1) <= base.Ui32(l0))
	} else {
		return base.B2i32(int32(0) <= l0-l1)
	}
}
func F_TransactionIdGetStatus(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v10 = int32(base.Ui32(l0) >> (uint(int32(15)) % 32))
	v12 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_TransactionIdGetStatus_0), base.I64_extend_i32_u(v10), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = int32(2)
		v19 = int32(base.Ui32(l0&int32(_a_F_TransactionIdGetStatus_1)) >> (uint(v18) % 32))
		v20 = int32(_a_F_TransactionIdGetStatus_0)
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetStatus[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v12<<(uint(v18)%32))))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v26))))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v29+v12<<(uint(int32(13))%32)+v19&int32(_a_F_TransactionIdGetStatus_2))))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v36
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetStatus[0]))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
		v42 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdGetStatus[1])))
		v43 = base.I32_rem_u_s(v10, v42)
		F_LWLockRelease(m, v40+v43<<(uint(int32(7))%32))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			return int32(base.Ui32(v28)>>(uint(l0<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
		}
	}
}
func F_TransactionIdIsCurrentTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[0]))
	if v13 == l0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[1]))
	if v18 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v126
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[2]))
	if v22 == int32(0) {
		v126 = int32(0)
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsCurrentTransactionId[3]))
	v97 = int32(0)
	v99 = v18 - int32(1)
	goto L34
L11:
	;
	v27 = v22
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v32 == int32(4) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v126 = int32(0)
	goto L7
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	if v90 != 0 {
		v27 = v90
		goto L12
	} else {
		goto L33
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(1)
	if l0 == v35 {
		v126 = v38
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v42 = v40 - int32(1)
	if v42 < int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(0)
	v49 = v42
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	v55 = int32(2)
	v56 = base.I32_div_s(v49-v47, v55)
	v57 = v56 + v47
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53+v57<<(uint(v55)%32))))
	if v61 == l0 {
		v126 = v38
		goto L7
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v61)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v76 = base.B2i32(base.Ui32(v61) < base.Ui32(l0))
	goto L22
L24:
	;
	goto L25
L25:
	;
	v76 = int32(base.Ui32(v61-l0) >> (uint(int32(31)) % 32))
	goto L22
L26:
	;
	v77 = v57 + int32(1)
	goto L28
L27:
	;
	v77 = v47
	goto L28
L28:
	;
	if v76 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v80 = v49
	goto L31
L30:
	;
	v80 = v57 - int32(1)
	goto L31
L31:
	;
	if v77 <= v80 {
		v47 = v77
		v49 = v80
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	goto L13
L34:
	;
	v104 = int32(2)
	v105 = base.I32_div_s(v99-v97, v104)
	v106 = v105 + v97
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v95+v106<<(uint(v104)%32))))
	v111 = base.B2i32(v110 == l0)
	if v110 == l0 {
		v126 = v111
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v126 = v111
	goto L7
L36:
	;
	v114 = base.B2i32(base.Ui32(v110) < base.Ui32(l0))
	if base.Ui32(v110) < base.Ui32(l0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v115 = v106 + int32(1)
	goto L39
L38:
	;
	v115 = v97
	goto L39
L39:
	;
	if base.Ui32(v110) < base.Ui32(l0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v118 = v99
	goto L42
L41:
	;
	v118 = v106 - int32(1)
	goto L42
L42:
	;
	if v115 <= v118 {
		v97 = v115
		v99 = v118
		goto L34
	} else {
		goto L43
	}
L43:
	;
	goto L35
}
func F_TransactionIdSetPageStatusInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v92 int32
	_ = v92
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	v17 = base.B2i32(l4 == int64(0))
	v18 = F_SimpleLruReadPage(m, int32(_a_F_TransactionIdSetPageStatusInternal_0), l5, v17, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if l0 == int32(0) {
		} else {
			v24 = int32(0)
			if base.B2i32(l3 != int32(1))|base.B2i32(l1 <= v24) == v24 {
				v37 = int32(0)
				for {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[0])))
					v47 = int32(1)
					v49 = int32(2)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l2+v37<<(uint(v49)%32))))
					v56 = int32(base.Ui32(v52&int32(_a_F_TransactionIdSetPageStatusInternal_1)) >> (uint(v49) % 32))
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v18<<(uint(v49)%32))))
					v64 = v56 + v63
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
					v69 = v52 << (uint(v47) % 32) & int32(6)
					if base.B2i32(v46 == v47)&base.B2i32(int32(base.Ui32(v65)>>(uint(v69)%32))&int32(3) == v47) != 0 {
					} else {
						v78 = v65 | int32(3)<<(uint(v69)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v78)
						if l4 == int64(0) {
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+36))
							v86 = v82 + v18<<(uint(int32(13))%32) + v56&int32(_a_F_TransactionIdSetPageStatusInternal_2)
							v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
							if base.Ui64(l4) <= base.Ui64(v87) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v86))) = l4
							}
						}
					}
					v92 = v37 + int32(1)
					if v92 != l1 {
						v37 = v92
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v109 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
			v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
			v111 = int32(2)
			v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v18<<(uint(v111)%32))))
			v116 = l0 & int32(_a_F_TransactionIdSetPageStatusInternal_1)
			v119 = v114 + int32(base.Ui32(v116)>>(uint(v111)%32))
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v124 = l0 << (uint(int32(1)) % 32) & int32(6)
			if l3 != int32(3) {
				v144 = v120&(int32(3)<<(uint(v124)%32)^int32(-1)) | l3<<(uint(v124)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v144)
				if l4 == int64(0) {
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
					v156 = v148 + v18<<(uint(int32(13))%32) + int32(base.Ui32(v116)>>(uint(int32(2))%32))&int32(_a_F_TransactionIdSetPageStatusInternal_2)
					v157 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
					if base.Ui64(l4) <= base.Ui64(v157) {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v156))) = l4
					}
				}
			} else {
				v128 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[0])))
				if v128&int32(1) == int32(0) {
					v144 = v120&(int32(3)<<(uint(v124)%32)^int32(-1)) | l3<<(uint(v124)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v144)
					if l4 == int64(0) {
					} else {
						v147 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
						v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
						v156 = v148 + v18<<(uint(int32(13))%32) + int32(base.Ui32(v116)>>(uint(int32(2))%32))&int32(_a_F_TransactionIdSetPageStatusInternal_2)
						v157 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
						if base.Ui64(l4) <= base.Ui64(v157) {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v156))) = l4
						}
					}
				} else {
					if int32(base.Ui32(v120)>>(uint(v124)%32))&int32(3) == int32(1) {
					} else {
						v144 = v120&(int32(3)<<(uint(v124)%32)^int32(-1)) | l3<<(uint(v124)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v144)
						if l4 == int64(0) {
						} else {
							v147 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
							v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
							v156 = v148 + v18<<(uint(int32(13))%32) + int32(base.Ui32(v116)>>(uint(int32(2))%32))&int32(_a_F_TransactionIdSetPageStatusInternal_2)
							v157 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
							if base.Ui64(l4) <= base.Ui64(v157) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v156))) = l4
							}
						}
					}
				}
			}
		}
		if int32(0) < l1 {
			v185 = int32(0)
			for {
				v193 = int32(2)
				v196 = *(*int32)(unsafe.Add(mBase, uint32(l2+v185<<(uint(v193)%32))))
				v200 = int32(base.Ui32(v196&int32(_a_F_TransactionIdSetPageStatusInternal_1)) >> (uint(v193) % 32))
				v202 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v203+v18<<(uint(v193)%32))))
				v208 = v200 + v207
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v213 = v196 << (uint(int32(1)) % 32) & int32(6)
				if l3 != int32(3) {
					v233 = v209&(int32(3)<<(uint(v213)%32)^int32(-1)) | l3<<(uint(v213)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v233)
					if l4 == int64(0) {
					} else {
						v236 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
						v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+36))
						v241 = v237 + v18<<(uint(int32(13))%32) + v200&int32(_a_F_TransactionIdSetPageStatusInternal_2)
						v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
						if base.Ui64(l4) <= base.Ui64(v242) {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v241))) = l4
						}
					}
				} else {
					v217 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[0])))
					if v217&int32(1) == int32(0) {
						v233 = v209&(int32(3)<<(uint(v213)%32)^int32(-1)) | l3<<(uint(v213)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v233)
						if l4 == int64(0) {
						} else {
							v236 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
							v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+36))
							v241 = v237 + v18<<(uint(int32(13))%32) + v200&int32(_a_F_TransactionIdSetPageStatusInternal_2)
							v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
							if base.Ui64(l4) <= base.Ui64(v242) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v241))) = l4
							}
						}
					} else {
						if int32(base.Ui32(v209)>>(uint(v213)%32))&int32(3) == int32(1) {
						} else {
							v233 = v209&(int32(3)<<(uint(v213)%32)^int32(-1)) | l3<<(uint(v213)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v233)
							if l4 == int64(0) {
							} else {
								v236 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
								v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+36))
								v241 = v237 + v18<<(uint(int32(13))%32) + v200&int32(_a_F_TransactionIdSetPageStatusInternal_2)
								v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
								if base.Ui64(l4) <= base.Ui64(v242) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v241))) = l4
								}
							}
						}
					}
				}
				v247 = v185 + int32(1)
				if v247 != l1 {
					v185 = v247
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v264 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetPageStatusInternal[1]))
		v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
		v267 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v265+v18))) = uint8(v267)
		return
	}
}
func F_TransactionTimeoutHandler(m *base.Module) {
	var v3 int32
	_ = v3
	Fn13831(m, int32(_a_F_TransactionTimeoutHandler_0))
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_UnlockApplyTransactionForSession(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = l2
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(267)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v10)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockApplyTransactionForSession[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	v19 = F_LockRelease(m, v8, l3, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_check_transaction_read_only(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		v70 = v4
		return v70
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[0])))
		if v7&int32(1) == int32(0) {
			v70 = v4
			return v70
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[1]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
			if base.B2i32(v14 == int32(2)) == int32(0) {
				return int32(1)
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[2])))
				if v22&int32(1) != 0 {
					v70 = v4
					return v70
				} else {
					v25 = int32(16777538)
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[1]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
					if int32(1) < v29 {
						v53 = v25
						v54 = int32(_a_F_check_transaction_read_only_0)
						*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[3])) = v53
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[5])) = v59
						v64 = F_format_elog_string(m, v54, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[6])) = v64
							v70 = int32(0)
							return v70
						}
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[7])))
						if v34 != 0 {
							v53 = v25
							v54 = int32(_a_F_check_transaction_read_only_1)
							*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[3])) = v53
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[5])) = v59
							v64 = F_format_elog_string(m, v54, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[6])) = v64
								v70 = int32(0)
								return v70
							}
						} else {
							v35 = int32(1)
							v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[8])))
							if v38 == v35 {
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[9]))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+316))
								v46 = base.B2i32(v44 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _c_F_check_transaction_read_only[8])) = uint8(v46)
								v48 = v46
							} else {
								v48 = int32(0)
							}
							if v48 == int32(0) {
								v70 = v35
								return v70
							} else {
								v53 = int32(1088)
								v54 = int32(_a_F_check_transaction_read_only_2)
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[3])) = v53
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[5])) = v59
								v64 = F_format_elog_string(m, v54, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_check_transaction_read_only[6])) = v64
									v70 = int32(0)
									return v70
								}
							}
						}
					}
				}
			}
		}
	}
}
