package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HalfvecCosineSimilarityDefault(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 float32
	_ = v6
	var v18 int32
	_ = v18
	var v22 float32
	_ = v22
	var v23 float32
	_ = v23
	var v24 float32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 float32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 float32
	_ = v210
	var v212 float32
	_ = v212
	var v214 float32
	_ = v214
	var v216 float32
	_ = v216
	var v218 int32
	_ = v218
	v4 = int32(0)
	v6 = float32(0)
	if l0 <= v4 {
		return math.Float64frombits(uint64(0x7ff8000000000000))
	} else {
		v18 = v4
		v22 = v6
		v23 = v6
		v24 = v6
		for {
			v26 = v18 << (uint(int32(1)) % 32)
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v26))))
			v33 = v28 & int32(1023)
			v37 = v28 << (uint(int32(16)) % 32) & int32(-2147483648)
			v40 = int32(31)
			v41 = int32(base.Ui32(v28)>>(uint(int32(10))%32)) & v40
			if v41 != v40 {
				if v41 != 0 {
					v113 = v33
					v114 = v41<<(uint(int32(23))%32) + v37 + int32(939524096)
				} else {
					if v33 != 0 {
						if v28&int32(512) != 0 {
							v101 = v33 << (uint(int32(1)) % 32)
							v103 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v33) {
								v101 = v33 << (uint(int32(2)) % 32)
								v103 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v33) {
									v101 = v33 << (uint(int32(3)) % 32)
									v103 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v33) {
										v101 = v33 << (uint(int32(4)) % 32)
										v103 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v33) {
											v101 = v33 << (uint(int32(5)) % 32)
											v103 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v33) {
												v101 = v33 << (uint(int32(6)) % 32)
												v103 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v33) {
													v101 = v33 << (uint(int32(7)) % 32)
													v103 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v33) {
														v101 = v33 << (uint(int32(8)) % 32)
														v103 = int32(880803840)
													} else {
														v96 = base.B2i32(v33 == int32(1))
														if v33 == int32(1) {
															v97 = int32(1024)
														} else {
															v97 = v33 << (uint(int32(9)) % 32)
														}
														if v33 == int32(1) {
															v100 = int32(864026624)
														} else {
															v100 = int32(872415232)
														}
														v101 = v97
														v103 = v100
													}
												}
											}
										}
									}
								}
							}
						}
						v113 = v101 & int32(1022)
						v114 = v103 | v37
					} else {
						v113 = int32(0)
						v114 = v37
					}
				}
			} else {
				if v33 == int32(0) {
					v113 = int32(0)
					v114 = v37 | int32(2139095040)
				} else {
					v113 = v33
					v114 = v37 | int32(2143289344)
				}
			}
			v118 = base.F32_reinterpret_i32(v114 | v113<<(uint(int32(13))%32))
			v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v26))))
			v125 = v120 & int32(1023)
			v129 = v120 << (uint(int32(16)) % 32) & int32(-2147483648)
			v132 = int32(31)
			v133 = int32(base.Ui32(v120)>>(uint(int32(10))%32)) & v132
			if v133 != v132 {
				if v133 != 0 {
					v205 = v125
					v206 = v133<<(uint(int32(23))%32) + v129 + int32(939524096)
				} else {
					if v125 != 0 {
						if v120&int32(512) != 0 {
							v193 = v125 << (uint(int32(1)) % 32)
							v195 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v125) {
								v193 = v125 << (uint(int32(2)) % 32)
								v195 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v125) {
									v193 = v125 << (uint(int32(3)) % 32)
									v195 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v125) {
										v193 = v125 << (uint(int32(4)) % 32)
										v195 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v125) {
											v193 = v125 << (uint(int32(5)) % 32)
											v195 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v125) {
												v193 = v125 << (uint(int32(6)) % 32)
												v195 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v125) {
													v193 = v125 << (uint(int32(7)) % 32)
													v195 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v125) {
														v193 = v125 << (uint(int32(8)) % 32)
														v195 = int32(880803840)
													} else {
														v188 = base.B2i32(v125 == int32(1))
														if v125 == int32(1) {
															v189 = int32(1024)
														} else {
															v189 = v125 << (uint(int32(9)) % 32)
														}
														if v125 == int32(1) {
															v192 = int32(864026624)
														} else {
															v192 = int32(872415232)
														}
														v193 = v189
														v195 = v192
													}
												}
											}
										}
									}
								}
							}
						}
						v205 = v193 & int32(1022)
						v206 = v195 | v129
					} else {
						v205 = int32(0)
						v206 = v129
					}
				}
			} else {
				if v125 == int32(0) {
					v205 = int32(0)
					v206 = v129 | int32(2139095040)
				} else {
					v205 = v125
					v206 = v129 | int32(2143289344)
				}
			}
			v210 = base.F32_reinterpret_i32(v206 | v205<<(uint(int32(13))%32))
			v212 = base.F32_add(base.F32_mul(v118, v210), v22)
			v214 = base.F32_add(base.F32_mul(v118, v118), v23)
			v216 = base.F32_add(base.F32_mul(v210, v210), v24)
			v218 = v18 + int32(1)
			if v218 != l0 {
				v18 = v218
				v22 = v212
				v23 = v214
				v24 = v216
				continue
			} else {
				break
			}
			break
		}
		return base.F64_div(base.F64_promote_f32(v212), base.F64_sqrt(base.F64_mul(base.F64_promote_f32(v214), base.F64_promote_f32(v216))))
	}
}
func F_HalfvecL2SquaredDistanceDefault(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 float32
	_ = v6
	var v13 int32
	_ = v13
	var v15 float32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 float32
	_ = v203
	var v205 float32
	_ = v205
	var v207 int32
	_ = v207
	var v214 float32
	_ = v214
	v4 = int32(0)
	v6 = float32(0)
	if v4 < l0 {
		v13 = v4
		v15 = v6
		for {
			v18 = v13 << (uint(int32(1)) % 32)
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v18))))
			v25 = v20 & int32(1023)
			v29 = v20 << (uint(int32(16)) % 32) & int32(-2147483648)
			v32 = int32(31)
			v33 = int32(base.Ui32(v20)>>(uint(int32(10))%32)) & v32
			if v33 != v32 {
				if v33 != 0 {
					v105 = v25
					v106 = v33<<(uint(int32(23))%32) + v29 + int32(939524096)
				} else {
					if v25 != 0 {
						if v20&int32(512) != 0 {
							v93 = v25 << (uint(int32(1)) % 32)
							v95 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v25) {
								v93 = v25 << (uint(int32(2)) % 32)
								v95 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v25) {
									v93 = v25 << (uint(int32(3)) % 32)
									v95 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v25) {
										v93 = v25 << (uint(int32(4)) % 32)
										v95 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v25) {
											v93 = v25 << (uint(int32(5)) % 32)
											v95 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v25) {
												v93 = v25 << (uint(int32(6)) % 32)
												v95 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v25) {
													v93 = v25 << (uint(int32(7)) % 32)
													v95 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v25) {
														v93 = v25 << (uint(int32(8)) % 32)
														v95 = int32(880803840)
													} else {
														v88 = base.B2i32(v25 == int32(1))
														if v25 == int32(1) {
															v89 = int32(1024)
														} else {
															v89 = v25 << (uint(int32(9)) % 32)
														}
														if v25 == int32(1) {
															v92 = int32(864026624)
														} else {
															v92 = int32(872415232)
														}
														v93 = v89
														v95 = v92
													}
												}
											}
										}
									}
								}
							}
						}
						v105 = v93 & int32(1022)
						v106 = v95 | v29
					} else {
						v105 = int32(0)
						v106 = v29
					}
				}
			} else {
				if v25 == int32(0) {
					v105 = int32(0)
					v106 = v29 | int32(2139095040)
				} else {
					v105 = v25
					v106 = v29 | int32(2143289344)
				}
			}
			v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v18))))
			v117 = v112 & int32(1023)
			v121 = v112 << (uint(int32(16)) % 32) & int32(-2147483648)
			v124 = int32(31)
			v125 = int32(base.Ui32(v112)>>(uint(int32(10))%32)) & v124
			if v125 != v124 {
				if v125 != 0 {
					v197 = v117
					v198 = v125<<(uint(int32(23))%32) + v121 + int32(939524096)
				} else {
					if v117 != 0 {
						if v112&int32(512) != 0 {
							v185 = v117 << (uint(int32(1)) % 32)
							v187 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v117) {
								v185 = v117 << (uint(int32(2)) % 32)
								v187 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v117) {
									v185 = v117 << (uint(int32(3)) % 32)
									v187 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v117) {
										v185 = v117 << (uint(int32(4)) % 32)
										v187 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v117) {
											v185 = v117 << (uint(int32(5)) % 32)
											v187 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v117) {
												v185 = v117 << (uint(int32(6)) % 32)
												v187 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v117) {
													v185 = v117 << (uint(int32(7)) % 32)
													v187 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v117) {
														v185 = v117 << (uint(int32(8)) % 32)
														v187 = int32(880803840)
													} else {
														v180 = base.B2i32(v117 == int32(1))
														if v117 == int32(1) {
															v181 = int32(1024)
														} else {
															v181 = v117 << (uint(int32(9)) % 32)
														}
														if v117 == int32(1) {
															v184 = int32(864026624)
														} else {
															v184 = int32(872415232)
														}
														v185 = v181
														v187 = v184
													}
												}
											}
										}
									}
								}
							}
						}
						v197 = v185 & int32(1022)
						v198 = v187 | v121
					} else {
						v197 = int32(0)
						v198 = v121
					}
				}
			} else {
				if v117 == int32(0) {
					v197 = int32(0)
					v198 = v121 | int32(2139095040)
				} else {
					v197 = v117
					v198 = v121 | int32(2143289344)
				}
			}
			v203 = base.F32_sub(base.F32_reinterpret_i32(v106|v105<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v198|v197<<(uint(int32(13))%32)))
			v205 = base.F32_add(base.F32_mul(v203, v203), v15)
			v207 = v13 + int32(1)
			if v207 != l0 {
				v13 = v207
				v15 = v205
				continue
			} else {
				break
			}
			break
		}
		v214 = v205
	} else {
		v214 = v6
	}
	return v214
}
func F_halfvec_avg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 float64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L22
	}
L2:
	;
	return int32(0)
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v18 != int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v25 != int32(701) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	if base.F64_eq(v28, float64(0)) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v11 + int32(16)
	return v88
L9:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v88 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v33 = int32(_a_F_halfvec_avg_0)
	v34 = v21 + v33
	v36 = v34 & v33
	F_CheckDim_1(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v41 = F_mul_size(m, int32(2), v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v43 = F_add_size(m, int32(8), v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v45 = F_palloc0(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v43 << (uint(int32(2)) % 32)
	if v36 == int32(0) {
		v88 = v45
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v58 = int32(0)
	goto L17
L17:
	;
	v66 = int32(1)
	v70 = v58 + v66
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v14+int32(24)+v70<<(uint(int32(3))%32))))
	v77 = F_Float4ToHalf(m, base.F32_demote_f64(base.F64_div(v74, v28)))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v88 = v45
	goto L8
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(8)+v58<<(uint(v66)%32)))) = uint16(v77)
	F_CheckElement_1(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	if v70 != v36 {
		v58 = v70
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_halfvec_avg_1)
	F_errmsg_internal(m, int32(_a_F_halfvec_avg_2), v11)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_halfvec_avg_3), int32(173), int32(_a_F_halfvec_avg_4))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 float32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 float32
	_ = v225
	var v233 int32
	_ = v233
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = base.B2i32(v19 < v20)
	if v19 < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v19 < v20 {
		goto L104
	} else {
		goto L105
	}
L5:
	;
	v22 = v19
	goto L7
L6:
	;
	v22 = v20
	goto L7
L7:
	;
	if v22 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = v30 << (uint(int32(1)) % 32)
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v25+v41))))
	v48 = v43 & int32(1023)
	v52 = v43 << (uint(int32(16)) % 32) & int32(-2147483648)
	v55 = int32(31)
	v56 = int32(base.Ui32(v43)>>(uint(int32(10))%32)) & v55
	if v56 != v55 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(1)
L11:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v25+v41))))
	v140 = v135 & int32(1023)
	v144 = v135 << (uint(int32(16)) % 32) & int32(-2147483648)
	v147 = int32(31)
	v148 = int32(base.Ui32(v135)>>(uint(int32(10))%32)) & v147
	if v148 != v147 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v133 = base.F32_reinterpret_i32(v129 | v128<<(uint(int32(13))%32))
	goto L11
L13:
	;
	v128 = v48
	v129 = v56<<(uint(int32(23))%32) + v52 + int32(939524096)
	goto L12
L14:
	;
	if v43&int32(512) != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v56 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v48 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v128 = int32(0)
	v129 = v52
	goto L12
L20:
	;
	v128 = int32(0)
	v129 = v52 | int32(2139095040)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v128 = v48
	v129 = v52 | int32(2143289344)
	goto L12
L23:
	;
	v128 = v116 & int32(1022)
	v129 = v118 | v52
	goto L12
