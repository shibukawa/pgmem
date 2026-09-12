package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistXLogUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	v3 = l2
	v5 = l4
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v5)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)) = uint16(v3)
	F_XLogBeginInsert(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_XLogRegisterData(m, v12+int32(12), int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_XLogRegisterBuffer(m, int32(0), l0, int32(8))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_XLogRegisterBufData(m, int32(0), l1, v3<<(uint(int32(1))%32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(0) < v5 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	if l5 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3+v43<<(uint(int32(2))%32))))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+6)))
	F_XLogRegisterBufData(m, int32(0), v49, v50&int32(8191))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v56 = v43 + int32(1)
	if v56 != v5 {
		v43 = v56
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_XLogRegisterBuffer(m, int32(1), l5, int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v73 = F_XLogInsert(m, int32(14), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v12 + int32(16)
	return v73
}
func F_gist_box_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 == v2 {
		v49 = v2
		m.G0 = v9 + int32(16)
		return v49
	} else {
		if v12 == int32(0) {
			v49 = v2
			m.G0 = v9 + int32(16)
			return v49
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23)+12)))
			if v25&int32(1) != 0 {
				v31 = (v11 - int32(1)) & int32(65535)
				if base.Ui32(int32(12)) <= base.Ui32(v31) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg_internal(m, int32(461458), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(478117), int32(939), int32(87816))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32))+uint32(_consts[22])))
					v40 = F_DirectFunctionCall2Coll(m, v38, int32(0), v17, v12)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v49 = base.B2i32(v40 != int32(0))
						m.G0 = v9 + int32(16)
						return v49
					}
				}
			} else {
				v46 = F_rtree_internal_consistent(m, v17, v12, v11)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = v46
					m.G0 = v9 + int32(16)
					return v49
				}
			}
		}
	}
}
func F_gist_box_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 float64
	_ = v15
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 float64
	_ = v19
	var v27 int32
	_ = v27
	var v33 float64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 float64
	_ = v37
	var v45 int32
	_ = v45
	var v51 float64
	_ = v51
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 float64
	_ = v55
	var v63 int32
	_ = v63
	var v69 float64
	_ = v69
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 float64
	_ = v73
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == v2 {
		v92 = base.B2i32(v10|v9 == int32(0))
		*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
		return v8
	} else {
		if v9 == int32(0) {
			v92 = base.B2i32(v10|v9 == int32(0))
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
			return v8
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
			v17 = int64(9223372036854775807)
			v18 = base.I64_reinterpret_f64(v15) & v17
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v19)&v17) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v18) {
					v33 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
					v35 = int64(9223372036854775807)
					v36 = base.I64_reinterpret_f64(v33) & v35
					v37 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v37)&v35) {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
							v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
							v53 = int64(9223372036854775807)
							v54 = base.I64_reinterpret_f64(v51) & v53
							v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
									v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
									v71 = int64(9223372036854775807)
									v72 = base.I64_reinterpret_f64(v69) & v71
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
										return v8
									} else {
										v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
										return v8
									}
								} else {
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
									return v8
								}
							} else {
								if base.F64_ne(v51, v55) != 0 {
									v92 = v2
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
									return v8
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
										v92 = v2
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
										return v8
									} else {
										v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
										v71 = int64(9223372036854775807)
										v72 = base.I64_reinterpret_f64(v69) & v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
											return v8
										} else {
											v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
											return v8
										}
									}
								}
							}
						} else {
							v45 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v45)
							return v8
						}
					} else {
						if base.F64_ne(v33, v37) != 0 {
							v92 = v2
							*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
							return v8
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
								v92 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
								return v8
							} else {
								v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
								v53 = int64(9223372036854775807)
								v54 = base.I64_reinterpret_f64(v51) & v53
								v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
										v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
										v71 = int64(9223372036854775807)
										v72 = base.I64_reinterpret_f64(v69) & v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
											return v8
										} else {
											v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
											return v8
										}
									} else {
										v63 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
										return v8
									}
								} else {
									if base.F64_ne(v51, v55) != 0 {
										v92 = v2
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
										return v8
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
											v92 = v2
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
											return v8
										} else {
											v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
											v71 = int64(9223372036854775807)
											v72 = base.I64_reinterpret_f64(v69) & v71
											v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
												return v8
											} else {
												v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
												return v8
											}
										}
									}
								}
							}
						}
					}
				} else {
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v27)
					return v8
				}
			} else {
				if base.F64_ne(v15, v19) != 0 {
					v92 = v2
					*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
					return v8
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v18) {
						v92 = v2
						*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
						return v8
					} else {
						v33 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						v35 = int64(9223372036854775807)
						v36 = base.I64_reinterpret_f64(v33) & v35
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v37)&v35) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
								v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
								v53 = int64(9223372036854775807)
								v54 = base.I64_reinterpret_f64(v51) & v53
								v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
										v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
										v71 = int64(9223372036854775807)
										v72 = base.I64_reinterpret_f64(v69) & v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
											return v8
										} else {
											v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
											return v8
										}
									} else {
										v63 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
										return v8
									}
								} else {
									if base.F64_ne(v51, v55) != 0 {
										v92 = v2
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
										return v8
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
											v92 = v2
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
											return v8
										} else {
											v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
											v71 = int64(9223372036854775807)
											v72 = base.I64_reinterpret_f64(v69) & v71
											v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
												return v8
											} else {
												v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
												return v8
											}
										}
									}
								}
							} else {
								v45 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v45)
								return v8
							}
						} else {
							if base.F64_ne(v33, v37) != 0 {
								v92 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
								return v8
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
									v92 = v2
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
									return v8
								} else {
									v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
									v53 = int64(9223372036854775807)
									v54 = base.I64_reinterpret_f64(v51) & v53
									v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
											v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
											v71 = int64(9223372036854775807)
											v72 = base.I64_reinterpret_f64(v69) & v71
											v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
												return v8
											} else {
												v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
												return v8
											}
										} else {
											v63 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
											return v8
										}
									} else {
										if base.F64_ne(v51, v55) != 0 {
											v92 = v2
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
											return v8
										} else {
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
												v92 = v2
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
												return v8
											} else {
												v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
												v71 = int64(9223372036854775807)
												v72 = base.I64_reinterpret_f64(v69) & v71
												v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
													*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
													return v8
												} else {
													v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
													*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
													return v8
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
	}
}
func F_gist_circle_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if base.Ui32(int32(20)) <= base.Ui32(v9) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
			F_errmsg_internal(m, int32(461458), v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(478117), int32(1492), int32(398921))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v32 = F_computeDistance(m, int32(0), v30, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v34)
			v36 = F_Float8GetDatum(m, v32)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v36
			}
		}
	}
}
func F_gist_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = int32(0)
	if v2 <= base.I32_extend8_s(l0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_consts[57])))
		v13 = v12
	} else {
		v13 = v2
	}
	return v13
}
func F_gist_point_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if base.Ui32(v9) <= base.Ui32(int32(19)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v14)+12)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_computeDistance(m, v16&int32(1), v19, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = F_Float8GetDatum(m, v21)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v25
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
			F_errmsg_internal(m, int32(461458), v7)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(478117), int32(1470), int32(398958))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
