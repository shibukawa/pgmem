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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v2 = l1
	v3 = int32(_a_F_BloomInitPage_0)
	v5 = int32(0)
	if v5|(l0&int32(3)|int32(1)) == v5 {
		v21 = l0 + v3
		v23 = l0 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_BloomInitPage_1)
	v44 = int32(_a_F_BloomInitPage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = int32(_a_F_BloomInitPage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	v55 = int32(_a_F_BloomInitPage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v55)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)) = uint16(v2)
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1164))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v11 = l1 + v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	v13 = v9 * v12
	v14 = int32(_a_F_BloomPageAddItem_0) - v13
	if base.Ui32(v9) <= base.Ui32(v14) {
		if v9 != 0 {
			base.MemoryCopy(m, l1+v13+int32(24), l2, v9)
		} else {
		}
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		v22 = v20 + int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v11))) = uint16(v22)
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+1164)))
		v27 = v22*v24 + int32(24)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v27)
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v8 = int32(_a_F_bloomBuildCallback_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
	*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v11
	v14 = l5 + int32(1184)
	v15 = F_BloomFormTuple(m, l5, l1, l2, l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+16)))
		v24 = v14 + v23
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
		v26 = v22 * v25
		v27 = int32(_a_F_bloomBuildCallback_1) - v26
		if base.Ui32(v22) <= base.Ui32(v27) {
			if v22 != 0 {
				base.MemoryCopy(m, v14+v26+int32(24), v15, v22)
			} else {
			}
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
			v35 = v33 + int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v35)
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+1164)))
			v40 = v35*v37 + int32(24)
			*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v40)
		} else {
		}
		if base.Ui32(v22) <= base.Ui32(v27) {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = v103 + int32(1)
			v107 = *(*int64)(unsafe.Add(mBase, uint32(l5)+1168))
			*(*int64)(unsafe.Add(mBase, uint32(l5)+1168)) = v107 + int64(1)
			*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v9
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
			F_MemoryContextReset(m, v113)
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return
			} else {
				return
			}
		} else {
			v44 = F_BloomNewBuffer(m, l0)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v46 = F_GenericXLogStart(m, l0)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v49 = F_GenericXLogRegisterBuffer(m, v46, v44, int32(1))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						base.MemoryCopy(m, v49, v14, int32(_a_F_bloomBuildCallback_2))
						F_GenericXLogFinish(m, v46)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v44)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[2]))
								if v58 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v61 = int32(0)
										F_PageInit(m, v14, int32(_a_F_bloomBuildCallback_2), int32(8))
										mBase = m.M
										v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+16)))
										v66 = v14 + v65
										v67 = int32(_a_F_bloomBuildCallback_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v67)
										*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v61)
										*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = int32(0)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
										v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+16)))
										v79 = v14 + v78
										v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
										v81 = v77 * v80
										v82 = int32(_a_F_bloomBuildCallback_1) - v81
										if base.Ui32(v77) <= base.Ui32(v82) {
											if v77 != 0 {
												base.MemoryCopy(m, v14+v81+int32(24), v15, v77)
											} else {
											}
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
											v90 = v88 + int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v90)
											v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+1164)))
											v95 = v90*v92 + int32(24)
											*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v95)
										} else {
										}
										if base.B2i32(base.Ui32(v77) <= base.Ui32(v82)) == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_bloomBuildCallback_4), int32(0))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bloomBuildCallback_5), int32(104), int32(_a_F_bloomBuildCallback_6))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1])))
											*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = v103 + int32(1)
											v107 = *(*int64)(unsafe.Add(mBase, uint32(l5)+1168))
											*(*int64)(unsafe.Add(mBase, uint32(l5)+1168)) = v107 + int64(1)
											*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v9
											v113 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
											F_MemoryContextReset(m, v113)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									v61 = int32(0)
									F_PageInit(m, v14, int32(_a_F_bloomBuildCallback_2), int32(8))
									mBase = m.M
									v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+16)))
									v66 = v14 + v65
									v67 = int32(_a_F_bloomBuildCallback_3)
									*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v67)
									*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v61)
									*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = int32(0)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1164))
									v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+16)))
									v79 = v14 + v78
									v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
									v81 = v77 * v80
									v82 = int32(_a_F_bloomBuildCallback_1) - v81
									if base.Ui32(v77) <= base.Ui32(v82) {
										if v77 != 0 {
											base.MemoryCopy(m, v14+v81+int32(24), v15, v77)
										} else {
										}
										v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
										v90 = v88 + int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v90)
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+1164)))
										v95 = v90*v92 + int32(24)
										*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v95)
									} else {
									}
									if base.B2i32(base.Ui32(v77) <= base.Ui32(v82)) == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(_a_F_bloomBuildCallback_4), int32(0))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_bloomBuildCallback_5), int32(104), int32(_a_F_bloomBuildCallback_6))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1])))
										*(*int32)(unsafe.Add(mBase, uint32(l5)+uint32(_c_F_bloomBuildCallback[1]))) = v103 + int32(1)
										v107 = *(*int64)(unsafe.Add(mBase, uint32(l5)+1168))
										*(*int64)(unsafe.Add(mBase, uint32(l5)+1168)) = v107 + int64(1)
										*(*int32)(unsafe.Add(mBase, _c_F_bloomBuildCallback[0])) = v9
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l5)+1176))
										F_MemoryContextReset(m, v113)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
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
	var v8 int64
	_ = v8
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	v8 = int64(1)
	v12 = base.I64_extend_i32_s(l1) << (uint(int64(10)) % 64)
	v14 = l0 << (uint(v8) % 64)
	if base.Ui64(v12) < base.Ui64(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v12
	goto L3
L2:
	;
	v16 = v14
	goto L3
L3:
	;
	if base.Ui64(v16) <= base.Ui64(int64(1048576)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = int64(1048576)
	goto L6
L5:
	;
	v19 = v16
	goto L6
L6:
	;
	v21 = v19 << (uint(int64(3)) % 64)
	if v21 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v50 = int64(4294967295)
	goto L9
L8:
	;
	v27 = int32(-1)
	v29 = v21
	goto L10
L9:
	;
	v51 = v8 << (uint(v50) % 64)
	v57 = F_palloc0(m, base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(3))%64)))+int32(24))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v34 = v27 + int32(1)
	v36 = int64(base.Ui64(v29) >> (uint(int64(1)) % 64))
	if v36 == int64(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v50 = base.I64_extend_i32_u(v34)
	goto L9
L12:
	;
	goto L11
L13:
	;
	if v27 < int32(31) {
		v27 = v34
		v29 = v36
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
	*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = l2
	v64 = int32(1)
	v71 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(base.F64_mul(base.F64_convert_i64_u(v51), float64(0.6931471805599453)), base.F64_convert_i64_s(l0))))
	if v71 <= v64 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v74 = v64
	goto L19
L18:
	;
	v74 = v71
	goto L19
L19:
	;
	if int32(10) <= v74 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v77 = int32(10)
	goto L22
L21:
	;
	v77 = v74
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v77
	return v57
}