L24:
	;
	v116 = v48 << (uint(int32(1)) % 32)
	v118 = int32(939524096)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v48) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = v48 << (uint(int32(2)) % 32)
	v118 = int32(931135488)
	goto L23
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(127)) < base.Ui32(v48) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v116 = v48 << (uint(int32(3)) % 32)
	v118 = int32(922746880)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(63)) < base.Ui32(v48) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = v48 << (uint(int32(4)) % 32)
	v118 = int32(914358272)
	goto L23
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(31)) < base.Ui32(v48) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v116 = v48 << (uint(int32(5)) % 32)
	v118 = int32(905969664)
	goto L23
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(15)) < base.Ui32(v48) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v116 = v48 << (uint(int32(6)) % 32)
	v118 = int32(897581056)
	goto L23
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(7)) < base.Ui32(v48) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v116 = v48 << (uint(int32(7)) % 32)
	v118 = int32(889192448)
	goto L23
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(3)) < base.Ui32(v48) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v116 = v48 << (uint(int32(8)) % 32)
	v118 = int32(880803840)
	goto L23
L46:
	;
	goto L47
L47:
	;
	v111 = base.B2i32(v48 == int32(1))
	if v48 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v112 = int32(1024)
	goto L50
L49:
	;
	v112 = v48 << (uint(int32(9)) % 32)
	goto L50
L50:
	;
	if v48 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v115 = int32(864026624)
	goto L53
L52:
	;
	v115 = int32(872415232)
	goto L53
L53:
	;
	v116 = v112
	v118 = v115
	goto L23
L54:
	;
	if base.F32_lt(v133, v225) != 0 {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v225 = base.F32_reinterpret_i32(v221 | v220<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v220 = v140
	v221 = v148<<(uint(int32(23))%32) + v144 + int32(939524096)
	goto L55
L57:
	;
	if v135&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v148 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v140 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v140 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v220 = int32(0)
	v221 = v144
	goto L55
L63:
	;
	v220 = int32(0)
	v221 = v144 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v220 = v140
	v221 = v144 | int32(2143289344)
	goto L55
L66:
	;
	v220 = v208 & int32(1022)
	v221 = v210 | v144
	goto L55
L67:
	;
	v208 = v140 << (uint(int32(1)) % 32)
	v210 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v140) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v208 = v140 << (uint(int32(2)) % 32)
	v210 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v140) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v208 = v140 << (uint(int32(3)) % 32)
	v210 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v140) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v208 = v140 << (uint(int32(4)) % 32)
	v210 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v140) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v208 = v140 << (uint(int32(5)) % 32)
	v210 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v140) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v208 = v140 << (uint(int32(6)) % 32)
	v210 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v140) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v208 = v140 << (uint(int32(7)) % 32)
	v210 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v208 = v140 << (uint(int32(8)) % 32)
	v210 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v203 = base.B2i32(v140 == int32(1))
	if v140 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v204 = int32(1024)
	goto L93
L92:
	;
	v204 = v140 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v140 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v207 = int32(864026624)
	goto L96
L95:
	;
	v207 = int32(872415232)
	goto L96
L96:
	;
	v208 = v204
	v210 = v207
	goto L66
L97:
	;
	return int32(-1)
L98:
	;
	goto L99
L99:
	;
	if base.F32_gt(v133, v225) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v233 = v30 + int32(1)
	if v233 == v22 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L10
L103:
	;
	v30 = v233
	goto L9
L104:
	;
	return int32(-1)
L105:
	;
	goto L106
L106:
	;
	return base.B2i32(v20 < v19)
}
func F_halfvec_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 float32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 float32
	_ = v227
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = base.B2i32(v20 < v21)
	if v20 < v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v254
L5:
	;
	v23 = v20
	goto L7
L6:
	;
	v23 = v21
	goto L7
L7:
	;
	if int32(0) < v23 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v20 < v21 {
		v254 = v2
		goto L4
	} else {
		goto L101
	}
L11:
	;
	v43 = v31 << (uint(int32(1)) % 32)
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v26+v43))))
	v50 = v45 & int32(1023)
	v54 = v45 << (uint(int32(16)) % 32) & int32(-2147483648)
	v57 = int32(31)
	v58 = int32(base.Ui32(v45)>>(uint(int32(10))%32)) & v57
	if v58 != v57 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	goto L10
L13:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+v26+v43))))
	v142 = v137 & int32(1023)
	v146 = v137 << (uint(int32(16)) % 32) & int32(-2147483648)
	v149 = int32(31)
	v150 = int32(base.Ui32(v137)>>(uint(int32(10))%32)) & v149
	if v150 != v149 {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	v135 = base.F32_reinterpret_i32(v131 | v130<<(uint(int32(13))%32))
	goto L13
L15:
	;
	v130 = v50
	v131 = v58<<(uint(int32(23))%32) + v54 + int32(939524096)
	goto L14
L16:
	;
	if v45&int32(512) != 0 {
		goto L26
	} else {
		goto L27
	}
L17:
	;
	if v58 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v50 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v130 = int32(0)
	v131 = v54
	goto L14
L22:
	;
	v130 = int32(0)
	v131 = v54 | int32(2139095040)
	goto L14
L23:
	;
	goto L24
L24:
	;
	v130 = v50
	v131 = v54 | int32(2143289344)
	goto L14
L25:
	;
	v130 = v118 & int32(1022)
	v131 = v120 | v54
	goto L14
L26:
	;
	v118 = v50 << (uint(int32(1)) % 32)
	v120 = int32(939524096)
	goto L25
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(int32(255)) < base.Ui32(v50) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v118 = v50 << (uint(int32(2)) % 32)
	v120 = int32(931135488)
	goto L25
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(127)) < base.Ui32(v50) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = v50 << (uint(int32(3)) % 32)
	v120 = int32(922746880)
	goto L25
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(int32(63)) < base.Ui32(v50) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v118 = v50 << (uint(int32(4)) % 32)
	v120 = int32(914358272)
	goto L25
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(int32(31)) < base.Ui32(v50) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = v50 << (uint(int32(5)) % 32)
	v120 = int32(905969664)
	goto L25
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(int32(15)) < base.Ui32(v50) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v118 = v50 << (uint(int32(6)) % 32)
	v120 = int32(897581056)
	goto L25
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(7)) < base.Ui32(v50) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v118 = v50 << (uint(int32(7)) % 32)
	v120 = int32(889192448)
	goto L25
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(3)) < base.Ui32(v50) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v118 = v50 << (uint(int32(8)) % 32)
	v120 = int32(880803840)
	goto L25
L48:
	;
	goto L49
L49:
	;
	v113 = base.B2i32(v50 == int32(1))
	if v50 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v114 = int32(1024)
	goto L52
L51:
	;
	v114 = v50 << (uint(int32(9)) % 32)
	goto L52
L52:
	;
	if v50 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v117 = int32(864026624)
	goto L55
L54:
	;
	v117 = int32(872415232)
	goto L55
L55:
	;
	v118 = v114
	v120 = v117
	goto L25
L56:
	;
	if base.F32_gt(v135, v227)|base.F32_lt(v135, v227) != 0 {
		v254 = v2
		goto L4
	} else {
		goto L99
	}
L57:
	;
	v227 = base.F32_reinterpret_i32(v223 | v222<<(uint(int32(13))%32))
	goto L56
L58:
	;
	v222 = v142
	v223 = v150<<(uint(int32(23))%32) + v146 + int32(939524096)
	goto L57
L59:
	;
	if v137&int32(512) != 0 {
		goto L69
	} else {
		goto L70
	}
L60:
	;
	if v150 != 0 {
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v142 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if v142 != 0 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v222 = int32(0)
	v223 = v146
	goto L57
L65:
	;
	v222 = int32(0)
	v223 = v146 | int32(2139095040)
	goto L57
L66:
	;
	goto L67
L67:
	;
	v222 = v142
	v223 = v146 | int32(2143289344)
	goto L57
L68:
	;
	v222 = v210 & int32(1022)
	v223 = v212 | v146
	goto L57
L69:
	;
	v210 = v142 << (uint(int32(1)) % 32)
	v212 = int32(939524096)
	goto L68
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(int32(255)) < base.Ui32(v142) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v210 = v142 << (uint(int32(2)) % 32)
	v212 = int32(931135488)
	goto L68
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(127)) < base.Ui32(v142) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v210 = v142 << (uint(int32(3)) % 32)
	v212 = int32(922746880)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if base.Ui32(int32(63)) < base.Ui32(v142) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v210 = v142 << (uint(int32(4)) % 32)
	v212 = int32(914358272)
	goto L68
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(int32(31)) < base.Ui32(v142) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v210 = v142 << (uint(int32(5)) % 32)
	v212 = int32(905969664)
	goto L68
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(int32(15)) < base.Ui32(v142) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v210 = v142 << (uint(int32(6)) % 32)
	v212 = int32(897581056)
	goto L68
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(int32(7)) < base.Ui32(v142) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v210 = v142 << (uint(int32(7)) % 32)
	v212 = int32(889192448)
	goto L68
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v210 = v142 << (uint(int32(8)) % 32)
	v212 = int32(880803840)
	goto L68
L91:
	;
	goto L92
L92:
	;
	v205 = base.B2i32(v142 == int32(1))
	if v142 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v206 = int32(1024)
	goto L95
L94:
	;
	v206 = v142 << (uint(int32(9)) % 32)
	goto L95
L95:
	;
	if v142 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v209 = int32(864026624)
	goto L98
L97:
	;
	v209 = int32(872415232)
	goto L98
L98:
	;
	v210 = v206
	v212 = v209
	goto L68
L99:
	;
	v232 = v31 + int32(1)
	if v232 != v23 {
		v31 = v232
		goto L11
	} else {
		goto L100
	}
L100:
	;
	goto L12
L101:
	;
	v254 = base.B2i32(v20 <= v21)
	goto L4
}
func F_halfvec_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 float32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 float32
	_ = v225
	var v231 int32
	_ = v231
	var v253 int32
	_ = v253
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v19 < v20 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v253
L5:
	;
	v253 = base.B2i32(v20 <= v19)
	goto L4
L6:
	;
	v22 = v19
	goto L8
L7:
	;
	v22 = v20
	goto L8
L8:
	;
	if v22 <= int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L10
L10:
	;
	v41 = v30 << (uint(int32(1)) % 32)
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v25+v41))))
	v48 = v43 & int32(1023)
	v52 = v43 << (uint(int32(16)) % 32) & int32(-2147483648)
	v55 = int32(31)
	v56 = int32(base.Ui32(v43)>>(uint(int32(10))%32)) & v55
	if v56 != v55 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return int32(1)
L12:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v25+v41))))
	v140 = v135 & int32(1023)
	v144 = v135 << (uint(int32(16)) % 32) & int32(-2147483648)
	v147 = int32(31)
	v148 = int32(base.Ui32(v135)>>(uint(int32(10))%32)) & v147
	if v148 != v147 {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	v133 = base.F32_reinterpret_i32(v129 | v128<<(uint(int32(13))%32))
	goto L12
L14:
	;
	v128 = v48
	v129 = v56<<(uint(int32(23))%32) + v52 + int32(939524096)
	goto L13
L15:
	;
	if v43&int32(512) != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	if v56 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v48 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v128 = int32(0)
	v129 = v52
	goto L13
L21:
	;
	v128 = int32(0)
	v129 = v52 | int32(2139095040)
	goto L13
L22:
	;
	goto L23
L23:
	;
	v128 = v48
	v129 = v52 | int32(2143289344)
	goto L13
L24:
	;
	v128 = v116 & int32(1022)
	v129 = v118 | v52
	goto L13
L25:
	;
	v116 = v48 << (uint(int32(1)) % 32)
	v118 = int32(939524096)
	goto L24
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(int32(255)) < base.Ui32(v48) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = v48 << (uint(int32(2)) % 32)
	v118 = int32(931135488)
	goto L24
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(int32(127)) < base.Ui32(v48) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = v48 << (uint(int32(3)) % 32)
	v118 = int32(922746880)
	goto L24
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(int32(63)) < base.Ui32(v48) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = v48 << (uint(int32(4)) % 32)
	v118 = int32(914358272)
	goto L24
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(int32(31)) < base.Ui32(v48) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v116 = v48 << (uint(int32(5)) % 32)
	v118 = int32(905969664)
	goto L24
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(int32(15)) < base.Ui32(v48) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v116 = v48 << (uint(int32(6)) % 32)
	v118 = int32(897581056)
	goto L24
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(7)) < base.Ui32(v48) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v116 = v48 << (uint(int32(7)) % 32)
	v118 = int32(889192448)
	goto L24
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(3)) < base.Ui32(v48) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v116 = v48 << (uint(int32(8)) % 32)
	v118 = int32(880803840)
	goto L24
L47:
	;
	goto L48
L48:
	;
	v111 = base.B2i32(v48 == int32(1))
	if v48 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v112 = int32(1024)
	goto L51
L50:
	;
	v112 = v48 << (uint(int32(9)) % 32)
	goto L51
L51:
	;
	if v48 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v115 = int32(864026624)
	goto L54
L53:
	;
	v115 = int32(872415232)
	goto L54
L54:
	;
	v116 = v112
	v118 = v115
	goto L24
L55:
	;
	if base.F32_lt(v133, v225) != 0 {
		v253 = int32(0)
		goto L4
	} else {
		goto L98
	}
L56:
	;
	v225 = base.F32_reinterpret_i32(v221 | v220<<(uint(int32(13))%32))
	goto L55
L57:
	;
	v220 = v140
	v221 = v148<<(uint(int32(23))%32) + v144 + int32(939524096)
	goto L56
L58:
	;
	if v135&int32(512) != 0 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	if v148 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v140 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v140 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v220 = int32(0)
	v221 = v144
	goto L56
