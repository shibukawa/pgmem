package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BitJaccardDistanceDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	v7 = int32(0)
	if base.Ui32(l0) < base.Ui32(int32(8)) {
		v89 = l3
		v90 = l4
		v91 = l5
		v92 = l0
		v93 = l1
		v94 = l2
	} else {
		v16 = int32(8)
		v17 = l0 - v16
		if v17&v16 == int32(0) {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v31 = int32(8)
			v32 = l2 + v31
			v34 = l1 + v31
			v35 = v17
			v36 = v34
			v37 = v32
			v38 = base.I64_popcnt(v22&v25) + l3
			v39 = base.I64_popcnt(v25) + l4
			v40 = base.I64_popcnt(v22) + l5
			v41 = v34
			v42 = v32
		} else {
			v35 = l0
			v36 = l1
			v37 = l2
			v38 = l3
			v39 = l4
			v40 = l5
			v41 = v7
			v42 = v7
		}
		if base.Ui32(v17) <= base.Ui32(int32(7)) {
			v89 = v38
			v90 = v39
			v91 = v40
			v92 = v17
			v93 = v41
			v94 = v42
		} else {
			v50 = v38
			v51 = v39
			v52 = v40
			v53 = v35
			v54 = v36
			v55 = v37
			for {
				v60 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
				v64 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
				v69 = base.I64_popcnt(v60&v61) + (base.I64_popcnt(v64&v65) + v50)
				v70 = int32(16)
				v71 = v55 + v70
				v73 = v54 + v70
				v77 = base.I64_popcnt(v60) + (base.I64_popcnt(v64) + v52)
				v81 = base.I64_popcnt(v61) + (base.I64_popcnt(v65) + v51)
				v83 = v53 - v70
				if base.Ui32(int32(7)) < base.Ui32(v83) {
					v50 = v69
					v51 = v81
					v52 = v77
					v53 = v83
					v54 = v73
					v55 = v71
					continue
				} else {
					break
				}
				break
			}
			v89 = v69
			v90 = v81
			v91 = v77
			v92 = v83
			v93 = v73
			v94 = v71
		}
	}
	if v92 != 0 {
		v100 = int32(0)
		v103 = v89
		v104 = v90
		v105 = v91
		for {
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v94))))
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v93))))
			v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v114&v116)+uint32(_consts[1115]))))
			v121 = v103 + v120
			v124 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_consts[1115]))))
			v125 = v105 + v124
			v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v116)+uint32(_consts[1115]))))
			v129 = v104 + v128
			v131 = v100 + int32(1)
			if v131 != v92 {
				v100 = v131
				v103 = v121
				v104 = v129
				v105 = v125
				continue
			} else {
				break
			}
			break
		}
		v136 = v121
		v137 = v129
		v138 = v125
	} else {
		v136 = v89
		v137 = v90
		v138 = v91
	}
	if v136 == int64(0) {
		return float64(1)
	} else {
		return base.F64_sub(float64(1), base.F64_div(base.F64_convert_i64_u(v136), base.F64_convert_i64_u(v137-v136+v138)))
	}
}
func F_BitSumCenter(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 float32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 float32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v80 float32
	_ = v80
	var v84 int32
	_ = v84
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= v3 {
	} else {
		v13 = int32(1)
		v16 = l0 + int32(8)
		v17 = int32(0)
		if v10 != v13 {
			v22 = v17
			v26 = v3
			for {
				v31 = int32(2)
				v33 = l1 + v22<<(uint(v31)%32)
				v34 = *(*float32)(unsafe.Add(mBase, uint32(v33)))
				v37 = v16 + int32(base.Ui32(v22)>>(uint(int32(3))%32))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
				v39 = int32(6)
				v40 = v22 & v39
				v44 = int32(1)
				*(*float32)(unsafe.Add(mBase, uint32(v33))) = base.F32_add(v34, base.F32_convert_i32_u(int32(base.Ui32(v38)>>(uint(v40^int32(7))%32))&v44))
				v50 = v33 + int32(4)
				v51 = *(*float32)(unsafe.Add(mBase, uint32(v50)))
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
				*(*float32)(unsafe.Add(mBase, uint32(v50))) = base.F32_add(v51, base.F32_convert_i32_u(int32(base.Ui32(v52)>>(uint(v40^v39)%32))&v44))
				v62 = v22 + v31
				v64 = v26 + v31
				if v64 != v10&int32(2147483646) {
					v22 = v62
					v26 = v64
					continue
				} else {
					break
				}
				break
			}
			v66 = v62
		} else {
			v66 = v17
		}
		if v10&v13 == int32(0) {
		} else {
			v79 = l1 + v66<<(uint(int32(2))%32)
			v80 = *(*float32)(unsafe.Add(mBase, uint32(v79)))
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v66)>>(uint(int32(3))%32))))))
			*(*float32)(unsafe.Add(mBase, uint32(v79))) = base.F32_add(v80, base.F32_convert_i32_u(int32(base.Ui32(v84)>>(uint((v66^int32(-1))&int32(7))%32))&int32(1)))
		}
	}
	return
}
func F_BitUpdateCenter(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 float32
	_ = v57
	var v58 float32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 float32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 float32
	_ = v99
	var v107 int32
	_ = v107
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	v14 = int32(8)
	v15 = base.I32_div_s(l1+int32(7), v14)
	v16 = int32(2)
	v19 = v15<<(uint(v16)%32) + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
	v22 = l0 + v14
	v26 = int32(base.Ui32(v19)>>(uint(v16)%32)) - v14
	if v26 != 0 {
		v29 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), v26)
		mBase = m.M
	} else {
	}
	if l1 <= int32(0) {
	} else {
		v32 = int32(1)
		v34 = int32(0)
		if l1 != v32 {
			v40 = v34
			v43 = int32(0)
			for {
				v52 = v22 + int32(base.Ui32(v40)>>(uint(int32(3))%32))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				v54 = int32(2)
				v56 = l2 + v40<<(uint(v54)%32)
				v57 = *(*float32)(unsafe.Add(mBase, uint32(v56)))
				v58 = float32(0.5)
				v60 = int32(6)
				v61 = v40 & v60
				v65 = v53 | base.F32_gt(v57, v58)<<(uint(v61^int32(7))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v65)
				v67 = *(*float32)(unsafe.Add(mBase, uint32(v56)+4))
				v73 = v65 | base.F32_gt(v67, v58)<<(uint(v61^v60)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v73)
				v76 = v40 + v54
				v78 = v43 + v54
				if v78 != l1&int32(2147483646) {
					v40 = v76
					v43 = v78
					continue
				} else {
					break
				}
				break
			}
			v80 = v76
		} else {
			v80 = v34
		}
		if l1&v32 == int32(0) {
		} else {
			v94 = v22 + int32(base.Ui32(v80)>>(uint(int32(3))%32))
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
			v99 = *(*float32)(unsafe.Add(mBase, uint32(l2+v80<<(uint(int32(2))%32))))
			v107 = v95 | base.F32_gt(v99, float32(0.5))<<(uint((v80^int32(-1))&int32(7))%32)
			*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v107)
		}
	}
	return
}
func F_bit_and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v15 == v16 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v21 = F_palloc(m, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v15
					v24 = int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18 & v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v27&v24 != int32(32) {
						v32 = int32(8)
						v38 = v8 + v32
						v39 = v13 + v32
						v41 = v21 + v32
						v43 = int32(0)
						for {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
							v46 = v44 & v45
							*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
							v48 = int32(1)
							v55 = v43 + v48
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							if base.Ui32(v55) < base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(8)) {
								v38 = v38 + v48
								v39 = v39 + v48
								v41 = v41 + v48
								v43 = v55
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v21
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(160707), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(504270), int32(1261), int32(438858))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
		}
	}
}
func F_bit_or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v15 == v16 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v21 = F_palloc(m, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v15
					v24 = int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18 & v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v27&v24 != int32(32) {
						v32 = int32(8)
						v38 = v8 + v32
						v39 = v13 + v32
						v41 = v21 + v32
						v43 = int32(0)
						for {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
							v46 = v44 | v45
							*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
							v48 = int32(1)
							v55 = v43 + v48
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							if base.Ui32(v55) < base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(8)) {
								v38 = v38 + v48
								v39 = v39 + v48
								v41 = v41 + v48
								v43 = v55
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v21
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(160666), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(504270), int32(1302), int32(217386))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
		}
	}
}
