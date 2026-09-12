package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BloomInitPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v2 = l1
	if l0&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l0, int32(0), int32(_a_F_BloomInitPage_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_BloomInitPage_1)
	v35 = int32(_a_F_BloomInitPage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v35)
	v41 = int32(_a_F_BloomInitPage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v45 = l0 + v44
	v46 = int32(_a_F_BloomInitPage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v46)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+2)) = uint16(v2)
	return
}
func F_BloomPageAddItem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1164))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v11 = l1 + v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	v13 = v9 * v12
	v14 = int32(_a_F_BloomPageAddItem_0) - v13
	if base.Ui32(v9) <= base.Ui32(v14) {
		if v9 != 0 {
			v19 = F__emscripten_memcpy_bulkmem(m, l1+v13+int32(24), l2, v9)
			mBase = m.M
		} else {
		}
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		v23 = v21 + int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v11))) = uint16(v23)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1164))
		v28 = v25*v23 + int32(24)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v28)
	} else {
	}
	return base.B2i32(base.Ui32(v9) <= base.Ui32(v14))
}
func F_bloomBuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int64
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	v8 = int32(_a_F_bloomBuildCallback_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
	*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v10
	v13 = l5 + int32(1184)
	v14 = F_BloomFormTuple(m, l5, l1, l2, l3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
		v23 = v13 + v22
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
		v25 = v21 * v24
		v26 = int32(_a_F_bloomBuildCallback_1) - v25
		if base.Ui32(v21) <= base.Ui32(v26) {
			v29 = int32(24)
			v31 = F___memcpy(m, v13+v25+v29, v14, v21)
			mBase = m.M
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
			v34 = v32 + int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v23))) = uint16(v34)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
			v39 = v36*v34 + v29
			*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v39)
		} else {
		}
		if base.Ui32(v21) <= base.Ui32(v26) {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = v43 + int32(1)
			v111 = *(*int64)(unsafe.Add(mBase, uint32(l5)+1168))
			*(*int64)(unsafe.Add(mBase, uint32(l5)+1168)) = v111 + int64(1)
			*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v9
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
			F_MemoryContextReset(m, v117)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return
			} else {
				return
			}
		} else {
			v47 = F_BloomNewBuffer(m, l0)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v49 = F_GenericXLogStart(m, l0)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					v52 = F_GenericXLogRegisterBuffer(m, v49, v47, int32(1))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v55 = F__emscripten_memcpy_bulkmem(m, v52, v13, int32(_a_F_bloomBuildCallback_2))
						mBase = m.M
						F_GenericXLogFinish(m, v49)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v47)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[2]))
								if v62 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										v65 = int32(0)
										F_PageInit(m, v13, int32(_a_F_bloomBuildCallback_2), int32(8))
										mBase = m.M
										v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
										v70 = v13 + v69
										v71 = int32(_a_F_bloomBuildCallback_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)) = uint16(v71)
										*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)) = uint16(v65)
										*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = int32(0)
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
										v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
										v83 = v13 + v82
										v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
										v85 = v81 * v84
										v86 = int32(_a_F_bloomBuildCallback_1) - v85
										if base.Ui32(v81) <= base.Ui32(v86) {
											v89 = int32(24)
											v91 = F___memcpy(m, v13+v85+v89, v14, v81)
											mBase = m.M
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
											v94 = v92 + int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v94)
											v96 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
											v99 = v96*v94 + v89
											*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v99)
										} else {
										}
										if base.B2i32(base.Ui32(v81) <= base.Ui32(v86)) == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_bloomBuildCallback_4), int32(0))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bloomBuildCallback_5), int32(104), int32(_a_F_bloomBuildCallback_6))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v105 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1])))
											*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = v105 + int32(1)
											v111 = *(*int64)(unsafe.Add(mBase, uint32(l5)+1168))
											*(*int64)(unsafe.Add(mBase, uint32(l5)+1168)) = v111 + int64(1)
											*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v9
											v117 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
											F_MemoryContextReset(m, v117)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									v65 = int32(0)
									F_PageInit(m, v13, int32(_a_F_bloomBuildCallback_2), int32(8))
									mBase = m.M
									v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
									v70 = v13 + v69
									v71 = int32(_a_F_bloomBuildCallback_3)
									*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)) = uint16(v71)
									*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)) = uint16(v65)
									*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = int32(0)
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
									v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
									v83 = v13 + v82
									v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
									v85 = v81 * v84
									v86 = int32(_a_F_bloomBuildCallback_1) - v85
									if base.Ui32(v81) <= base.Ui32(v86) {
										v89 = int32(24)
										v91 = F___memcpy(m, v13+v85+v89, v14, v81)
										mBase = m.M
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
										v94 = v92 + int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v94)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
										v99 = v96*v94 + v89
										*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v99)
									} else {
									}
									if base.B2i32(base.Ui32(v81) <= base.Ui32(v86)) == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(_a_F_bloomBuildCallback_4), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_bloomBuildCallback_5), int32(104), int32(_a_F_bloomBuildCallback_6))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1])))
										*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = v105 + int32(1)
										v111 = *(*int64)(unsafe.Add(mBase, uint32(l5)+1168))
										*(*int64)(unsafe.Add(mBase, uint32(l5)+1168)) = v111 + int64(1)
										*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v9
										v117 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
										F_MemoryContextReset(m, v117)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_bloom_create(m *base.Module, l0 int64, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v36 int32
	_ = v36
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v75 float64
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v9 = int64(1)
	v13 = base.I64_extend_i32_s(l1) << (uint(int64(10)) % 64)
	v15 = l0 << (uint(v9) % 64)
	if base.Ui64(v13) < base.Ui64(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v13
	goto L3
L2:
	;
	v17 = v15
	goto L3
L3:
	;
	if base.Ui64(v17) <= base.Ui64(int64(1048576)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int64(1048576)
	goto L6
L5:
	;
	v20 = v17
	goto L6
L6:
	;
	v22 = v20 << (uint(int64(3)) % 64)
	if v22 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v55 = int64(4294967295)
	goto L9
L8:
	;
	v28 = int32(-1)
	v30 = v22
	goto L10
L9:
	;
	v56 = v9 << (uint(v55) % 64)
	v62 = F_palloc0(m, base.I32_wrap_i64(int64(base.Ui64(v56)>>(uint(int64(3))%64)))+int32(24))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v36 = v28 + int32(1)
	if base.Ui64(v30) < base.Ui64(int64(2)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v55 = base.I64_extend_i32_u(v36)
	goto L9
L12:
	;
	goto L11
L13:
	;
	if v28 < int32(31) {
		v28 = v36
		v30 = int64(base.Ui64(v30) >> (uint(int64(1)) % 64))
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return int32(0)
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v62)+16)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = l2
	v75 = base.F64_nearest(base.F64_div(base.F64_mul(base.F64_convert_i64_u(v56), float64(0.6931471805599453)), base.F64_convert_i64_s(l0)))
	if base.F64_lt(base.F64_abs(v75), float64(2.147483648e+09)) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if int32(10) <= v81 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v79 = base.I32_trunc_f64_s(v75)
	v81 = v79
	goto L17
L19:
	;
	goto L20
L20:
	;
	v81 = int32(-2147483648)
	goto L17
L21:
	;
	v84 = int32(10)
	goto L23
L22:
	;
	v84 = v81
	goto L23
L23:
	;
	if v84 <= int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v87 = int32(1)
	goto L26
L25:
	;
	v87 = v84
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v87
	return v62
}