L64:
	;
	v220 = int32(0)
	v221 = v144 | int32(2139095040)
	goto L56
L65:
	;
	goto L66
L66:
	;
	v220 = v140
	v221 = v144 | int32(2143289344)
	goto L56
L67:
	;
	v220 = v208 & int32(1022)
	v221 = v210 | v144
	goto L56
L68:
	;
	v208 = v140 << (uint(int32(1)) % 32)
	v210 = int32(939524096)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(int32(255)) < base.Ui32(v140) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v208 = v140 << (uint(int32(2)) % 32)
	v210 = int32(931135488)
	goto L67
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(127)) < base.Ui32(v140) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v208 = v140 << (uint(int32(3)) % 32)
	v210 = int32(922746880)
	goto L67
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(63)) < base.Ui32(v140) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v208 = v140 << (uint(int32(4)) % 32)
	v210 = int32(914358272)
	goto L67
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(31)) < base.Ui32(v140) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v208 = v140 << (uint(int32(5)) % 32)
	v210 = int32(905969664)
	goto L67
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v140) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v208 = v140 << (uint(int32(6)) % 32)
	v210 = int32(897581056)
	goto L67
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(7)) < base.Ui32(v140) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v208 = v140 << (uint(int32(7)) % 32)
	v210 = int32(889192448)
	goto L67
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v208 = v140 << (uint(int32(8)) % 32)
	v210 = int32(880803840)
	goto L67
L90:
	;
	goto L91
L91:
	;
	v203 = base.B2i32(v140 == int32(1))
	if v140 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v204 = int32(1024)
	goto L94
L93:
	;
	v204 = v140 << (uint(int32(9)) % 32)
	goto L94
L94:
	;
	if v140 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v207 = int32(864026624)
	goto L97
L96:
	;
	v207 = int32(872415232)
	goto L97
L97:
	;
	v208 = v204
	v210 = v207
	goto L67
L98:
	;
	if base.F32_gt(v133, v225) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v231 = v30 + int32(1)
	if v231 == v22 {
		goto L5
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L11
L102:
	;
	v30 = v231
	goto L10
}
func F_halfvec_l2_normalize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 float64
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 float64
	_ = v147
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 float32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 float32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v22 = F_mul_size(m, int32(2), v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_add_size(m, int32(8), v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = F_palloc0(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)) = uint16(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v24 << (uint(int32(2)) % 32)
	v32 = int32(0)
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v33 <= v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return v26
L7:
	;
	v36 = int32(8)
	v37 = v17 + v36
	v39 = v26 + v36
	v40 = v32
	v50 = float64(0)
	goto L8
L8:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+v40<<(uint(int32(1))%32)))))
	v61 = v56 & int32(1023)
	v65 = v56 << (uint(int32(16)) % 32) & int32(-2147483648)
	v68 = int32(31)
	v69 = int32(base.Ui32(v56)>>(uint(int32(10))%32)) & v68
	if v69 != v68 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	if base.F64_gt(v149, float64(0)) == int32(0) {
		goto L6
	} else {
		goto L54
	}
L10:
	;
	v147 = base.F64_promote_f32(base.F32_reinterpret_i32(v142 | v141<<(uint(int32(13))%32)))
	v149 = base.F64_add(base.F64_mul(v147, v147), v50)
	v151 = v40 + int32(1)
	if v151 != v33 {
		v40 = v151
		v50 = v149
		goto L8
	} else {
		goto L53
	}
L11:
	;
	goto L10
L12:
	;
	v141 = v61
	v142 = v69<<(uint(int32(23))%32) + v65 + int32(939524096)
	goto L11
L13:
	;
	if v56&int32(512) != 0 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	if v69 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v61 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v141 = int32(0)
	v142 = v65
	goto L11
L19:
	;
	v141 = int32(0)
	v142 = v65 | int32(2139095040)
	goto L11
L20:
	;
	goto L21
L21:
	;
	v141 = v61
	v142 = v65 | int32(2143289344)
	goto L11
L22:
	;
	v141 = v129 & int32(1022)
	v142 = v131 | v65
	goto L11
L23:
	;
	v129 = v61 << (uint(int32(1)) % 32)
	v131 = int32(939524096)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(int32(255)) < base.Ui32(v61) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v129 = v61 << (uint(int32(2)) % 32)
	v131 = int32(931135488)
	goto L22
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(int32(127)) < base.Ui32(v61) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v129 = v61 << (uint(int32(3)) % 32)
	v131 = int32(922746880)
	goto L22
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(63)) < base.Ui32(v61) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v129 = v61 << (uint(int32(4)) % 32)
	v131 = int32(914358272)
	goto L22
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(int32(31)) < base.Ui32(v61) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v129 = v61 << (uint(int32(5)) % 32)
	v131 = int32(905969664)
	goto L22
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(int32(15)) < base.Ui32(v61) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v129 = v61 << (uint(int32(6)) % 32)
	v131 = int32(897581056)
	goto L22
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(int32(7)) < base.Ui32(v61) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v129 = v61 << (uint(int32(7)) % 32)
	v131 = int32(889192448)
	goto L22
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v129 = v61 << (uint(int32(8)) % 32)
	v131 = int32(880803840)
	goto L22
L45:
	;
	goto L46
L46:
	;
	v124 = base.B2i32(v61 == int32(1))
	if v61 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v125 = int32(1024)
	goto L49
L48:
	;
	v125 = v61 << (uint(int32(9)) % 32)
	goto L49
L49:
	;
	if v61 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v128 = int32(864026624)
	goto L52
L51:
	;
	v128 = int32(872415232)
	goto L52
L52:
	;
	v129 = v125
	v131 = v128
	goto L22
L53:
	;
	goto L9
L54:
	;
	v159 = int32(0)
	goto L55
L55:
	;
	v173 = v159 << (uint(int32(1)) % 32)
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+v173))))
	v180 = v175 & int32(1023)
	v184 = v175 << (uint(int32(16)) % 32) & int32(-2147483648)
	v187 = int32(31)
	v188 = int32(base.Ui32(v175)>>(uint(int32(10))%32)) & v187
	if v188 != v187 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v360 = int32(0)
	if v358 <= v360 {
		goto L6
	} else {
		goto L121
	}
L57:
	;
	v268 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(base.F32_reinterpret_i32(v261|v260<<(uint(int32(13))%32))), base.F64_sqrt(v149)))
	v269 = base.I32_reinterpret_f32(v268)
	v271 = int32(base.Ui32(v269) >> (uint(int32(16)) % 32))
	v275 = base.F32_abs(v268)
	if base.F32_eq(v275, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v354 = v271 & int32(_a_F_halfvec_l2_normalize_0)
		goto L100
	} else {
		goto L101
	}
L58:
	;
	goto L57
L59:
	;
	v260 = v180
	v261 = v188<<(uint(int32(23))%32) + v184 + int32(939524096)
	goto L58
L60:
	;
	if v175&int32(512) != 0 {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	if v188 != 0 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v180 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v180 != 0 {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v260 = int32(0)
	v261 = v184
	goto L58
L66:
	;
	v260 = int32(0)
	v261 = v184 | int32(2139095040)
	goto L58
L67:
	;
	goto L68
L68:
	;
	v260 = v180
	v261 = v184 | int32(2143289344)
	goto L58
L69:
	;
	v260 = v248 & int32(1022)
	v261 = v250 | v184
	goto L58
L70:
	;
	v248 = v180 << (uint(int32(1)) % 32)
	v250 = int32(939524096)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(255)) < base.Ui32(v180) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v248 = v180 << (uint(int32(2)) % 32)
	v250 = int32(931135488)
	goto L69
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(127)) < base.Ui32(v180) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v248 = v180 << (uint(int32(3)) % 32)
	v250 = int32(922746880)
	goto L69
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(63)) < base.Ui32(v180) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v248 = v180 << (uint(int32(4)) % 32)
	v250 = int32(914358272)
	goto L69
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(31)) < base.Ui32(v180) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v248 = v180 << (uint(int32(5)) % 32)
	v250 = int32(905969664)
	goto L69
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(15)) < base.Ui32(v180) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v248 = v180 << (uint(int32(6)) % 32)
	v250 = int32(897581056)
	goto L69
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(7)) < base.Ui32(v180) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v248 = v180 << (uint(int32(7)) % 32)
	v250 = int32(889192448)
	goto L69
L89:
	;
	goto L90
L90:
	;
	if base.Ui32(int32(3)) < base.Ui32(v180) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v248 = v180 << (uint(int32(8)) % 32)
	v250 = int32(880803840)
	goto L69
L92:
	;
	goto L93
L93:
	;
	v243 = base.B2i32(v180 == int32(1))
	if v180 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v244 = int32(1024)
	goto L96
L95:
	;
	v244 = v180 << (uint(int32(9)) % 32)
	goto L96
L96:
	;
	if v180 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v247 = int32(864026624)
	goto L99
L98:
	;
	v247 = int32(872415232)
	goto L99
L99:
	;
	v248 = v244
	v250 = v247
	goto L69
L100:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v173+v39))) = uint16(v354)
	v357 = v159 + int32(1)
	v358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v357 < v358 {
		v159 = v357
		goto L55
	} else {
		goto L120
	}
L101:
	;
	v279 = v271 & int32(_a_F_halfvec_l2_normalize_1)
	v281 = v269 & int32(_a_F_halfvec_l2_normalize_2)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v275)) {
		v354 = v279 | int32(base.Ui32(v281)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_l2_normalize_3)
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v293 = int32(base.Ui32(v269)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v293) < base.Ui32(int32(99)) {
		v354 = v279
		goto L100
	} else {
		goto L103
	}
L103:
	;
	if base.Ui32(v293) <= base.Ui32(int32(112)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v305 = int32(1)<<(uint(v293-int32(90))%32) + int32(base.Ui32(v281)>>(uint(int32(113)-v293)%32))
	v307 = v305 | v269
	v308 = v305
	goto L106
L105:
	;
	v307 = v269
	v308 = v281
	goto L106
L106:
	;
	v310 = int32(base.Ui32(v308) >> (uint(int32(13)) % 32))
	v315 = int32(1)
	v319 = int32(3)
	v320 = int32(base.Ui32(v308)>>(uint(int32(12))%32)) & v319
	if base.B2i32(v320 != v319)&(base.B2i32(v307&int32(4095) == int32(0))|base.B2i32(v320 != v315)) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v331 = v310
	goto L109
L108:
	;
	v331 = v310 + v315
	goto L109
L109:
	;
	v333 = base.B2i32(v331 == int32(1024))
	if v331 == int32(1024) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v334 = int32(-126)
	goto L112
L111:
	;
	v334 = int32(-127)
	goto L112
L112:
	;
	v335 = v334 + v293
	if int32(16) <= v335 {
		v354 = v279 | int32(_a_F_halfvec_l2_normalize_4)
		goto L100
	} else {
		goto L113
	}
L113:
	;
	if int32(-15) < v335 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v345 = v335<<(uint(int32(10))%32) + int32(_a_F_halfvec_l2_normalize_5) | v279
	goto L116
L115:
	;
	v345 = v279
	goto L116
L116:
	;
	if v331 == int32(1024) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v347 = int32(0)
	goto L119
L118:
	;
	v347 = v331
	goto L119
L119:
	;
	v354 = v345 | v347
	goto L100
L120:
	;
	goto L56
L121:
	;
	v363 = v360
	goto L122
L122:
	;
	v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v363<<(uint(int32(1))%32)))))
	if v379&int32(_a_F_halfvec_l2_normalize_6) != int32(_a_F_halfvec_l2_normalize_4) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L128
	}
L124:
	;
	v385 = v363 + int32(1)
	if v358 != v385 {
		v363 = v385
		goto L122
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	goto L123
L127:
	;
	goto L6
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_l2_squared_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13913(m, l0, int32(_a_F_halfvec_l2_squared_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_halfvec_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 float32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 float32
	_ = v224
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	v20 = base.B2i32(v18 < v19)
	if v18 < v19 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v244
L5:
	;
	v244 = v20
	goto L4
L6:
	;
	v21 = v18
	goto L8
L7:
	;
	v21 = v19
	goto L8
L8:
	;
	if v21 <= int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v24 = int32(8)
	v29 = int32(0)
	goto L10
L10:
	;
	v38 = int32(1)
	v40 = v29 << (uint(v38) % 32)
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+v24+v40))))
	v47 = v42 & int32(1023)
	v51 = v42 << (uint(int32(16)) % 32) & int32(-2147483648)
	v54 = int32(31)
	v55 = int32(base.Ui32(v42)>>(uint(int32(10))%32)) & v54
	if v55 != v54 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return int32(0)
L12:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+v24+v40))))
	v139 = v134 & int32(1023)
	v143 = v134 << (uint(int32(16)) % 32) & int32(-2147483648)
	v146 = int32(31)
	v147 = int32(base.Ui32(v134)>>(uint(int32(10))%32)) & v146
	if v147 != v146 {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	v132 = base.F32_reinterpret_i32(v128 | v127<<(uint(int32(13))%32))
	goto L12
L14:
	;
	v127 = v47
	v128 = v55<<(uint(int32(23))%32) + v51 + int32(939524096)
	goto L13
L15:
	;
	if v42&int32(512) != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	if v55 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v47 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v47 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v127 = int32(0)
	v128 = v51
	goto L13
L21:
	;
	v127 = int32(0)
	v128 = v51 | int32(2139095040)
	goto L13
L22:
	;
	goto L23
L23:
	;
	v127 = v47
	v128 = v51 | int32(2143289344)
	goto L13
L24:
	;
	v127 = v115 & int32(1022)
	v128 = v117 | v51
	goto L13
L25:
	;
	v115 = v47 << (uint(int32(1)) % 32)
	v117 = int32(939524096)
	goto L24
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(int32(255)) < base.Ui32(v47) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v115 = v47 << (uint(int32(2)) % 32)
	v117 = int32(931135488)
	goto L24
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(int32(127)) < base.Ui32(v47) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = v47 << (uint(int32(3)) % 32)
	v117 = int32(922746880)
	goto L24
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(int32(63)) < base.Ui32(v47) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v115 = v47 << (uint(int32(4)) % 32)
	v117 = int32(914358272)
	goto L24
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(int32(31)) < base.Ui32(v47) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v115 = v47 << (uint(int32(5)) % 32)
	v117 = int32(905969664)
	goto L24
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(int32(15)) < base.Ui32(v47) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = v47 << (uint(int32(6)) % 32)
	v117 = int32(897581056)
	goto L24
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(7)) < base.Ui32(v47) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v115 = v47 << (uint(int32(7)) % 32)
	v117 = int32(889192448)
	goto L24
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(3)) < base.Ui32(v47) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v115 = v47 << (uint(int32(8)) % 32)
	v117 = int32(880803840)
	goto L24
L47:
	;
	goto L48
L48:
	;
	v110 = base.B2i32(v47 == int32(1))
	if v47 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v111 = int32(1024)
	goto L51
L50:
	;
	v111 = v47 << (uint(int32(9)) % 32)
	goto L51
L51:
	;
	if v47 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v114 = int32(864026624)
	goto L54
L53:
	;
	v114 = int32(872415232)
	goto L54
L54:
	;
	v115 = v111
	v117 = v114
	goto L24
L55:
	;
	if base.F32_lt(v132, v224) != 0 {
		v244 = v38
		goto L4
	} else {
		goto L98
	}
L56:
	;
	v224 = base.F32_reinterpret_i32(v220 | v219<<(uint(int32(13))%32))
	goto L55
L57:
	;
	v219 = v139
	v220 = v147<<(uint(int32(23))%32) + v143 + int32(939524096)
	goto L56
L58:
	;
	if v134&int32(512) != 0 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	if v147 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v139 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v139 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v219 = int32(0)
	v220 = v143
	goto L56
L64:
	;
	v219 = int32(0)
	v220 = v143 | int32(2139095040)
	goto L56
L65:
	;
	goto L66
L66:
	;
	v219 = v139
	v220 = v143 | int32(2143289344)
	goto L56
L67:
	;
	v219 = v207 & int32(1022)
	v220 = v209 | v143
	goto L56
L68:
	;
	v207 = v139 << (uint(int32(1)) % 32)
	v209 = int32(939524096)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(int32(255)) < base.Ui32(v139) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v207 = v139 << (uint(int32(2)) % 32)
	v209 = int32(931135488)
	goto L67
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(127)) < base.Ui32(v139) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v207 = v139 << (uint(int32(3)) % 32)
	v209 = int32(922746880)
	goto L67
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(63)) < base.Ui32(v139) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v207 = v139 << (uint(int32(4)) % 32)
	v209 = int32(914358272)
	goto L67
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(31)) < base.Ui32(v139) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v207 = v139 << (uint(int32(5)) % 32)
	v209 = int32(905969664)
	goto L67
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v139) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v207 = v139 << (uint(int32(6)) % 32)
	v209 = int32(897581056)
	goto L67
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(7)) < base.Ui32(v139) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v207 = v139 << (uint(int32(7)) % 32)
	v209 = int32(889192448)
	goto L67
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(3)) < base.Ui32(v139) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v207 = v139 << (uint(int32(8)) % 32)
	v209 = int32(880803840)
	goto L67
L90:
	;
	goto L91
L91:
	;
	v202 = base.B2i32(v139 == int32(1))
	if v139 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v203 = int32(1024)
	goto L94
L93:
	;
	v203 = v139 << (uint(int32(9)) % 32)
	goto L94
L94:
	;
	if v139 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v206 = int32(864026624)
	goto L97
L96:
	;
	v206 = int32(872415232)
	goto L97
L97:
	;
	v207 = v203
	v209 = v206
	goto L67
L98:
	;
	if base.F32_gt(v132, v224) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v230 = v29 + int32(1)
	if v230 == v21 {
		goto L5
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L11
L102:
	;
	v29 = v230
	goto L10
}
func F_halfvec_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int64
	_ = v297
	var v301 int64
	_ = v301
	var v303 int32
	_ = v303
	var v306 int64
	_ = v306
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int64
	_ = v434
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v439 int64
	_ = v439
	var v441 int32
	_ = v441
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int64
	_ = v464
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int64
	_ = v507
	var v511 int64
	_ = v511
	var v513 int32
	_ = v513
	var v516 int64
	_ = v516
	var v518 int32
	_ = v518
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int64
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1425 int32
	_ = v1425
	var v1428 int64
	_ = v1428
	var v1430 int64
	_ = v1430
	var v1431 int64
	_ = v1431
	var v1433 int64
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1437 int64
	_ = v1437
	var v1438 int64
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int64
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1464 int64
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1474 int32
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int64
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1504 int64
	_ = v1504
	var v1508 int64
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1513 int64
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1526 int32
	_ = v1526
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1542 int32
	_ = v1542
	var v1548 int32
	_ = v1548
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1614 int32
	_ = v1614
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int64
	_ = v1641
	var v1643 int64
	_ = v1643
	var v1644 int64
	_ = v1644
	var v1646 int64
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1650 int64
	_ = v1650
	var v1651 int64
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1671 int64
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1676 int64
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1686 int32
	_ = v1686
	var v1687 int64
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1692 int64
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1714 int64
	_ = v1714
	var v1718 int64
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1723 int64
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1738 int32
	_ = v1738
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2082 int32
	_ = v2082
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2105 int32
	_ = v2105
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2247 int64
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2296 int32
	_ = v2296
	var v2299 int32
	_ = v2299
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2339 int32
	_ = v2339
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2354 int32
	_ = v2354
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2397 int32
	_ = v2397
	var v2400 int32
	_ = v2400
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2426 int32
	_ = v2426
	var v2429 int32
	_ = v2429
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
	v15 = F_mul_size(m, int32(16), v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_add_size(m, v15, int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = F_palloc(m, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v22)
	v24 = int32(1)
	v26 = v20 + v24
	if v14 <= int32(0) {
		v2429 = v26
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v2435 = int32(93)
	*(*uint16)(unsafe.Add(mBase, uint32(v2429))) = uint16(v2435)
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2437 != v10 {
		goto L424
	} else {
		goto L425
	}
L7:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+8)))
	v34 = v29 & int32(1023)
	v38 = v29 << (uint(int32(16)) % 32) & int32(-2147483648)
	v41 = int32(31)
	v42 = int32(base.Ui32(v29)>>(uint(int32(10))%32)) & v41
	if v42 != v41 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v120 = int32(0)
	v137 = base.I32_reinterpret_f32(base.F32_reinterpret_i32(v115 | v114<<(uint(int32(13))%32)))
	v139 = v137 & int32(_a_F_halfvec_out_0)
	v142 = int32(255)
	v143 = int32(base.Ui32(v137)>>(uint(int32(23))%32)) & v142
	if v143|v139 != 0 {
		goto L52
	} else {
		goto L53
	}
L9:
	;
	goto L8
L10:
	;
	v114 = v34
	v115 = v42<<(uint(int32(23))%32) + v38 + int32(939524096)
	goto L9
L11:
	;
	if v29&int32(512) != 0 {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v42 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v34 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v34 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v114 = int32(0)
	v115 = v38
	goto L9
L17:
	;
	v114 = int32(0)
	v115 = v38 | int32(2139095040)
	goto L9
L18:
	;
	goto L19
L19:
	;
	v114 = v34
	v115 = v38 | int32(2143289344)
	goto L9
L20:
	;
	v114 = v102 & int32(1022)
	v115 = v104 | v38
	goto L9
L21:
	;
	v102 = v34 << (uint(int32(1)) % 32)
	v104 = int32(939524096)
	goto L20
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(int32(255)) < base.Ui32(v34) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v102 = v34 << (uint(int32(2)) % 32)
	v104 = int32(931135488)
	goto L20
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(127)) < base.Ui32(v34) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v102 = v34 << (uint(int32(3)) % 32)
	v104 = int32(922746880)
	goto L20
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(63)) < base.Ui32(v34) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v102 = v34 << (uint(int32(4)) % 32)
	v104 = int32(914358272)
	goto L20
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(31)) < base.Ui32(v34) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v34 << (uint(int32(5)) % 32)
	v104 = int32(905969664)
	goto L20
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(15)) < base.Ui32(v34) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = v34 << (uint(int32(6)) % 32)
	v104 = int32(897581056)
	goto L20
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(7)) < base.Ui32(v34) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v102 = v34 << (uint(int32(7)) % 32)
	v104 = int32(889192448)
	goto L20
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(3)) < base.Ui32(v34) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v102 = v34 << (uint(int32(8)) % 32)
	v104 = int32(880803840)
	goto L20
L43:
	;
	goto L44
L44:
	;
	v97 = base.B2i32(v34 == int32(1))
	if v34 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v98 = int32(1024)
	goto L47
L46:
	;
	v98 = v34 << (uint(int32(9)) % 32)
	goto L47
L47:
	;
	if v34 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v101 = int32(864026624)
	goto L50
L49:
	;
	v101 = int32(872415232)
	goto L50
L50:
	;
	v102 = v98
	v104 = v101
	goto L20
L51:
	;
	v1217 = v1216 + v26
	if v14 == int32(1) {
		v2429 = v1217
		goto L6
	} else {
		goto L214
	}
L52:
	;
	v148 = base.B2i32(v143 != v142)
	goto L54
L53:
	;
	v148 = v120
	goto L54
L54:
	;
	if v148 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v139 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(int32(23)) < base.Ui32(v143-int32(127)) {
		goto L75
	} else {
		goto L76
	}
L58:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_halfvec_out[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2)) = uint8(v152)
	v155 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_halfvec_out[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v155)
	v1216 = int32(3)
	goto L51
L59:
	;
	goto L60
L60:
	;
	if v137 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v160 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v160)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v164 = v26 + int32(base.Ui32(v137)>>(uint(int32(31))%32))
	if v143 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = int64(8751735898823355977)
	if v137 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v172 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v172)
	if v137 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v171 = int32(9)
	goto L69
L68:
	;
	v171 = int32(8)
	goto L69
L69:
	;
	v1216 = v171
	goto L51
L70:
	;
	v178 = int32(2)
	goto L72
L71:
	;
	v178 = int32(1)
	goto L72
L72:
	;
	v1216 = v178
	goto L51
L73:
	;
	v826 = int32(0)
	if v137 < v826 {
		goto L141
	} else {
		goto L142
	}
L74:
	;
	if base.Ui32(int32(_a_F_halfvec_out_1)) < base.Ui32(v767) {
		v809 = v767
		v814 = v772
		v825 = int32(8)
		goto L73
	} else {
		goto L132
	}
L75:
	;
	v195 = v139 << (uint(int32(2)) % 32)
	if v143 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v183 = int32(-1)
	v185 = int32(150) - v143
	if v139&(v183<<(uint(v185)%32)^v183) != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v767 = int32(base.Ui32(v139|int32(_a_F_halfvec_out_2)) >> (uint(v185) % 32))
	v772 = v120
	goto L74
L78:
	;
	v198 = v195 | int32(33554432)
	goto L80
L79:
	;
	v198 = v195
	goto L80
L80:
	;
	v201 = int32(2)
	v206 = v198 + (base.B2i32(v139 != int32(0)) | base.B2i32(base.Ui32(v143) < base.Ui32(v201)) ^ int32(-1))
	v208 = v198 | v201
	if v143 != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v759 = v742 + v746
	v760 = v741 + v758
	if base.Ui32(v760) <= base.Ui32(int32(99999999)) {
		v767 = v760
		v772 = v759
		goto L74
	} else {
		goto L131
	}
L82:
	;
	v677 = int32(0)
	v678 = int32(10)
	v679 = base.I32_div_u_s(v663, v678)
	v681 = base.I32_div_u_s(v667, v678)
	if base.Ui32(v681) < base.Ui32(v679) {
		goto L125
	} else {
		goto L126
	}
L83:
	;
	v581 = int32(0)
	v582 = int32(10)
	v583 = base.I32_div_u_s(v567, v582)
	v585 = base.I32_div_u_s(v571, v582)
	if base.Ui32(v583) <= base.Ui32(v585) {
		goto L119
	} else {
		goto L120
	}
L84:
	;
	v212 = v143 - int32(152)
	goto L86
L85:
	;
	v212 = int32(-151)
	goto L86
L86:
	;
	if int32(0) <= v212 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v218 = int32(base.Ui32(v212*int32(_a_F_halfvec_out_3)) >> (uint(int32(18)) % 32))
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v218<<(uint(int32(3))%32))+uint32(_c_F_halfvec_out[2])))
	v223 = v221 & int64(4294967295)
	v224 = base.I64_extend_i32_u(v206)
	v226 = int64(32)
	v228 = base.I32_wrap_i64(int64(base.Ui64(v223*v224) >> (uint(v226) % 64)))
	v230 = int64(base.Ui64(v221) >> (uint(v226) % 64))
	v231 = v224 * v230
	v233 = v228 + base.I32_wrap_i64(v231)
	v240 = v218 - v212
	v245 = v240 + int32(base.Ui32(v218*int32(_a_F_halfvec_out_4))>>(uint(int32(19))%32))
	v246 = int32(5) - v245
	v249 = v245 + int32(27)
	v251 = (base.B2i32(base.Ui32(v233) < base.Ui32(v228))+base.I32_wrap_i64(int64(base.Ui64(v231)>>(uint(v226)%64))))<<(uint(v246)%32) | int32(base.Ui32(v233)>>(uint(v249)%32))
	v252 = base.I64_extend_i32_u(v208)
	v256 = base.I32_wrap_i64(int64(base.Ui64(v223*v252) >> (uint(v226) % 64)))
	v257 = v252 * v230
	v259 = v256 + base.I32_wrap_i64(v257)
	v267 = (base.B2i32(base.Ui32(v259) < base.Ui32(v256))+base.I32_wrap_i64(int64(base.Ui64(v257)>>(uint(v226)%64))))<<(uint(v246)%32) | int32(base.Ui32(v259)>>(uint(v249)%32))
	v268 = base.I64_extend_i32_u(v198)
	v272 = base.I32_wrap_i64(int64(base.Ui64(v223*v268) >> (uint(v226) % 64)))
	v273 = v268 * v230
	v275 = v272 + base.I32_wrap_i64(v273)
	v283 = (base.B2i32(base.Ui32(v275) < base.Ui32(v272))+base.I32_wrap_i64(int64(base.Ui64(v273)>>(uint(v226)%64))))<<(uint(v246)%32) | int32(base.Ui32(v275)>>(uint(v249)%32))
	v284 = int32(0)
	if v218 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v427 = v212 * int32(-732923)
	v429 = int32(base.Ui32(v427) >> (uint(int32(20)) % 32))
	v430 = v212 + v429
	v434 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_halfvec_out_5)-v430<<(uint(int32(3))%32))))
	v436 = v434 & int64(4294967295)
	v437 = base.I64_extend_i32_u(v206)
	v439 = int64(32)
	v441 = base.I32_wrap_i64(int64(base.Ui64(v436*v437) >> (uint(v439) % 64)))
	v443 = int64(base.Ui64(v434) >> (uint(v439) % 64))
	v444 = v437 * v443
	v446 = v441 + base.I32_wrap_i64(v444)
	v457 = v429 - int32(base.Ui32(v430*int32(-1217359))>>(uint(int32(19))%32))
	v458 = int32(4) - v457
	v461 = v457 + int32(28)
	v463 = (base.B2i32(base.Ui32(v446) < base.Ui32(v441))+base.I32_wrap_i64(int64(base.Ui64(v444)>>(uint(v439)%64))))<<(uint(v458)%32) | int32(base.Ui32(v446)>>(uint(v461)%32))
	v464 = base.I64_extend_i32_u(v198)
	v468 = base.I32_wrap_i64(int64(base.Ui64(v436*v464) >> (uint(v439) % 64)))
	v469 = v464 * v443
	v471 = v468 + base.I32_wrap_i64(v469)
	v479 = (base.B2i32(base.Ui32(v471) < base.Ui32(v468))+base.I32_wrap_i64(int64(base.Ui64(v469)>>(uint(v439)%64))))<<(uint(v458)%32) | int32(base.Ui32(v471)>>(uint(v461)%32))
	v480 = base.I64_extend_i32_u(v208)
	v484 = base.I32_wrap_i64(int64(base.Ui64(v436*v480) >> (uint(v439) % 64)))
	v485 = v443 * v480
	v487 = v484 + base.I32_wrap_i64(v485)
	v495 = (base.B2i32(base.Ui32(v487) < base.Ui32(v484))+base.I32_wrap_i64(int64(base.Ui64(v485)>>(uint(v439)%64))))<<(uint(v458)%32) | int32(base.Ui32(v487)>>(uint(v461)%32))
	v497 = v495 - int32(1)
	if v429 != 0 {
		goto L111
	} else {
		goto L112
	}
L90:
	;
	v288 = int32(10)
	v289 = base.I32_div_u_s(v267-int32(1), v288)
	v291 = base.I32_div_u_s(v251, v288)
	if base.Ui32(v289) <= base.Ui32(v291) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v335 = v284
	goto L92
L92:
	;
	v341 = base.I32_rem_u_s(v198, int32(5))
	if v341 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v294 = v218 - int32(1)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v294<<(uint(int32(3))%32))+uint32(_c_F_halfvec_out[2])))
	v301 = int64(32)
	v303 = base.I32_wrap_i64(int64(base.Ui64(v297&int64(4294967295)*v268) >> (uint(v301) % 64)))
	v306 = int64(base.Ui64(v297)>>(uint(v301)%64)) * v268
	v308 = v303 + base.I32_wrap_i64(v306)
	v319 = v240 + int32(base.Ui32(v294*int32(_a_F_halfvec_out_4))>>(uint(int32(19))%32))
	v327 = base.I32_rem_u_s((base.B2i32(base.Ui32(v308) < base.Ui32(v303))+base.I32_wrap_i64(int64(base.Ui64(v306)>>(uint(v301)%64))))<<(uint(int32(6)-v319)%32)|int32(base.Ui32(v308)>>(uint(v319+int32(26))%32)), int32(10))
	v328 = v327
	goto L95
L94:
	;
	v328 = v284
	goto L95
L95:
	;
	if base.Ui32(int32(33)) < base.Ui32(v212) {
		v660 = v283
		v663 = v267
		v665 = v218
		v666 = v328
		v667 = v251
		goto L82
	} else {
		goto L96
	}
L96:
	;
	v335 = v328
	goto L92
L97:
	;
	v347 = v198
	v348 = v284
	goto L100
L98:
	;
	goto L99
L99:
	;
	v372 = int32(0)
	v374 = base.I32_rem_u_s(v208, int32(5))
	if v374 == v372 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v364 = v348 + int32(1)
	v365 = int32(5)
	v366 = base.I32_div_u_s(v347, v365)
	v368 = base.I32_rem_u_s(v366, v365)
	if v368 == int32(0) {
		v347 = v366
		v348 = v364
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if base.Ui32(v364) < base.Ui32(v218) {
		v660 = v283
		v663 = v267
		v665 = v218
		v666 = v335
		v667 = v251
		goto L82
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v564 = v283
	v567 = v267
	v569 = v218
	v570 = v335
	v571 = v251
	goto L83
L104:
	;
	v380 = v372
	v383 = v208
	goto L107
L105:
	;
	v407 = v372
	goto L106
L106:
	;
	v660 = v283
	v663 = v267 - base.B2i32(base.Ui32(v218) <= base.Ui32(v407))
	v665 = v218
	v666 = v335
	v667 = v251
	goto L82
L107:
	;
	v397 = v380 + int32(1)
	v398 = int32(5)
	v399 = base.I32_div_u_s(v383, v398)
	v401 = base.I32_rem_u_s(v399, v398)
	if v401 == int32(0) {
		v380 = v397
		v383 = v399
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v407 = v397
	goto L106
L109:
	;
	goto L108
L110:
	;
	v552 = int32(-1)
	if v198&(v552<<(uint(v429-int32(1))%32)^v552)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v427)) != 0 {
		v660 = v479
		v663 = v495
		v665 = v430
		v666 = v541
		v667 = v463
		goto L82
	} else {
		goto L118
	}
L111:
	;
	v498 = int32(10)
	v499 = base.I32_div_u_s(v497, v498)
	v501 = base.I32_div_u_s(v463, v498)
	if base.Ui32(v499) <= base.Ui32(v501) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v548 = v120
	goto L113
L113:
	;
	v564 = v479
	v567 = v497
	v569 = v430
	v570 = v548
	v571 = v463
	goto L83
L114:
	;
	v504 = int32(1) - v430
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v504<<(uint(int32(3))%32))+uint32(_c_F_halfvec_out[3])))
	v511 = int64(32)
	v513 = base.I32_wrap_i64(int64(base.Ui64(v507&int64(4294967295)*v464) >> (uint(v511) % 64)))
	v516 = int64(base.Ui64(v507)>>(uint(v511)%64)) * v464
	v518 = v513 + base.I32_wrap_i64(v516)
	v531 = v429 + (int32(base.Ui32(v504*int32(_a_F_halfvec_out_4))>>(uint(int32(19))%32)) ^ int32(-1))
	v539 = base.I32_rem_u_s((base.B2i32(base.Ui32(v518) < base.Ui32(v513))+base.I32_wrap_i64(int64(base.Ui64(v516)>>(uint(v511)%64))))<<(uint(int32(4)-v531)%32)|int32(base.Ui32(v518)>>(uint(v531+int32(28))%32)), int32(10))
	v541 = v539
	goto L116
L115:
	;
	v541 = v120
	goto L116
L116:
	;
	if v429 != int32(1) {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v548 = v541
	goto L113
L118:
	;
	v564 = v479
	v567 = v495
	v569 = v430
	v570 = v541
	v571 = v463
	goto L83
L119:
	;
	v629 = v564
	v630 = v581
	v635 = v570
	v636 = v571
	v646 = int32(0)
	goto L121
L120:
	;
	v591 = v564
	v592 = v581
	v593 = v583
	v594 = int32(1)
	v595 = v585
	v597 = v570
	goto L122
L121:
	;
	v648 = v635 & int32(255)
	v741 = v629
	v742 = v630
	v746 = v569
	v758 = (v646|base.B2i32(v648 != int32(5))|v629)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v648)) | base.B2i32(v629 == v636)
	goto L81
L122:
	;
	v609 = v592 + int32(1)
	v610 = int32(10)
	v611 = base.I32_div_u_s(v591, v610)
	v614 = v591 - v611*v610
	v619 = v594 & base.B2i32(v597&int32(255) == int32(0))
	v621 = base.I32_div_u_s(v593, v610)
	v623 = base.I32_div_u_s(v595, v610)
	if base.Ui32(v623) < base.Ui32(v621) {
		v591 = v611
		v592 = v609
		v593 = v621
		v594 = v619
		v595 = v623
		v597 = v614
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v629 = v611
	v630 = v609
	v635 = v614
	v636 = v595
	v646 = v619 ^ int32(1)
	goto L121
L124:
	;
	goto L123
L125:
	;
	v685 = v660
	v686 = v677
	v687 = v679
	v689 = v681
	goto L128
L126:
	;
	v716 = v660
	v717 = v677
	v722 = v666
	v723 = v667
	goto L127
L127:
	;
	v741 = v716
	v742 = v717
	v746 = v665
	v758 = base.B2i32(v716 == v723) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v722&int32(255)))
	goto L81
L128:
	;
	v703 = v686 + int32(1)
	v704 = int32(10)
	v705 = base.I32_div_u_s(v685, v704)
	v707 = base.I32_div_u_s(v687, v704)
	v709 = base.I32_div_u_s(v689, v704)
	if base.Ui32(v709) < base.Ui32(v707) {
		v685 = v705
		v686 = v703
		v687 = v707
		v689 = v709
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v716 = v705
	v717 = v703
	v722 = v685 - v705*int32(10)
	v723 = v689
	goto L127
L130:
	;
	goto L129
L131:
	;
	v809 = v760
	v814 = v759
	v825 = int32(9)
	goto L73
L132:
	;
	if base.Ui32(int32(_a_F_halfvec_out_6)) < base.Ui32(v767) {
		v809 = v767
		v814 = v772
		v825 = int32(7)
		goto L73
	} else {
		goto L133
	}
L133:
	;
	if base.Ui32(int32(_a_F_halfvec_out_7)) < base.Ui32(v767) {
		v809 = v767
		v814 = v772
		v825 = int32(6)
		goto L73
	} else {
		goto L134
	}
L134:
	;
	if base.Ui32(int32(_a_F_halfvec_out_8)) < base.Ui32(v767) {
		v809 = v767
		v814 = v772
		v825 = int32(5)
		goto L73
	} else {
		goto L135
	}
L135:
	;
	if base.Ui32(int32(999)) < base.Ui32(v767) {
		v809 = v767
		v814 = v772
		v825 = int32(4)
		goto L73
	} else {
		goto L136
	}
L136:
	;
	if base.Ui32(int32(99)) < base.Ui32(v767) {
		v809 = v767
		v814 = v772
		v825 = int32(3)
		goto L73
	} else {
		goto L137
	}
L137:
	;
	if base.Ui32(int32(9)) < base.Ui32(v767) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v805 = int32(2)
	goto L140
L139:
	;
	v805 = int32(1)
	goto L140
L140:
	;
	v809 = v767
	v814 = v772
	v825 = v805
	goto L73
L141:
	;
	v829 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v829)
	v832 = int32(1)
	goto L143
L142:
	;
	v832 = v826
	goto L143
L143:
	;
	v833 = v825 + v814
	if base.Ui32(v833+int32(3)) <= base.Ui32(int32(9)) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v1044 = int32(0)
	if base.Ui32(int32(_a_F_halfvec_out_9)) <= base.Ui32(v809) {
		goto L184
	} else {
		goto L185
	}
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v838))) = v1040
	v1042 = v1039
	goto L144
L146:
	;
	v838 = v26 + v832
	v839 = int32(0)
	if v833 <= v839 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L148
L148:
	;
	if v814 != 0 {
		goto L154
	} else {
		goto L155
	}
L149:
	;
	v1039 = int32(2) - v833
	v1040 = int64(3472328296227679792)
	goto L145
L150:
	;
	goto L151
L151:
	;
	if int32(0) <= v814 {
		v1039 = v839
		v1040 = int64(3472328296227680304)
		goto L145
	} else {
		goto L152
	}
L152:
	;
	v1042 = int32(1)
	goto L144
L153:
	;
	v898 = int32(0)
	if base.Ui32(int32(_a_F_halfvec_out_9)) <= base.Ui32(v881) {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	v881 = v809
	v886 = v825
	goto L153
L155:
	;
	goto L156
L156:
	;
	v852 = v809
	v854 = v825
	goto L157
L157:
	;
	if v852&int32(1) != 0 {
		v881 = v852
		v886 = v854
		goto L153
	} else {
		goto L159
	}
L158:
	;
	v881 = v852
	v886 = v854
	goto L153
L159:
	;
	v875 = base.I32_div_u_s(v852, int32(10))
	if int32(0)-v852 == v875*int32(-10) {
		v852 = v875
		v854 = v854 - int32(1)
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v905 = v881
	v906 = v898
	goto L164
L162:
	;
	v951 = v881
	v952 = v898
	goto L163
L163:
	;
	if base.Ui32(v951) < base.Ui32(int32(100)) {
		goto L168
	} else {
		goto L169
	}
L164:
	;
	v922 = v26 + v832 + v886 - v906
	v926 = base.I32_div_u_s(v905, int32(_a_F_halfvec_out_9))
	v929 = v905 + v926*int32(-10000)
	v930 = int32(100)
	v931 = base.I32_div_u_s(v929, v930)
	v932 = int32(1)
	v934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v931<<(uint(v932)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v922-int32(3)))) = uint16(v934)
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v929-v931*v930)<<(uint(v932)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v922-v932))) = uint16(v943)
	v946 = v906 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v905) {
		v905 = v926
		v906 = v946
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v951 = v926
	v952 = v946
	goto L163
L166:
	;
	goto L165
L167:
	;
	v993 = v833 - int32(1)
	v994 = v26 + v832
	if base.Ui32(int32(10)) <= base.Ui32(v991) {
		goto L172
	} else {
		goto L173
	}
L168:
	;
	v990 = v952
	v991 = v951
	goto L167
L169:
	;
	goto L170
L170:
	;
	v975 = int32(_a_F_halfvec_out_10)
	v977 = int32(100)
	v978 = base.I32_div_u_s(v951&v975, v977)
	v986 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v951-v978*v977)&v975<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v26+v832+v886+(v952^int32(-1))))) = uint16(v986)
	v990 = v952 | int32(2)
	v991 = v978
	goto L167
L171:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v994))) = uint8(v1008)
	if base.Ui32(int32(2)) <= base.Ui32(v886) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1001 = v991 << (uint(int32(1)) % 32)
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+uint32(_c_F_halfvec_out[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+(v832+v886-v990)))) = uint8(v1002)
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+uint32(_c_F_halfvec_out[4]))))
	v1008 = v1004
	goto L171
L173:
	;
	goto L174
L174:
	;
	v1008 = v991 | int32(48)
	goto L171
L175:
	;
	v1012 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v994)+1)) = uint8(v1012)
	v1017 = v886 + int32(1)
	goto L177
L176:
	;
	v1017 = int32(1)
	goto L177
L177:
	;
	v1018 = v1017 + v832
	v1019 = v26 + v1018
	v1020 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019))) = uint8(v1020)
	v1025 = base.B2i32(v993 < int32(0))
	if v993 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1026 = int32(45)
	goto L180
L179:
	;
	v1026 = int32(43)
	goto L180
L180:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+1)) = uint8(v1026)
	if v993 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1030 = int32(1) - v833
	goto L183
L182:
	;
	v1030 = v993
	goto L183
L183:
	;
	v1035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1030<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1019)+2)) = uint16(v1035)
	v1216 = v1018 + int32(4)
	goto L51
L184:
	;
	v1051 = v1044
	v1052 = v809
	goto L187
L185:
	;
	v1097 = v1044
	v1098 = v809
	goto L186
L186:
	;
	if base.Ui32(v1098) < base.Ui32(int32(100)) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v1068 = v1042 + v838 + v825 - v1051
	v1069 = int32(4)
	v1072 = base.I32_div_u_s(v1052, int32(_a_F_halfvec_out_9))
	v1075 = v1052 + v1072*int32(-10000)
	v1076 = int32(100)
	v1077 = base.I32_div_u_s(v1075, v1076)
	v1078 = int32(1)
	v1080 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1077<<(uint(v1078)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1068-v1069))) = uint16(v1080)
	v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1075-v1077*v1076)<<(uint(v1078)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1068-int32(2)))) = uint16(v1089)
	v1092 = v1051 + v1069
	if base.Ui32(int32(99999999)) < base.Ui32(v1052) {
		v1051 = v1092
		v1052 = v1072
		goto L187
	} else {
		goto L189
	}
L188:
	;
	v1097 = v1092
	v1098 = v1072
	goto L186
L189:
	;
	goto L188
L190:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1137) {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v1136 = v1097
	v1137 = v1098
	goto L190
L192:
	;
	goto L193
L193:
	;
	v1119 = int32(2)
	v1121 = int32(_a_F_halfvec_out_10)
	v1123 = int32(100)
	v1124 = base.I32_div_u_s(v1098&v1121, v1123)
	v1132 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1098-v1124*v1123)&v1121<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1042+v838+v825-v1097-v1119))) = uint16(v1132)
	v1136 = v1097 | v1119
	v1137 = v1124
	goto L190
L194:
	;
	v1153 = int32(1)
	if v1042 == v1153 {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1137<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1042+v838+v825-v1136-int32(2)))) = uint16(v1147)
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1151 = v1137 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1042+v838))) = uint8(v1151)
	goto L194
L198:
	;
	v1216 = v1193 + int32(base.Ui32(v137)>>(uint(int32(31))%32))
	goto L51
L199:
	;
	if v833&int32(4) != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	if v814 < int32(0) {
		goto L211
	} else {
		goto L212
	}
L202:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v838)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v838))) = v1158
	v1161 = int32(5)
	goto L204
L203:
	;
	v1161 = v1153
	goto L204
L204:
	;
	if v833&int32(2) != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1164 = v1161 + v838
	v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1164))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1164-int32(1)))) = uint16(v1167)
	v1172 = v1161 | int32(2)
	goto L207
L206:
	;
	v1172 = v1161
	goto L207
L207:
	;
	if v833&int32(1) != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1175 = v1172 + v838
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1175-int32(1)))) = uint8(v1178)
	goto L210
L209:
	;
	goto L210
L210:
	;
	v1182 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v838+v833))) = uint8(v1182)
	v1193 = v825 + int32(1)
	goto L198
L211:
	;
	v1190 = int32(2) - v814
	goto L213
L212:
	;
	v1190 = v833
	goto L213
L213:
	;
	v1193 = v1190
	goto L198
L214:
	;
	v1223 = v1217
	v1226 = v24
	goto L215
L215:
	;
	v1229 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1223))) = uint8(v1229)
	v1231 = int32(1)
	v1232 = v1223 + v1231
	v1236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(8)+v1226<<(uint(v1231)%32)))))
	v1241 = v1236 & int32(1023)
	v1245 = v1236 << (uint(int32(16)) % 32) & int32(-2147483648)
	v1248 = int32(31)
	v1249 = int32(base.Ui32(v1236)>>(uint(int32(10))%32)) & v1248
	if v1249 != v1248 {
		goto L221
	} else {
		goto L222
	}
L216:
	;
	v2429 = v2424
	goto L6
L217:
	;
	v1327 = int32(0)
	v1344 = base.I32_reinterpret_f32(base.F32_reinterpret_i32(v1322 | v1321<<(uint(int32(13))%32)))
	v1346 = v1344 & int32(_a_F_halfvec_out_0)
	v1349 = int32(255)
	v1350 = int32(base.Ui32(v1344)>>(uint(int32(23))%32)) & v1349
	if v1350|v1346 != 0 {
		goto L261
	} else {
		goto L262
	}
L218:
	;
	goto L217
L219:
	;
	v1321 = v1241
	v1322 = v1249<<(uint(int32(23))%32) + v1245 + int32(939524096)
	goto L218
L220:
	;
	if v1236&int32(512) != 0 {
		goto L230
	} else {
		goto L231
	}
L221:
	;
	if v1249 != 0 {
		goto L219
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v1241 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	if v1241 != 0 {
		goto L220
	} else {
		goto L225
	}
L225:
	;
	v1321 = int32(0)
	v1322 = v1245
	goto L218
L226:
	;
	v1321 = int32(0)
	v1322 = v1245 | int32(2139095040)
	goto L218
L227:
	;
	goto L228
L228:
	;
	v1321 = v1241
	v1322 = v1245 | int32(2143289344)
	goto L218
L229:
	;
	v1321 = v1309 & int32(1022)
	v1322 = v1311 | v1245
	goto L218
L230:
	;
	v1309 = v1241 << (uint(int32(1)) % 32)
	v1311 = int32(939524096)
	goto L229
L231:
	;
	goto L232
L232:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1241) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1309 = v1241 << (uint(int32(2)) % 32)
	v1311 = int32(931135488)
	goto L229
L234:
	;
	goto L235
L235:
	;
	if base.Ui32(int32(127)) < base.Ui32(v1241) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1309 = v1241 << (uint(int32(3)) % 32)
	v1311 = int32(922746880)
	goto L229
L237:
	;
	goto L238
L238:
	;
	if base.Ui32(int32(63)) < base.Ui32(v1241) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1309 = v1241 << (uint(int32(4)) % 32)
	v1311 = int32(914358272)
	goto L229
L240:
	;
	goto L241
L241:
	;
	if base.Ui32(int32(31)) < base.Ui32(v1241) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1309 = v1241 << (uint(int32(5)) % 32)
	v1311 = int32(905969664)
	goto L229
L243:
	;
	goto L244
L244:
	;
	if base.Ui32(int32(15)) < base.Ui32(v1241) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1309 = v1241 << (uint(int32(6)) % 32)
	v1311 = int32(897581056)
	goto L229
L246:
	;
	goto L247
L247:
	;
	if base.Ui32(int32(7)) < base.Ui32(v1241) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1309 = v1241 << (uint(int32(7)) % 32)
	v1311 = int32(889192448)
	goto L229
L249:
	;
	goto L250
L250:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1241) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1309 = v1241 << (uint(int32(8)) % 32)
	v1311 = int32(880803840)
	goto L229
L252:
	;
	goto L253
L253:
	;
	v1304 = base.B2i32(v1241 == int32(1))
	if v1241 == int32(1) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1305 = int32(1024)
	goto L256
L255:
	;
	v1305 = v1241 << (uint(int32(9)) % 32)
	goto L256
L256:
	;
	if v1241 == int32(1) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1308 = int32(864026624)
	goto L259
L258:
	;
	v1308 = int32(872415232)
	goto L259
L259:
	;
	v1309 = v1305
	v1311 = v1308
	goto L229
L260:
	;
	v2424 = v2423 + v1232
	v2426 = v1226 + int32(1)
	if v2426 != v14 {
		v1223 = v2424
		v1226 = v2426
		goto L215
	} else {
		goto L423
	}
L261:
	;
	v1355 = base.B2i32(v1350 != v1349)
	goto L263
L262:
	;
	v1355 = v1327
	goto L263
L263:
	;
	if v1355 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	if v1346 != 0 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	if base.Ui32(int32(23)) < base.Ui32(v1350-int32(127)) {
		goto L284
	} else {
		goto L285
	}
L267:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_halfvec_out[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1232)+2)) = uint8(v1359)
	v1362 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_halfvec_out[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1232))) = uint16(v1362)
	v2423 = int32(3)
	goto L260
L268:
	;
	goto L269
L269:
	;
	if v1344 < int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1367 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1232))) = uint8(v1367)
	goto L272
L271:
	;
	goto L272
L272:
	;
	v1371 = v1232 + int32(base.Ui32(v1344)>>(uint(int32(31))%32))
	if v1350 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1371))) = int64(8751735898823355977)
	if v1344 < int32(0) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	v1379 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1371))) = uint8(v1379)
	if v1344 < int32(0) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v1378 = int32(9)
	goto L278
L277:
	;
	v1378 = int32(8)
	goto L278
L278:
	;
	v2423 = v1378
	goto L260
L279:
	;
	v1385 = int32(2)
	goto L281
L280:
	;
	v1385 = int32(1)
	goto L281
L281:
	;
	v2423 = v1385
	goto L260
L282:
	;
	v2033 = int32(0)
	if v1344 < v2033 {
		goto L350
	} else {
		goto L351
	}
L283:
	;
	if base.Ui32(int32(_a_F_halfvec_out_1)) < base.Ui32(v1974) {
		v2016 = v1974
		v2021 = v1979
		v2032 = int32(8)
		goto L282
	} else {
		goto L341
	}
L284:
	;
	v1402 = v1346 << (uint(int32(2)) % 32)
	if v1350 != 0 {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v1390 = int32(-1)
	v1392 = int32(150) - v1350
	if v1346&(v1390<<(uint(v1392)%32)^v1390) != 0 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1974 = int32(base.Ui32(v1346|int32(_a_F_halfvec_out_2)) >> (uint(v1392) % 32))
	v1979 = v1327
	goto L283
L287:
	;
	v1405 = v1402 | int32(33554432)
	goto L289
L288:
	;
	v1405 = v1402
	goto L289
L289:
	;
	v1408 = int32(2)
	v1413 = v1405 + (base.B2i32(v1346 != int32(0)) | base.B2i32(base.Ui32(v1350) < base.Ui32(v1408)) ^ int32(-1))
	v1415 = v1405 | v1408
	if v1350 != 0 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	v1966 = v1949 + v1953
	v1967 = v1948 + v1965
	if base.Ui32(v1967) <= base.Ui32(int32(99999999)) {
		v1974 = v1967
		v1979 = v1966
		goto L283
	} else {
		goto L340
	}
L291:
	;
	v1884 = int32(0)
	v1885 = int32(10)
	v1886 = base.I32_div_u_s(v1870, v1885)
	v1888 = base.I32_div_u_s(v1874, v1885)
	if base.Ui32(v1888) < base.Ui32(v1886) {
		goto L334
	} else {
		goto L335
	}
L292:
	;
	v1788 = int32(0)
	v1789 = int32(10)
	v1790 = base.I32_div_u_s(v1774, v1789)
	v1792 = base.I32_div_u_s(v1778, v1789)
	if base.Ui32(v1790) <= base.Ui32(v1792) {
		goto L328
	} else {
		goto L329
	}
L293:
	;
	v1419 = v1350 - int32(152)
	goto L295
L294:
	;
	v1419 = int32(-151)
	goto L295
L295:
	;
	if int32(0) <= v1419 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1425 = int32(base.Ui32(v1419*int32(_a_F_halfvec_out_3)) >> (uint(int32(18)) % 32))
	v1428 = *(*int64)(unsafe.Add(mBase, uint32(v1425<<(uint(int32(3))%32))+uint32(_c_F_halfvec_out[2])))
	v1430 = v1428 & int64(4294967295)
	v1431 = base.I64_extend_i32_u(v1413)
	v1433 = int64(32)
	v1435 = base.I32_wrap_i64(int64(base.Ui64(v1430*v1431) >> (uint(v1433) % 64)))
	v1437 = int64(base.Ui64(v1428) >> (uint(v1433) % 64))
	v1438 = v1431 * v1437
	v1440 = v1435 + base.I32_wrap_i64(v1438)
	v1447 = v1425 - v1419
	v1452 = v1447 + int32(base.Ui32(v1425*int32(_a_F_halfvec_out_4))>>(uint(int32(19))%32))
	v1453 = int32(5) - v1452
	v1456 = v1452 + int32(27)
	v1458 = (base.B2i32(base.Ui32(v1440) < base.Ui32(v1435))+base.I32_wrap_i64(int64(base.Ui64(v1438)>>(uint(v1433)%64))))<<(uint(v1453)%32) | int32(base.Ui32(v1440)>>(uint(v1456)%32))
	v1459 = base.I64_extend_i32_u(v1415)
	v1463 = base.I32_wrap_i64(int64(base.Ui64(v1430*v1459) >> (uint(v1433) % 64)))
	v1464 = v1459 * v1437
	v1466 = v1463 + base.I32_wrap_i64(v1464)
	v1474 = (base.B2i32(base.Ui32(v1466) < base.Ui32(v1463))+base.I32_wrap_i64(int64(base.Ui64(v1464)>>(uint(v1433)%64))))<<(uint(v1453)%32) | int32(base.Ui32(v1466)>>(uint(v1456)%32))
	v1475 = base.I64_extend_i32_u(v1405)
	v1479 = base.I32_wrap_i64(int64(base.Ui64(v1430*v1475) >> (uint(v1433) % 64)))
	v1480 = v1475 * v1437
	v1482 = v1479 + base.I32_wrap_i64(v1480)
	v1490 = (base.B2i32(base.Ui32(v1482) < base.Ui32(v1479))+base.I32_wrap_i64(int64(base.Ui64(v1480)>>(uint(v1433)%64))))<<(uint(v1453)%32) | int32(base.Ui32(v1482)>>(uint(v1456)%32))
	v1491 = int32(0)
	if v1425 != 0 {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	goto L298
L298:
	;
	v1634 = v1419 * int32(-732923)
	v1636 = int32(base.Ui32(v1634) >> (uint(int32(20)) % 32))
	v1637 = v1419 + v1636
	v1641 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_halfvec_out_5)-v1637<<(uint(int32(3))%32))))
	v1643 = v1641 & int64(4294967295)
	v1644 = base.I64_extend_i32_u(v1413)
	v1646 = int64(32)
	v1648 = base.I32_wrap_i64(int64(base.Ui64(v1643*v1644) >> (uint(v1646) % 64)))
	v1650 = int64(base.Ui64(v1641) >> (uint(v1646) % 64))
	v1651 = v1644 * v1650
	v1653 = v1648 + base.I32_wrap_i64(v1651)
	v1664 = v1636 - int32(base.Ui32(v1637*int32(-1217359))>>(uint(int32(19))%32))
	v1665 = int32(4) - v1664
	v1668 = v1664 + int32(28)
	v1670 = (base.B2i32(base.Ui32(v1653) < base.Ui32(v1648))+base.I32_wrap_i64(int64(base.Ui64(v1651)>>(uint(v1646)%64))))<<(uint(v1665)%32) | int32(base.Ui32(v1653)>>(uint(v1668)%32))
	v1671 = base.I64_extend_i32_u(v1405)
	v1675 = base.I32_wrap_i64(int64(base.Ui64(v1643*v1671) >> (uint(v1646) % 64)))
	v1676 = v1671 * v1650
	v1678 = v1675 + base.I32_wrap_i64(v1676)
	v1686 = (base.B2i32(base.Ui32(v1678) < base.Ui32(v1675))+base.I32_wrap_i64(int64(base.Ui64(v1676)>>(uint(v1646)%64))))<<(uint(v1665)%32) | int32(base.Ui32(v1678)>>(uint(v1668)%32))
	v1687 = base.I64_extend_i32_u(v1415)
	v1691 = base.I32_wrap_i64(int64(base.Ui64(v1643*v1687) >> (uint(v1646) % 64)))
	v1692 = v1650 * v1687
	v1694 = v1691 + base.I32_wrap_i64(v1692)
	v1702 = (base.B2i32(base.Ui32(v1694) < base.Ui32(v1691))+base.I32_wrap_i64(int64(base.Ui64(v1692)>>(uint(v1646)%64))))<<(uint(v1665)%32) | int32(base.Ui32(v1694)>>(uint(v1668)%32))
	v1704 = v1702 - int32(1)
	if v1636 != 0 {
		goto L320
	} else {
		goto L321
	}
L299:
	;
	v1495 = int32(10)
	v1496 = base.I32_div_u_s(v1474-int32(1), v1495)
	v1498 = base.I32_div_u_s(v1458, v1495)
	if base.Ui32(v1496) <= base.Ui32(v1498) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	v1542 = v1491
	goto L301
L301:
	;
	v1548 = base.I32_rem_u_s(v1405, int32(5))
	if v1548 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L302:
	;
	v1501 = v1425 - int32(1)
	v1504 = *(*int64)(unsafe.Add(mBase, uint32(v1501<<(uint(int32(3))%32))+uint32(_c_F_halfvec_out[2])))
	v1508 = int64(32)
	v1510 = base.I32_wrap_i64(int64(base.Ui64(v1504&int64(4294967295)*v1475) >> (uint(v1508) % 64)))
	v1513 = int64(base.Ui64(v1504)>>(uint(v1508)%64)) * v1475
	v1515 = v1510 + base.I32_wrap_i64(v1513)
	v1526 = v1447 + int32(base.Ui32(v1501*int32(_a_F_halfvec_out_4))>>(uint(int32(19))%32))
	v1534 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1515) < base.Ui32(v1510))+base.I32_wrap_i64(int64(base.Ui64(v1513)>>(uint(v1508)%64))))<<(uint(int32(6)-v1526)%32)|int32(base.Ui32(v1515)>>(uint(v1526+int32(26))%32)), int32(10))
	v1535 = v1534
	goto L304
L303:
	;
	v1535 = v1491
	goto L304
L304:
	;
	if base.Ui32(int32(33)) < base.Ui32(v1419) {
		v1867 = v1490
		v1870 = v1474
		v1872 = v1425
		v1873 = v1535
		v1874 = v1458
		goto L291
	} else {
		goto L305
	}
L305:
	;
	v1542 = v1535
	goto L301
L306:
	;
	v1554 = v1405
	v1555 = v1491
	goto L309
L307:
	;
	goto L308
L308:
	;
	v1579 = int32(0)
	v1581 = base.I32_rem_u_s(v1415, int32(5))
	if v1581 == v1579 {
		goto L313
	} else {
		goto L314
	}
L309:
	;
	v1571 = v1555 + int32(1)
	v1572 = int32(5)
	v1573 = base.I32_div_u_s(v1554, v1572)
	v1575 = base.I32_rem_u_s(v1573, v1572)
	if v1575 == int32(0) {
		v1554 = v1573
		v1555 = v1571
		goto L309
	} else {
		goto L311
	}
L310:
	;
	if base.Ui32(v1571) < base.Ui32(v1425) {
		v1867 = v1490
		v1870 = v1474
		v1872 = v1425
		v1873 = v1542
		v1874 = v1458
		goto L291
	} else {
		goto L312
	}
L311:
	;
	goto L310
L312:
	;
	v1771 = v1490
	v1774 = v1474
	v1776 = v1425
	v1777 = v1542
	v1778 = v1458
	goto L292
L313:
	;
	v1587 = v1579
	v1590 = v1415
	goto L316
L314:
	;
	v1614 = v1579
	goto L315
L315:
	;
	v1867 = v1490
	v1870 = v1474 - base.B2i32(base.Ui32(v1425) <= base.Ui32(v1614))
	v1872 = v1425
	v1873 = v1542
	v1874 = v1458
	goto L291
L316:
	;
	v1604 = v1587 + int32(1)
	v1605 = int32(5)
	v1606 = base.I32_div_u_s(v1590, v1605)
	v1608 = base.I32_rem_u_s(v1606, v1605)
	if v1608 == int32(0) {
		v1587 = v1604
		v1590 = v1606
		goto L316
	} else {
		goto L318
	}
L317:
	;
	v1614 = v1604
	goto L315
L318:
	;
	goto L317
L319:
	;
	v1759 = int32(-1)
	if v1405&(v1759<<(uint(v1636-int32(1))%32)^v1759)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v1634)) != 0 {
		v1867 = v1686
		v1870 = v1702
		v1872 = v1637
		v1873 = v1748
		v1874 = v1670
		goto L291
	} else {
		goto L327
	}
L320:
	;
	v1705 = int32(10)
	v1706 = base.I32_div_u_s(v1704, v1705)
	v1708 = base.I32_div_u_s(v1670, v1705)
	if base.Ui32(v1706) <= base.Ui32(v1708) {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	v1755 = v1327
	goto L322
L322:
	;
	v1771 = v1686
	v1774 = v1704
	v1776 = v1637
	v1777 = v1755
	v1778 = v1670
	goto L292
L323:
	;
	v1711 = int32(1) - v1637
	v1714 = *(*int64)(unsafe.Add(mBase, uint32(v1711<<(uint(int32(3))%32))+uint32(_c_F_halfvec_out[3])))
	v1718 = int64(32)
	v1720 = base.I32_wrap_i64(int64(base.Ui64(v1714&int64(4294967295)*v1671) >> (uint(v1718) % 64)))
	v1723 = int64(base.Ui64(v1714)>>(uint(v1718)%64)) * v1671
	v1725 = v1720 + base.I32_wrap_i64(v1723)
	v1738 = v1636 + (int32(base.Ui32(v1711*int32(_a_F_halfvec_out_4))>>(uint(int32(19))%32)) ^ int32(-1))
	v1746 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1725) < base.Ui32(v1720))+base.I32_wrap_i64(int64(base.Ui64(v1723)>>(uint(v1718)%64))))<<(uint(int32(4)-v1738)%32)|int32(base.Ui32(v1725)>>(uint(v1738+int32(28))%32)), int32(10))
	v1748 = v1746
	goto L325
L324:
	;
	v1748 = v1327
	goto L325
L325:
	;
	if v1636 != int32(1) {
		goto L319
	} else {
		goto L326
	}
L326:
	;
	v1755 = v1748
	goto L322
L327:
	;
	v1771 = v1686
	v1774 = v1702
	v1776 = v1637
	v1777 = v1748
	v1778 = v1670
	goto L292
L328:
	;
	v1836 = v1771
	v1837 = v1788
	v1842 = v1777
	v1843 = v1778
	v1853 = int32(0)
	goto L330
L329:
	;
	v1798 = v1771
	v1799 = v1788
	v1800 = v1790
	v1801 = int32(1)
	v1802 = v1792
	v1804 = v1777
	goto L331
L330:
	;
	v1855 = v1842 & int32(255)
	v1948 = v1836
	v1949 = v1837
	v1953 = v1776
	v1965 = (v1853|base.B2i32(v1855 != int32(5))|v1836)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1855)) | base.B2i32(v1836 == v1843)
	goto L290
L331:
	;
	v1816 = v1799 + int32(1)
	v1817 = int32(10)
	v1818 = base.I32_div_u_s(v1798, v1817)
	v1821 = v1798 - v1818*v1817
	v1826 = v1801 & base.B2i32(v1804&int32(255) == int32(0))
	v1828 = base.I32_div_u_s(v1800, v1817)
	v1830 = base.I32_div_u_s(v1802, v1817)
	if base.Ui32(v1830) < base.Ui32(v1828) {
		v1798 = v1818
		v1799 = v1816
		v1800 = v1828
		v1801 = v1826
		v1802 = v1830
		v1804 = v1821
		goto L331
	} else {
		goto L333
	}
L332:
	;
	v1836 = v1818
	v1837 = v1816
	v1842 = v1821
	v1843 = v1802
	v1853 = v1826 ^ int32(1)
	goto L330
L333:
	;
	goto L332
L334:
	;
	v1892 = v1867
	v1893 = v1884
	v1894 = v1886
	v1896 = v1888
	goto L337
L335:
	;
	v1923 = v1867
	v1924 = v1884
	v1929 = v1873
	v1930 = v1874
	goto L336
L336:
	;
	v1948 = v1923
	v1949 = v1924
	v1953 = v1872
	v1965 = base.B2i32(v1923 == v1930) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1929&int32(255)))
	goto L290
L337:
	;
	v1910 = v1893 + int32(1)
	v1911 = int32(10)
	v1912 = base.I32_div_u_s(v1892, v1911)
	v1914 = base.I32_div_u_s(v1894, v1911)
	v1916 = base.I32_div_u_s(v1896, v1911)
	if base.Ui32(v1916) < base.Ui32(v1914) {
		v1892 = v1912
		v1893 = v1910
		v1894 = v1914
		v1896 = v1916
		goto L337
	} else {
		goto L339
	}
L338:
	;
	v1923 = v1912
	v1924 = v1910
	v1929 = v1892 - v1912*int32(10)
	v1930 = v1896
	goto L336
L339:
	;
	goto L338
L340:
	;
	v2016 = v1967
	v2021 = v1966
	v2032 = int32(9)
	goto L282
L341:
	;
	if base.Ui32(int32(_a_F_halfvec_out_6)) < base.Ui32(v1974) {
		v2016 = v1974
		v2021 = v1979
		v2032 = int32(7)
		goto L282
	} else {
		goto L342
	}
L342:
	;
	if base.Ui32(int32(_a_F_halfvec_out_7)) < base.Ui32(v1974) {
		v2016 = v1974
		v2021 = v1979
		v2032 = int32(6)
		goto L282
	} else {
		goto L343
	}
L343:
	;
	if base.Ui32(int32(_a_F_halfvec_out_8)) < base.Ui32(v1974) {
		v2016 = v1974
		v2021 = v1979
		v2032 = int32(5)
		goto L282
	} else {
		goto L344
	}
L344:
	;
	if base.Ui32(int32(999)) < base.Ui32(v1974) {
		v2016 = v1974
		v2021 = v1979
		v2032 = int32(4)
		goto L282
	} else {
		goto L345
	}
L345:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1974) {
		v2016 = v1974
		v2021 = v1979
		v2032 = int32(3)
		goto L282
	} else {
		goto L346
	}
L346:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1974) {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v2012 = int32(2)
	goto L349
L348:
	;
	v2012 = int32(1)
	goto L349
L349:
	;
	v2016 = v1974
	v2021 = v1979
	v2032 = v2012
	goto L282
L350:
	;
	v2036 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1232))) = uint8(v2036)
	v2039 = int32(1)
	goto L352
L351:
	;
	v2039 = v2033
	goto L352
L352:
	;
	v2040 = v2032 + v2021
	if base.Ui32(v2040+int32(3)) <= base.Ui32(int32(9)) {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	v2251 = int32(0)
	if base.Ui32(int32(_a_F_halfvec_out_9)) <= base.Ui32(v2016) {
		goto L393
	} else {
		goto L394
	}
L354:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2045))) = v2247
	v2249 = v2246
	goto L353
L355:
	;
	v2045 = v1232 + v2039
	v2046 = int32(0)
	if v2040 <= v2046 {
		goto L358
	} else {
		goto L359
	}
L356:
	;
	goto L357
L357:
	;
	if v2021 != 0 {
		goto L363
	} else {
		goto L364
	}
L358:
	;
	v2246 = int32(2) - v2040
	v2247 = int64(3472328296227679792)
	goto L354
L359:
	;
	goto L360
L360:
	;
	if int32(0) <= v2021 {
		v2246 = v2046
		v2247 = int64(3472328296227680304)
		goto L354
	} else {
		goto L361
	}
L361:
	;
	v2249 = int32(1)
	goto L353
L362:
	;
	v2105 = int32(0)
	if base.Ui32(int32(_a_F_halfvec_out_9)) <= base.Ui32(v2088) {
		goto L370
	} else {
		goto L371
	}
L363:
	;
	v2088 = v2016
	v2093 = v2032
	goto L362
L364:
	;
	goto L365
L365:
	;
	v2059 = v2016
	v2061 = v2032
	goto L366
L366:
	;
	if v2059&int32(1) != 0 {
		v2088 = v2059
		v2093 = v2061
		goto L362
	} else {
		goto L368
	}
L367:
	;
	v2088 = v2059
	v2093 = v2061
	goto L362
L368:
	;
	v2082 = base.I32_div_u_s(v2059, int32(10))
	if int32(0)-v2059 == v2082*int32(-10) {
		v2059 = v2082
		v2061 = v2061 - int32(1)
		goto L366
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	v2112 = v2088
	v2113 = v2105
	goto L373
L371:
	;
	v2158 = v2088
	v2159 = v2105
	goto L372
L372:
	;
	if base.Ui32(v2158) < base.Ui32(int32(100)) {
		goto L377
	} else {
		goto L378
	}
L373:
	;
	v2129 = v1232 + v2039 + v2093 - v2113
	v2133 = base.I32_div_u_s(v2112, int32(_a_F_halfvec_out_9))
	v2136 = v2112 + v2133*int32(-10000)
	v2137 = int32(100)
	v2138 = base.I32_div_u_s(v2136, v2137)
	v2139 = int32(1)
	v2141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2138<<(uint(v2139)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2129-int32(3)))) = uint16(v2141)
	v2150 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2136-v2138*v2137)<<(uint(v2139)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2129-v2139))) = uint16(v2150)
	v2153 = v2113 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v2112) {
		v2112 = v2133
		v2113 = v2153
		goto L373
	} else {
		goto L375
	}
L374:
	;
	v2158 = v2133
	v2159 = v2153
	goto L372
L375:
	;
	goto L374
L376:
	;
	v2200 = v2040 - int32(1)
	v2201 = v1232 + v2039
	if base.Ui32(int32(10)) <= base.Ui32(v2198) {
		goto L381
	} else {
		goto L382
	}
L377:
	;
	v2197 = v2159
	v2198 = v2158
	goto L376
L378:
	;
	goto L379
L379:
	;
	v2182 = int32(_a_F_halfvec_out_10)
	v2184 = int32(100)
	v2185 = base.I32_div_u_s(v2158&v2182, v2184)
	v2193 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2158-v2185*v2184)&v2182<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1232+v2039+v2093+(v2159^int32(-1))))) = uint16(v2193)
	v2197 = v2159 | int32(2)
	v2198 = v2185
	goto L376
L380:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2201))) = uint8(v2215)
	if base.Ui32(int32(2)) <= base.Ui32(v2093) {
		goto L384
	} else {
		goto L385
	}
L381:
	;
	v2208 = v2198 << (uint(int32(1)) % 32)
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2208)+uint32(_c_F_halfvec_out[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1232+(v2039+v2093-v2197)))) = uint8(v2209)
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2208)+uint32(_c_F_halfvec_out[4]))))
	v2215 = v2211
	goto L380
L382:
	;
	goto L383
L383:
	;
	v2215 = v2198 | int32(48)
	goto L380
L384:
	;
	v2219 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2201)+1)) = uint8(v2219)
	v2224 = v2093 + int32(1)
	goto L386
L385:
	;
	v2224 = int32(1)
	goto L386
L386:
	;
	v2225 = v2224 + v2039
	v2226 = v1232 + v2225
	v2227 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2226))) = uint8(v2227)
	v2232 = base.B2i32(v2200 < int32(0))
	if v2200 < int32(0) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v2233 = int32(45)
	goto L389
L388:
	;
	v2233 = int32(43)
	goto L389
L389:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2226)+1)) = uint8(v2233)
	if v2200 < int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v2237 = int32(1) - v2040
	goto L392
L391:
	;
	v2237 = v2200
	goto L392
L392:
	;
	v2242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2237<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2226)+2)) = uint16(v2242)
	v2423 = v2225 + int32(4)
	goto L260
L393:
	;
	v2258 = v2251
	v2259 = v2016
	goto L396
L394:
	;
	v2304 = v2251
	v2305 = v2016
	goto L395
L395:
	;
	if base.Ui32(v2305) < base.Ui32(int32(100)) {
		goto L400
	} else {
		goto L401
	}
L396:
	;
	v2275 = v2249 + v2045 + v2032 - v2258
	v2276 = int32(4)
	v2279 = base.I32_div_u_s(v2259, int32(_a_F_halfvec_out_9))
	v2282 = v2259 + v2279*int32(-10000)
	v2283 = int32(100)
	v2284 = base.I32_div_u_s(v2282, v2283)
	v2285 = int32(1)
	v2287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2284<<(uint(v2285)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2275-v2276))) = uint16(v2287)
	v2296 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2282-v2284*v2283)<<(uint(v2285)%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2275-int32(2)))) = uint16(v2296)
	v2299 = v2258 + v2276
	if base.Ui32(int32(99999999)) < base.Ui32(v2259) {
		v2258 = v2299
		v2259 = v2279
		goto L396
	} else {
		goto L398
	}
L397:
	;
	v2304 = v2299
	v2305 = v2279
	goto L395
L398:
	;
	goto L397
L399:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2344) {
		goto L404
	} else {
		goto L405
	}
L400:
	;
	v2343 = v2304
	v2344 = v2305
	goto L399
L401:
	;
	goto L402
L402:
	;
	v2326 = int32(2)
	v2328 = int32(_a_F_halfvec_out_10)
	v2330 = int32(100)
	v2331 = base.I32_div_u_s(v2305&v2328, v2330)
	v2339 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2305-v2331*v2330)&v2328<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2249+v2045+v2032-v2304-v2326))) = uint16(v2339)
	v2343 = v2304 | v2326
	v2344 = v2331
	goto L399
L403:
	;
	v2360 = int32(1)
	if v2249 == v2360 {
		goto L408
	} else {
		goto L409
	}
L404:
	;
	v2354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2344<<(uint(int32(1))%32))+uint32(_c_F_halfvec_out[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2249+v2045+v2032-v2343-int32(2)))) = uint16(v2354)
	goto L403
L405:
	;
	goto L406
L406:
	;
	v2358 = v2344 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2249+v2045))) = uint8(v2358)
	goto L403
L407:
	;
	v2423 = v2400 + int32(base.Ui32(v1344)>>(uint(int32(31))%32))
	goto L260
L408:
	;
	if v2040&int32(4) != 0 {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	goto L410
L410:
	;
	if v2021 < int32(0) {
		goto L420
	} else {
		goto L421
	}
L411:
	;
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v2045)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v2045))) = v2365
	v2368 = int32(5)
	goto L413
L412:
	;
	v2368 = v2360
	goto L413
L413:
	;
	if v2040&int32(2) != 0 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v2371 = v2368 + v2045
	v2374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2371-int32(1)))) = uint16(v2374)
	v2379 = v2368 | int32(2)
	goto L416
L415:
	;
	v2379 = v2368
	goto L416
L416:
	;
	if v2040&int32(1) != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2382 = v2379 + v2045
	v2385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2382))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2382-int32(1)))) = uint8(v2385)
	goto L419
L418:
	;
	goto L419
L419:
	;
	v2389 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2045+v2040))) = uint8(v2389)
	v2400 = v2032 + int32(1)
	goto L407
L420:
	;
	v2397 = int32(2) - v2021
	goto L422
L421:
	;
	v2397 = v2040
	goto L422
L422:
	;
	v2400 = v2397
	goto L407
L423:
	;
	goto L216
L424:
	;
	F_pfree(m, v10)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L1
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	return v20
L427:
	;
	goto L426
}
func F_halfvec_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_pq_begintypsend(m, v9)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	F_enlargeStringInfo(m, v9, int32(2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v25 = int32(8)
	v29 = v18<<(uint(v25)%32) | int32(base.Ui32(v18)>>(uint(v25)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v22+v23))) = uint16(v29)
	v31 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22 + v31
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)))
	F_enlargeStringInfo(m, v9, v31)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v41 = int32(8)
	v45 = v34<<(uint(v41)%32) | int32(base.Ui32(v34)>>(uint(v41)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v38+v39))) = uint16(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v38 + int32(2)
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if int32(0) < v50 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v56 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v93 << (uint(int32(2)) % 32)
	goto L13
L9:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(8)+v56<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, v9, int32(2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v72 = int32(8)
	v76 = v65<<(uint(v72)%32) | int32(base.Ui32(v65)>>(uint(v72)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v69+v70))) = uint16(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69 + int32(2)
	v82 = v56 + int32(1)
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if v82 < v83 {
		v56 = v82
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	m.G0 = v9 + int32(16)
	return v92
}
