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
	var v20 float32
	_ = v20
	var v21 float32
	_ = v21
	var v22 float32
	_ = v22
	var v27 int32
	_ = v27
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
	var v120 float32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 float32
	_ = v213
	var v215 float32
	_ = v215
	var v217 float32
	_ = v217
	var v219 float32
	_ = v219
	var v221 int32
	_ = v221
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	v4 = int32(0)
	v6 = float32(0)
	if l0 <= v4 {
		v237 = float64(0)
		v238 = float64(0)
	} else {
		v18 = v4
		v20 = v6
		v21 = v6
		v22 = v6
		for {
			v27 = v18 << (uint(int32(1)) % 32)
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v27))))
			v34 = v29 & int32(1023)
			v38 = v29 << (uint(int32(16)) % 32) & int32(-2147483648)
			v41 = int32(31)
			v42 = int32(base.Ui32(v29)>>(uint(int32(10))%32)) & v41
			if v42 != v41 {
				if v42 != 0 {
					v114 = v34
					v115 = v42<<(uint(int32(23))%32) + v38 + int32(939524096)
				} else {
					if v34 != 0 {
						if v29&int32(512) != 0 {
							v102 = v34 << (uint(int32(1)) % 32)
							v104 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v34) {
								v102 = v34 << (uint(int32(2)) % 32)
								v104 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v34) {
									v102 = v34 << (uint(int32(3)) % 32)
									v104 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v34) {
										v102 = v34 << (uint(int32(4)) % 32)
										v104 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v34) {
											v102 = v34 << (uint(int32(5)) % 32)
											v104 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v34) {
												v102 = v34 << (uint(int32(6)) % 32)
												v104 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v34) {
													v102 = v34 << (uint(int32(7)) % 32)
													v104 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v34) {
														v102 = v34 << (uint(int32(8)) % 32)
														v104 = int32(880803840)
													} else {
														v97 = base.B2i32(v34 == int32(1))
														if v34 == int32(1) {
															v98 = int32(1024)
														} else {
															v98 = v34 << (uint(int32(9)) % 32)
														}
														if v34 == int32(1) {
															v101 = int32(864026624)
														} else {
															v101 = int32(872415232)
														}
														v102 = v98
														v104 = v101
													}
												}
											}
										}
									}
								}
							}
						}
						v114 = v102 & int32(1022)
						v115 = v104 | v38
					} else {
						v114 = int32(0)
						v115 = v38
					}
				}
			} else {
				if v34 == int32(0) {
					v114 = int32(0)
					v115 = v38 | int32(2139095040)
				} else {
					v114 = v34
					v115 = v38 | int32(2143289344)
				}
			}
			v120 = base.F32_reinterpret_i32(v115 | v114<<(uint(int32(13))%32))
			v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v27))))
			v127 = v122 & int32(1023)
			v131 = v122 << (uint(int32(16)) % 32) & int32(-2147483648)
			v134 = int32(31)
			v135 = int32(base.Ui32(v122)>>(uint(int32(10))%32)) & v134
			if v135 != v134 {
				if v135 != 0 {
					v207 = v127
					v208 = v135<<(uint(int32(23))%32) + v131 + int32(939524096)
				} else {
					if v127 != 0 {
						if v122&int32(512) != 0 {
							v195 = v127 << (uint(int32(1)) % 32)
							v197 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v127) {
								v195 = v127 << (uint(int32(2)) % 32)
								v197 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v127) {
									v195 = v127 << (uint(int32(3)) % 32)
									v197 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v127) {
										v195 = v127 << (uint(int32(4)) % 32)
										v197 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v127) {
											v195 = v127 << (uint(int32(5)) % 32)
											v197 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v127) {
												v195 = v127 << (uint(int32(6)) % 32)
												v197 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v127) {
													v195 = v127 << (uint(int32(7)) % 32)
													v197 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v127) {
														v195 = v127 << (uint(int32(8)) % 32)
														v197 = int32(880803840)
													} else {
														v190 = base.B2i32(v127 == int32(1))
														if v127 == int32(1) {
															v191 = int32(1024)
														} else {
															v191 = v127 << (uint(int32(9)) % 32)
														}
														if v127 == int32(1) {
															v194 = int32(864026624)
														} else {
															v194 = int32(872415232)
														}
														v195 = v191
														v197 = v194
													}
												}
											}
										}
									}
								}
							}
						}
						v207 = v195 & int32(1022)
						v208 = v197 | v131
					} else {
						v207 = int32(0)
						v208 = v131
					}
				}
			} else {
				if v127 == int32(0) {
					v207 = int32(0)
					v208 = v131 | int32(2139095040)
				} else {
					v207 = v127
					v208 = v131 | int32(2143289344)
				}
			}
			v213 = base.F32_reinterpret_i32(v208 | v207<<(uint(int32(13))%32))
			v215 = base.F32_add(base.F32_mul(v120, v213), v20)
			v217 = base.F32_add(base.F32_mul(v120, v120), v21)
			v219 = base.F32_add(base.F32_mul(v213, v213), v22)
			v221 = v18 + int32(1)
			if v221 != l0 {
				v18 = v221
				v20 = v215
				v21 = v217
				v22 = v219
				continue
			} else {
				break
			}
			break
		}
		v237 = base.F64_mul(base.F64_promote_f32(v217), base.F64_promote_f32(v219))
		v238 = base.F64_promote_f32(v215)
	}
	return base.F64_div(v238, base.F64_sqrt(v237))
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
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 float32
	_ = v205
	var v207 float32
	_ = v207
	var v209 int32
	_ = v209
	var v216 float32
	_ = v216
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
			v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v18))))
			v118 = v113 & int32(1023)
			v122 = v113 << (uint(int32(16)) % 32) & int32(-2147483648)
			v125 = int32(31)
			v126 = int32(base.Ui32(v113)>>(uint(int32(10))%32)) & v125
			if v126 != v125 {
				if v126 != 0 {
					v198 = v118
					v199 = v126<<(uint(int32(23))%32) + v122 + int32(939524096)
				} else {
					if v118 != 0 {
						if v113&int32(512) != 0 {
							v186 = v118 << (uint(int32(1)) % 32)
							v188 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v118) {
								v186 = v118 << (uint(int32(2)) % 32)
								v188 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v118) {
									v186 = v118 << (uint(int32(3)) % 32)
									v188 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v118) {
										v186 = v118 << (uint(int32(4)) % 32)
										v188 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v118) {
											v186 = v118 << (uint(int32(5)) % 32)
											v188 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v118) {
												v186 = v118 << (uint(int32(6)) % 32)
												v188 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v118) {
													v186 = v118 << (uint(int32(7)) % 32)
													v188 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v118) {
														v186 = v118 << (uint(int32(8)) % 32)
														v188 = int32(880803840)
													} else {
														v181 = base.B2i32(v118 == int32(1))
														if v118 == int32(1) {
															v182 = int32(1024)
														} else {
															v182 = v118 << (uint(int32(9)) % 32)
														}
														if v118 == int32(1) {
															v185 = int32(864026624)
														} else {
															v185 = int32(872415232)
														}
														v186 = v182
														v188 = v185
													}
												}
											}
										}
									}
								}
							}
						}
						v198 = v186 & int32(1022)
						v199 = v188 | v122
					} else {
						v198 = int32(0)
						v199 = v122
					}
				}
			} else {
				if v118 == int32(0) {
					v198 = int32(0)
					v199 = v122 | int32(2139095040)
				} else {
					v198 = v118
					v199 = v122 | int32(2143289344)
				}
			}
			v205 = base.F32_sub(base.F32_reinterpret_i32(v106|v105<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v199|v198<<(uint(int32(13))%32)))
			v207 = base.F32_add(base.F32_mul(v205, v205), v15)
			v209 = v13 + int32(1)
			if v209 != l0 {
				v13 = v209
				v15 = v207
				continue
			} else {
				break
			}
			break
		}
		v216 = v207
	} else {
		v216 = v6
	}
	return v216
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
	var v59 int32
	_ = v59
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
	var v86 int32
	_ = v86
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
	return v86
L9:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v86 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v33 = int32(65535)
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
		v86 = v45
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v59 = int32(0)
	goto L17
L17:
	;
	v66 = int32(1)
	v70 = v59 + v66
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
	v86 = v45
	goto L8
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(8)+v59<<(uint(v66)%32)))) = uint16(v77)
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
		v59 = v70
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(342707)
	F_errmsg_internal(m, int32(26326), v11)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(523988), int32(173), int32(26817))
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
	var v134 float32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 float32
	_ = v227
	var v235 int32
	_ = v235
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
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+(v17+v25)))))
	v141 = v136 & int32(1023)
	v145 = v136 << (uint(int32(16)) % 32) & int32(-2147483648)
	v148 = int32(31)
	v149 = int32(base.Ui32(v136)>>(uint(int32(10))%32)) & v148
	if v149 != v148 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v134 = base.F32_reinterpret_i32(v129 | v128<<(uint(int32(13))%32))
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
	if base.F32_lt(v134, v227) != 0 {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v227 = base.F32_reinterpret_i32(v222 | v221<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v221 = v141
	v222 = v149<<(uint(int32(23))%32) + v145 + int32(939524096)
	goto L55
L57:
	;
	if v136&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v149 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v141 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v141 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v221 = int32(0)
	v222 = v145
	goto L55
L63:
	;
	v221 = int32(0)
	v222 = v145 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v221 = v141
	v222 = v145 | int32(2143289344)
	goto L55
L66:
	;
	v221 = v209 & int32(1022)
	v222 = v211 | v145
	goto L55
L67:
	;
	v209 = v141 << (uint(int32(1)) % 32)
	v211 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v141) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v209 = v141 << (uint(int32(2)) % 32)
	v211 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v141) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v209 = v141 << (uint(int32(3)) % 32)
	v211 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v141) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v209 = v141 << (uint(int32(4)) % 32)
	v211 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v141) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v209 = v141 << (uint(int32(5)) % 32)
	v211 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v141) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v209 = v141 << (uint(int32(6)) % 32)
	v211 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v141) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v209 = v141 << (uint(int32(7)) % 32)
	v211 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v141) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v209 = v141 << (uint(int32(8)) % 32)
	v211 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v204 = base.B2i32(v141 == int32(1))
	if v141 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v205 = int32(1024)
	goto L93
L92:
	;
	v205 = v141 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v141 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v208 = int32(864026624)
	goto L96
L95:
	;
	v208 = int32(872415232)
	goto L96
L96:
	;
	v209 = v205
	v211 = v208
	goto L66
L97:
	;
	return int32(-1)
L98:
	;
	goto L99
L99:
	;
	if base.F32_gt(v134, v227) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v235 = v30 + int32(1)
	if v235 == v22 {
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
	v30 = v235
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
	var v136 float32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 float32
	_ = v229
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
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
	return v250
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
		v250 = v2
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
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43+(v18+v26)))))
	v143 = v138 & int32(1023)
	v147 = v138 << (uint(int32(16)) % 32) & int32(-2147483648)
	v150 = int32(31)
	v151 = int32(base.Ui32(v138)>>(uint(int32(10))%32)) & v150
	if v151 != v150 {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	v136 = base.F32_reinterpret_i32(v131 | v130<<(uint(int32(13))%32))
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
	if base.F32_gt(v136, v229)|base.F32_lt(v136, v229) != 0 {
		v250 = v2
		goto L4
	} else {
		goto L99
	}
L57:
	;
	v229 = base.F32_reinterpret_i32(v224 | v223<<(uint(int32(13))%32))
	goto L56
L58:
	;
	v223 = v143
	v224 = v151<<(uint(int32(23))%32) + v147 + int32(939524096)
	goto L57
L59:
	;
	if v138&int32(512) != 0 {
		goto L69
	} else {
		goto L70
	}
L60:
	;
	if v151 != 0 {
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v143 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if v143 != 0 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v223 = int32(0)
	v224 = v147
	goto L57
L65:
	;
	v223 = int32(0)
	v224 = v147 | int32(2139095040)
	goto L57
L66:
	;
	goto L67
L67:
	;
	v223 = v143
	v224 = v147 | int32(2143289344)
	goto L57
L68:
	;
	v223 = v211 & int32(1022)
	v224 = v213 | v147
	goto L57
L69:
	;
	v211 = v143 << (uint(int32(1)) % 32)
	v213 = int32(939524096)
	goto L68
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(int32(255)) < base.Ui32(v143) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v211 = v143 << (uint(int32(2)) % 32)
	v213 = int32(931135488)
	goto L68
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(127)) < base.Ui32(v143) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v211 = v143 << (uint(int32(3)) % 32)
	v213 = int32(922746880)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if base.Ui32(int32(63)) < base.Ui32(v143) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v211 = v143 << (uint(int32(4)) % 32)
	v213 = int32(914358272)
	goto L68
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(int32(31)) < base.Ui32(v143) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v211 = v143 << (uint(int32(5)) % 32)
	v213 = int32(905969664)
	goto L68
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(int32(15)) < base.Ui32(v143) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v211 = v143 << (uint(int32(6)) % 32)
	v213 = int32(897581056)
	goto L68
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(int32(7)) < base.Ui32(v143) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v211 = v143 << (uint(int32(7)) % 32)
	v213 = int32(889192448)
	goto L68
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(int32(3)) < base.Ui32(v143) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v211 = v143 << (uint(int32(8)) % 32)
	v213 = int32(880803840)
	goto L68
L91:
	;
	goto L92
L92:
	;
	v206 = base.B2i32(v143 == int32(1))
	if v143 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v207 = int32(1024)
	goto L95
L94:
	;
	v207 = v143 << (uint(int32(9)) % 32)
	goto L95
L95:
	;
	if v143 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v210 = int32(864026624)
	goto L98
L97:
	;
	v210 = int32(872415232)
	goto L98
L98:
	;
	v211 = v207
	v213 = v210
	goto L68
L99:
	;
	v234 = v31 + int32(1)
	if v234 != v23 {
		v31 = v234
		goto L11
	} else {
		goto L100
	}
L100:
	;
	goto L12
L101:
	;
	v250 = base.B2i32(v20 <= v21)
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
	var v134 float32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 float32
	_ = v227
	var v233 int32
	_ = v233
	var v252 int32
	_ = v252
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
	return v252
L5:
	;
	v252 = base.B2i32(v20 <= v19)
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
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+(v17+v25)))))
	v141 = v136 & int32(1023)
	v145 = v136 << (uint(int32(16)) % 32) & int32(-2147483648)
	v148 = int32(31)
	v149 = int32(base.Ui32(v136)>>(uint(int32(10))%32)) & v148
	if v149 != v148 {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	v134 = base.F32_reinterpret_i32(v129 | v128<<(uint(int32(13))%32))
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
	if base.F32_lt(v134, v227) != 0 {
		v252 = int32(0)
		goto L4
	} else {
		goto L98
	}
L56:
	;
	v227 = base.F32_reinterpret_i32(v222 | v221<<(uint(int32(13))%32))
	goto L55
L57:
	;
	v221 = v141
	v222 = v149<<(uint(int32(23))%32) + v145 + int32(939524096)
	goto L56
L58:
	;
	if v136&int32(512) != 0 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	if v149 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v141 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v141 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v221 = int32(0)
	v222 = v145
	goto L56
L64:
	;
	v221 = int32(0)
	v222 = v145 | int32(2139095040)
	goto L56
L65:
	;
	goto L66
L66:
	;
	v221 = v141
	v222 = v145 | int32(2143289344)
	goto L56
L67:
	;
	v221 = v209 & int32(1022)
	v222 = v211 | v145
	goto L56
L68:
	;
	v209 = v141 << (uint(int32(1)) % 32)
	v211 = int32(939524096)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(int32(255)) < base.Ui32(v141) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v209 = v141 << (uint(int32(2)) % 32)
	v211 = int32(931135488)
	goto L67
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(127)) < base.Ui32(v141) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v209 = v141 << (uint(int32(3)) % 32)
	v211 = int32(922746880)
	goto L67
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(63)) < base.Ui32(v141) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v209 = v141 << (uint(int32(4)) % 32)
	v211 = int32(914358272)
	goto L67
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(31)) < base.Ui32(v141) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v209 = v141 << (uint(int32(5)) % 32)
	v211 = int32(905969664)
	goto L67
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v141) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v209 = v141 << (uint(int32(6)) % 32)
	v211 = int32(897581056)
	goto L67
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(7)) < base.Ui32(v141) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v209 = v141 << (uint(int32(7)) % 32)
	v211 = int32(889192448)
	goto L67
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(3)) < base.Ui32(v141) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v209 = v141 << (uint(int32(8)) % 32)
	v211 = int32(880803840)
	goto L67
L90:
	;
	goto L91
L91:
	;
	v204 = base.B2i32(v141 == int32(1))
	if v141 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v205 = int32(1024)
	goto L94
L93:
	;
	v205 = v141 << (uint(int32(9)) % 32)
	goto L94
L94:
	;
	if v141 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v208 = int32(864026624)
	goto L97
L96:
	;
	v208 = int32(872415232)
	goto L97
L97:
	;
	v209 = v205
	v211 = v208
	goto L67
L98:
	;
	if base.F32_gt(v134, v227) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v233 = v30 + int32(1)
	if v233 == v22 {
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
	v30 = v233
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
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 float32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 float32
	_ = v277
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
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
	if base.F64_gt(v150, float64(0)) == int32(0) {
		goto L6
	} else {
		goto L54
	}
L10:
	;
	v148 = base.F64_promote_f32(base.F32_reinterpret_i32(v142 | v141<<(uint(int32(13))%32)))
	v150 = base.F64_add(base.F64_mul(v148, v148), v50)
	v152 = v40 + int32(1)
	if v152 != v33 {
		v40 = v152
		v50 = v150
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
	v160 = int32(0)
	goto L55
L55:
	;
	v174 = v160 << (uint(int32(1)) % 32)
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+v174))))
	v181 = v176 & int32(1023)
	v185 = v176 << (uint(int32(16)) % 32) & int32(-2147483648)
	v188 = int32(31)
	v189 = int32(base.Ui32(v176)>>(uint(int32(10))%32)) & v188
	if v189 != v188 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v362 = int32(0)
	if v360 <= v362 {
		goto L6
	} else {
		goto L124
	}
L57:
	;
	v270 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(base.F32_reinterpret_i32(v262|v261<<(uint(int32(13))%32))), base.F64_sqrt(v150)))
	v271 = base.I32_reinterpret_f32(v270)
	v273 = int32(base.Ui32(v271) >> (uint(int32(16)) % 32))
	v277 = base.F32_abs(v270)
	if base.F32_eq(v277, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v356 = v273 & int32(64512)
		goto L100
	} else {
		goto L101
	}
L58:
	;
	goto L57
L59:
	;
	v261 = v181
	v262 = v189<<(uint(int32(23))%32) + v185 + int32(939524096)
	goto L58
L60:
	;
	if v176&int32(512) != 0 {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	if v189 != 0 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v181 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v181 != 0 {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v261 = int32(0)
	v262 = v185
	goto L58
L66:
	;
	v261 = int32(0)
	v262 = v185 | int32(2139095040)
	goto L58
L67:
	;
	goto L68
L68:
	;
	v261 = v181
	v262 = v185 | int32(2143289344)
	goto L58
L69:
	;
	v261 = v249 & int32(1022)
	v262 = v251 | v185
	goto L58
L70:
	;
	v249 = v181 << (uint(int32(1)) % 32)
	v251 = int32(939524096)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(255)) < base.Ui32(v181) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v249 = v181 << (uint(int32(2)) % 32)
	v251 = int32(931135488)
	goto L69
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(127)) < base.Ui32(v181) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v249 = v181 << (uint(int32(3)) % 32)
	v251 = int32(922746880)
	goto L69
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(63)) < base.Ui32(v181) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v249 = v181 << (uint(int32(4)) % 32)
	v251 = int32(914358272)
	goto L69
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(31)) < base.Ui32(v181) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v249 = v181 << (uint(int32(5)) % 32)
	v251 = int32(905969664)
	goto L69
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(15)) < base.Ui32(v181) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v249 = v181 << (uint(int32(6)) % 32)
	v251 = int32(897581056)
	goto L69
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(7)) < base.Ui32(v181) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v249 = v181 << (uint(int32(7)) % 32)
	v251 = int32(889192448)
	goto L69
L89:
	;
	goto L90
L90:
	;
	if base.Ui32(int32(3)) < base.Ui32(v181) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v249 = v181 << (uint(int32(8)) % 32)
	v251 = int32(880803840)
	goto L69
L92:
	;
	goto L93
L93:
	;
	v244 = base.B2i32(v181 == int32(1))
	if v181 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v245 = int32(1024)
	goto L96
L95:
	;
	v245 = v181 << (uint(int32(9)) % 32)
	goto L96
L96:
	;
	if v181 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v248 = int32(864026624)
	goto L99
L98:
	;
	v248 = int32(872415232)
	goto L99
L99:
	;
	v249 = v245
	v251 = v248
	goto L69
L100:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v174+v39))) = uint16(v356)
	v359 = v160 + int32(1)
	v360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v359 < v360 {
		v160 = v359
		goto L55
	} else {
		goto L123
	}
L101:
	;
	v281 = v271 & int32(8388607)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v277)) {
		v356 = v273&int32(32768) | int32(base.Ui32(v281)>>(uint(int32(13))%32)) | int32(32256)
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v293 = v273 & int32(32768)
	v297 = int32(base.Ui32(v271)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v297) < base.Ui32(int32(99)) {
		v356 = v293
		goto L100
	} else {
		goto L103
	}
L103:
	;
	if base.Ui32(v297) <= base.Ui32(int32(112)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v309 = int32(1)<<(uint(v297-int32(90))%32) + int32(base.Ui32(v281)>>(uint(int32(113)-v297)%32))
	v311 = v309
	v312 = v309 | v271
	goto L106
L105:
	;
	v311 = v281
	v312 = v271
	goto L106
L106:
	;
	v314 = int32(base.Ui32(v311) >> (uint(int32(13)) % 32))
	v317 = int32(3)
	v318 = int32(base.Ui32(v311)>>(uint(int32(12))%32)) & v317
	if v318 != v317 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v335 = base.B2i32(v329 == int32(1024))
	if v329 == int32(1024) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	if v318 != int32(1) {
		v329 = v314
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v329 = v314 + int32(1)
	goto L107
L111:
	;
	if v312&int32(4095) == int32(0) {
		v329 = v314
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v336 = int32(-126)
	goto L115
L114:
	;
	v336 = int32(-127)
	goto L115
L115:
	;
	v337 = v336 + v297
	if int32(16) <= v337 {
		v356 = v293 | int32(31744)
		goto L100
	} else {
		goto L116
	}
L116:
	;
	if int32(-15) < v337 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v347 = v337<<(uint(int32(10))%32) + int32(15360) | v293
	goto L119
L118:
	;
	v347 = v293
	goto L119
L119:
	;
	if v329 == int32(1024) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v349 = int32(0)
	goto L122
L121:
	;
	v349 = v329
	goto L122
L122:
	;
	v356 = v347 | v349
	goto L100
L123:
	;
	goto L56
L124:
	;
	v365 = v362
	goto L125
L125:
	;
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v365<<(uint(int32(1))%32)))))
	if v381&int32(32767) != int32(31744) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L131
	}
L127:
	;
	v387 = v365 + int32(1)
	if v360 != v387 {
		v365 = v387
		goto L125
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	goto L126
L130:
	;
	goto L6
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_l2_squared_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 float32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
			if v17 != v18 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
						v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
						F_errmsg(m, int32(499586), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523988), int32(80), int32(159834))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v40 = int32(8)
				v45 = *(*int32)(unsafe.Add(mBase, _consts[1145]))
				v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v17), v10+v40, v15+v40)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = F_Float8GetDatum(m, base.F64_promote_f32(v46))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v49
					}
				}
			}
		}
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
	var v226 float32
	_ = v226
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
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
	return v247
L5:
	;
	v247 = v20
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
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+(v16+v24)))))
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
	v133 = base.F32_reinterpret_i32(v128 | v127<<(uint(int32(13))%32))
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
	if base.F32_lt(v133, v226) != 0 {
		v247 = v38
		goto L4
	} else {
		goto L98
	}
L56:
	;
	v226 = base.F32_reinterpret_i32(v221 | v220<<(uint(int32(13))%32))
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
	if base.F32_gt(v133, v226) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v232 = v29 + int32(1)
	if v232 == v21 {
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
	v29 = v232
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
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v304 int64
	_ = v304
	var v308 int64
	_ = v308
	var v310 int32
	_ = v310
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v444 int64
	_ = v444
	var v446 int32
	_ = v446
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v451 int32
	_ = v451
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v489 int32
	_ = v489
	var v490 int64
	_ = v490
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v516 int64
	_ = v516
	var v520 int64
	_ = v520
	var v522 int32
	_ = v522
	var v525 int64
	_ = v525
	var v527 int32
	_ = v527
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int64
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1455 int64
	_ = v1455
	var v1457 int64
	_ = v1457
	var v1458 int64
	_ = v1458
	var v1460 int64
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int64
	_ = v1464
	var v1465 int64
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int64
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1491 int64
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1501 int32
	_ = v1501
	var v1502 int64
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1507 int64
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1535 int64
	_ = v1535
	var v1539 int64
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1544 int64
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1557 int32
	_ = v1557
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int64
	_ = v1670
	var v1672 int64
	_ = v1672
	var v1673 int64
	_ = v1673
	var v1675 int64
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1679 int64
	_ = v1679
	var v1680 int64
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1700 int64
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1705 int64
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1715 int32
	_ = v1715
	var v1716 int64
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1721 int64
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1747 int64
	_ = v1747
	var v1751 int64
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1756 int64
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1771 int32
	_ = v1771
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1786 int32
	_ = v1786
	var v1792 int32
	_ = v1792
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1962 int32
	_ = v1962
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2003 int32
	_ = v2003
	var v2013 int32
	_ = v2013
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2055 int32
	_ = v2055
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2112 int32
	_ = v2112
	var v2119 int32
	_ = v2119
	var v2125 int32
	_ = v2125
	var v2135 int32
	_ = v2135
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2229 int32
	_ = v2229
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2241 int32
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2287 int64
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2329 int32
	_ = v2329
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2422 int32
	_ = v2422
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
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
		v2477 = v26
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v2483 = int32(93)
	*(*uint16)(unsafe.Add(mBase, uint32(v2477))) = uint16(v2483)
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2485 != v10 {
		goto L432
	} else {
		goto L433
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
	v121 = int32(0)
	v138 = base.I32_reinterpret_f32(base.F32_reinterpret_i32(v115 | v114<<(uint(int32(13))%32)))
	v140 = v138 & int32(8388607)
	v143 = int32(255)
	v144 = int32(base.Ui32(v138)>>(uint(int32(23))%32)) & v143
	if v140|v144 != 0 {
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
	v1241 = v1240 + v26
	if v14 == int32(1) {
		v2477 = v1241
		goto L6
	} else {
		goto L218
	}
L52:
	;
	v149 = base.B2i32(v144 != v143)
	goto L54
L53:
	;
	v149 = v121
	goto L54
L54:
	;
	if v149 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v140 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(int32(23)) < base.Ui32(v144-int32(127)) {
		goto L75
	} else {
		goto L76
	}
L58:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1139])))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2)) = uint8(v153)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1140])))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v156)
	v1240 = int32(3)
	goto L51
L59:
	;
	goto L60
L60:
	;
	if v138 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v161 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v161)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v165 = v26 + int32(base.Ui32(v138)>>(uint(int32(31))%32))
	if v144 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(8751735898823355977)
	if v138 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v173 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v173)
	if v138 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v172 = int32(9)
	goto L69
L68:
	;
	v172 = int32(8)
	goto L69
L69:
	;
	v1240 = v172
	goto L51
L70:
	;
	v179 = int32(2)
	goto L72
L71:
	;
	v179 = int32(1)
	goto L72
L72:
	;
	v1240 = v179
	goto L51
L73:
	;
	v832 = v831 + v824
	v833 = int32(0)
	if v138 < v833 {
		goto L143
	} else {
		goto L144
	}
L74:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v772) {
		v814 = v772
		v824 = v782
		v831 = int32(8)
		goto L73
	} else {
		goto L134
	}
L75:
	;
	v197 = int32(2)
	v203 = v140 << (uint(v197) % 32)
	if v144 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v184 = int32(-1)
	v186 = int32(150) - v144
	if v140&(v184<<(uint(v186)%32)^v184) != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v772 = int32(base.Ui32(v140|int32(8388608)) >> (uint(v186) % 32))
	v782 = v121
	goto L74
L78:
	;
	v206 = v203 | int32(33554432)
	goto L80
L79:
	;
	v206 = v203
	goto L80
L80:
	;
	v207 = base.B2i32(v140 != int32(0)) | base.B2i32(base.Ui32(v144) < base.Ui32(v197)) ^ int32(-1) + v206
	v209 = v206 | int32(2)
	if v144 != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v765 = v747 + v753
	v766 = v764 + v749
	if base.Ui32(v766) <= base.Ui32(int32(99999999)) {
		v772 = v766
		v782 = v765
		goto L74
	} else {
		goto L133
	}
L82:
	;
	v683 = int32(0)
	v684 = int32(10)
	v685 = base.I32_div_u_s(v673, v684)
	v687 = base.I32_div_u_s(v675, v684)
	if base.Ui32(v687) < base.Ui32(v685) {
		goto L127
	} else {
		goto L128
	}
L83:
	;
	v587 = int32(0)
	v588 = int32(10)
	v589 = base.I32_div_u_s(v577, v588)
	v591 = base.I32_div_u_s(v579, v588)
	if base.Ui32(v589) <= base.Ui32(v591) {
		goto L121
	} else {
		goto L122
	}
L84:
	;
	v213 = v144 - int32(152)
	goto L86
L85:
	;
	v213 = int32(-151)
	goto L86
L86:
	;
	if int32(0) <= v213 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v219 = int32(base.Ui32(v213*int32(78913)) >> (uint(int32(18)) % 32))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v219<<(uint(int32(3))%32))+uint32(_consts[1141])))
	v226 = v224 & int64(4294967295)
	v227 = base.I64_extend_i32_u(v207)
	v229 = int64(32)
	v231 = base.I32_wrap_i64(int64(base.Ui64(v226*v227) >> (uint(v229) % 64)))
	v233 = int64(base.Ui64(v224) >> (uint(v229) % 64))
	v234 = v233 * v227
	v236 = v231 + base.I32_wrap_i64(v234)
	v243 = v219 - v213
	v248 = v243 + int32(base.Ui32(v219*int32(1217359))>>(uint(int32(19))%32))
	v249 = int32(5) - v248
	v252 = v248 + int32(27)
	v254 = (base.B2i32(base.Ui32(v236) < base.Ui32(v231))+base.I32_wrap_i64(int64(base.Ui64(v234)>>(uint(v229)%64))))<<(uint(v249)%32) | int32(base.Ui32(v236)>>(uint(v252)%32))
	v255 = base.I64_extend_i32_u(v209)
	v259 = base.I32_wrap_i64(int64(base.Ui64(v226*v255) >> (uint(v229) % 64)))
	v260 = v255 * v233
	v262 = v259 + base.I32_wrap_i64(v260)
	v270 = (base.B2i32(base.Ui32(v262) < base.Ui32(v259))+base.I32_wrap_i64(int64(base.Ui64(v260)>>(uint(v229)%64))))<<(uint(v249)%32) | int32(base.Ui32(v262)>>(uint(v252)%32))
	v271 = base.I64_extend_i32_u(v206)
	v275 = base.I32_wrap_i64(int64(base.Ui64(v226*v271) >> (uint(v229) % 64)))
	v276 = v271 * v233
	v278 = v275 + base.I32_wrap_i64(v276)
	v286 = (base.B2i32(base.Ui32(v278) < base.Ui32(v275))+base.I32_wrap_i64(int64(base.Ui64(v276)>>(uint(v229)%64))))<<(uint(v249)%32) | int32(base.Ui32(v278)>>(uint(v252)%32))
	v287 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v213) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v432 = v213 * int32(-732923)
	v434 = int32(base.Ui32(v432) >> (uint(int32(20)) % 32))
	v435 = v434 + v213
	v439 = *(*int64)(unsafe.Add(mBase, uint32(int32(1876512)-v435<<(uint(int32(3))%32))))
	v441 = v439 & int64(4294967295)
	v442 = base.I64_extend_i32_u(v207)
	v444 = int64(32)
	v446 = base.I32_wrap_i64(int64(base.Ui64(v441*v442) >> (uint(v444) % 64)))
	v448 = int64(base.Ui64(v439) >> (uint(v444) % 64))
	v449 = v448 * v442
	v451 = v446 + base.I32_wrap_i64(v449)
	v462 = v434 - int32(base.Ui32(v435*int32(-1217359))>>(uint(int32(19))%32))
	v463 = int32(4) - v462
	v466 = v462 + int32(28)
	v468 = (base.B2i32(base.Ui32(v451) < base.Ui32(v446))+base.I32_wrap_i64(int64(base.Ui64(v449)>>(uint(v444)%64))))<<(uint(v463)%32) | int32(base.Ui32(v451)>>(uint(v466)%32))
	v469 = base.I64_extend_i32_u(v206)
	v473 = base.I32_wrap_i64(int64(base.Ui64(v441*v469) >> (uint(v444) % 64)))
	v474 = v469 * v448
	v476 = v473 + base.I32_wrap_i64(v474)
	v484 = (base.B2i32(base.Ui32(v476) < base.Ui32(v473))+base.I32_wrap_i64(int64(base.Ui64(v474)>>(uint(v444)%64))))<<(uint(v463)%32) | int32(base.Ui32(v476)>>(uint(v466)%32))
	v485 = base.I64_extend_i32_u(v209)
	v489 = base.I32_wrap_i64(int64(base.Ui64(v441*v485) >> (uint(v444) % 64)))
	v490 = v448 * v485
	v492 = v489 + base.I32_wrap_i64(v490)
	v500 = (base.B2i32(base.Ui32(v492) < base.Ui32(v489))+base.I32_wrap_i64(int64(base.Ui64(v490)>>(uint(v444)%64))))<<(uint(v463)%32) | int32(base.Ui32(v492)>>(uint(v466)%32))
	v502 = v500 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v432) {
		goto L111
	} else {
		goto L112
	}
L90:
	;
	v293 = int32(10)
	v294 = base.I32_div_u_s(v270-int32(1), v293)
	v296 = base.I32_div_u_s(v254, v293)
	if base.Ui32(v294) <= base.Ui32(v296) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v341 = v287
	goto L92
L92:
	;
	v346 = base.I32_rem_u_s(v206, int32(5))
	if v346 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v299 = v219 - int32(1)
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v299<<(uint(int32(3))%32))+uint32(_consts[1141])))
	v308 = int64(32)
	v310 = base.I32_wrap_i64(int64(base.Ui64(v304&int64(4294967295)*v271) >> (uint(v308) % 64)))
	v313 = int64(base.Ui64(v304)>>(uint(v308)%64)) * v271
	v315 = v310 + base.I32_wrap_i64(v313)
	v326 = v243 + int32(base.Ui32(v299*int32(1217359))>>(uint(int32(19))%32))
	v334 = base.I32_rem_u_s((base.B2i32(base.Ui32(v315) < base.Ui32(v310))+base.I32_wrap_i64(int64(base.Ui64(v313)>>(uint(v308)%64))))<<(uint(int32(6)-v326)%32)|int32(base.Ui32(v315)>>(uint(v326+int32(26))%32)), int32(10))
	v335 = v334
	goto L95
L94:
	;
	v335 = v287
	goto L95
L95:
	;
	if base.Ui32(int32(33)) < base.Ui32(v213) {
		v667 = v286
		v669 = v335
		v672 = v219
		v673 = v270
		v675 = v254
		goto L82
	} else {
		goto L96
	}
L96:
	;
	v341 = v335
	goto L92
L97:
	;
	v351 = v206
	v353 = v287
	goto L100
L98:
	;
	goto L99
L99:
	;
	v377 = int32(0)
	v379 = base.I32_rem_u_s(v209, int32(5))
	if v379 == v377 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v369 = v353 + int32(1)
	v370 = int32(5)
	v371 = base.I32_div_u_s(v351, v370)
	v373 = base.I32_rem_u_s(v371, v370)
	if v373 == int32(0) {
		v351 = v371
		v353 = v369
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if base.Ui32(v369) < base.Ui32(v219) {
		v667 = v286
		v669 = v341
		v672 = v219
		v673 = v270
		v675 = v254
		goto L82
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v571 = v286
	v573 = v341
	v576 = v219
	v577 = v270
	v579 = v254
	goto L83
L104:
	;
	v384 = v377
	v388 = v209
	goto L107
L105:
	;
	v411 = v377
	goto L106
L106:
	;
	v667 = v286
	v669 = v341
	v672 = v219
	v673 = v270 - base.B2i32(base.Ui32(v219) <= base.Ui32(v411))
	v675 = v254
	goto L82
L107:
	;
	v402 = v384 + int32(1)
	v403 = int32(5)
	v404 = base.I32_div_u_s(v388, v403)
	v406 = base.I32_rem_u_s(v404, v403)
	if v406 == int32(0) {
		v384 = v402
		v388 = v404
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v411 = v402
	goto L106
L109:
	;
	goto L108
L110:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v432) {
		v667 = v484
		v669 = v549
		v672 = v435
		v673 = v500
		v675 = v468
		goto L82
	} else {
		goto L118
	}
L111:
	;
	v505 = int32(10)
	v506 = base.I32_div_u_s(v502, v505)
	v508 = base.I32_div_u_s(v468, v505)
	if base.Ui32(v506) <= base.Ui32(v508) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v555 = v121
	goto L113
L113:
	;
	v571 = v484
	v573 = v555
	v576 = v435
	v577 = v502
	v579 = v468
	goto L83
L114:
	;
	v511 = int32(1) - v435
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v511<<(uint(int32(3))%32))+uint32(_consts[1142])))
	v520 = int64(32)
	v522 = base.I32_wrap_i64(int64(base.Ui64(v516&int64(4294967295)*v469) >> (uint(v520) % 64)))
	v525 = int64(base.Ui64(v516)>>(uint(v520)%64)) * v469
	v527 = v522 + base.I32_wrap_i64(v525)
	v540 = v434 + (int32(base.Ui32(v511*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v548 = base.I32_rem_u_s((base.B2i32(base.Ui32(v527) < base.Ui32(v522))+base.I32_wrap_i64(int64(base.Ui64(v525)>>(uint(v520)%64))))<<(uint(int32(4)-v540)%32)|int32(base.Ui32(v527)>>(uint(v540+int32(28))%32)), int32(10))
	v549 = v548
	goto L116
L115:
	;
	v549 = v121
	goto L116
L116:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v432) {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v555 = v549
	goto L113
L118:
	;
	v561 = int32(-1)
	if v206&(v561<<(uint(v434-int32(1))%32)^v561) != 0 {
		v667 = v484
		v669 = v549
		v672 = v435
		v673 = v500
		v675 = v468
		goto L82
	} else {
		goto L119
	}
L119:
	;
	v571 = v484
	v573 = v549
	v576 = v435
	v577 = v500
	v579 = v468
	goto L83
L120:
	;
	v654 = v639 & int32(255)
	v747 = v635
	v749 = v637
	v753 = v576
	v764 = (v652|base.B2i32(v654 != int32(5))|v637)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v654)) | base.B2i32(v637 == v644)
	goto L81
L121:
	;
	v635 = v587
	v637 = v571
	v639 = v573
	v644 = v579
	v652 = int32(0)
	goto L120
L122:
	;
	goto L123
L123:
	;
	v597 = v587
	v598 = v571
	v600 = v573
	v602 = v589
	v604 = v591
	v605 = int32(1)
	goto L124
L124:
	;
	v615 = v597 + int32(1)
	v616 = int32(10)
	v617 = base.I32_div_u_s(v598, v616)
	v620 = v598 - v617*v616
	v625 = v605 & base.B2i32(v600&int32(255) == int32(0))
	v627 = base.I32_div_u_s(v602, v616)
	v629 = base.I32_div_u_s(v604, v616)
	if base.Ui32(v629) < base.Ui32(v627) {
		v597 = v615
		v598 = v617
		v600 = v620
		v602 = v627
		v604 = v629
		v605 = v625
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v635 = v615
	v637 = v617
	v639 = v620
	v644 = v604
	v652 = v625 ^ int32(1)
	goto L120
L126:
	;
	goto L125
L127:
	;
	v691 = v683
	v692 = v667
	v693 = v685
	v695 = v687
	goto L130
L128:
	;
	v722 = v683
	v723 = v667
	v725 = v669
	v731 = v675
	goto L129
L129:
	;
	v747 = v722
	v749 = v723
	v753 = v672
	v764 = base.B2i32(v723 == v731) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v725&int32(255)))
	goto L81
L130:
	;
	v709 = v691 + int32(1)
	v710 = int32(10)
	v711 = base.I32_div_u_s(v692, v710)
	v713 = base.I32_div_u_s(v693, v710)
	v715 = base.I32_div_u_s(v695, v710)
	if base.Ui32(v715) < base.Ui32(v713) {
		v691 = v709
		v692 = v711
		v693 = v713
		v695 = v715
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v722 = v709
	v723 = v711
	v725 = v692 - v711*int32(10)
	v731 = v695
	goto L129
L132:
	;
	goto L131
L133:
	;
	v814 = v766
	v824 = v765
	v831 = int32(9)
	goto L73
L134:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v772) {
		v814 = v772
		v824 = v782
		v831 = int32(7)
		goto L73
	} else {
		goto L135
	}
L135:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v772) {
		v814 = v772
		v824 = v782
		v831 = int32(6)
		goto L73
	} else {
		goto L136
	}
L136:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v772) {
		v814 = v772
		v824 = v782
		v831 = int32(5)
		goto L73
	} else {
		goto L137
	}
L137:
	;
	if base.Ui32(int32(999)) < base.Ui32(v772) {
		v814 = v772
		v824 = v782
		v831 = int32(4)
		goto L73
	} else {
		goto L138
	}
L138:
	;
	if base.Ui32(int32(99)) < base.Ui32(v772) {
		v814 = v772
		v824 = v782
		v831 = int32(3)
		goto L73
	} else {
		goto L139
	}
L139:
	;
	if base.Ui32(int32(9)) < base.Ui32(v772) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v811 = int32(2)
	goto L142
L141:
	;
	v811 = int32(1)
	goto L142
L142:
	;
	v814 = v772
	v824 = v782
	v831 = v811
	goto L73
L143:
	;
	v836 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v836)
	v839 = int32(1)
	goto L145
L144:
	;
	v839 = v833
	goto L145
L145:
	;
	if base.Ui32(v832+int32(3)) <= base.Ui32(int32(9)) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v1060 = int32(0)
	if base.Ui32(v814) < base.Ui32(int32(10000)) {
		goto L188
	} else {
		goto L189
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v844))) = v1056
	v1058 = v1055
	goto L146
L148:
	;
	v844 = v26 + v839
	v845 = int32(0)
	if v832 <= v845 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	if v824 != 0 {
		goto L156
	} else {
		goto L157
	}
L151:
	;
	v1055 = int32(2) - v832
	v1056 = int64(3472328296227679792)
	goto L147
L152:
	;
	goto L153
L153:
	;
	if int32(0) <= v824 {
		v1055 = v845
		v1056 = int64(3472328296227680304)
		goto L147
	} else {
		goto L154
	}
L154:
	;
	v1058 = int32(1)
	goto L146
L155:
	;
	v904 = int32(0)
	if base.Ui32(v888) < base.Ui32(int32(10000)) {
		goto L164
	} else {
		goto L165
	}
L156:
	;
	v888 = v814
	v894 = v831
	goto L155
L157:
	;
	goto L158
L158:
	;
	v857 = v814
	v862 = v831
	goto L159
L159:
	;
	if v857&int32(1) != 0 {
		v888 = v857
		v894 = v862
		goto L155
	} else {
		goto L161
	}
L160:
	;
	v888 = v857
	v894 = v862
	goto L155
L161:
	;
	v881 = base.I32_div_u_s(v857, int32(10))
	if int32(0)-v857 == v881*int32(-10) {
		v857 = v881
		v862 = v862 - int32(1)
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	if base.Ui32(v963) < base.Ui32(int32(100)) {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	v961 = v904
	v963 = v888
	goto L163
L165:
	;
	goto L166
L166:
	;
	v911 = v904
	v912 = v888
	goto L167
L167:
	;
	v928 = v26 + v839 + v894 - v911
	v932 = base.I32_div_u_s(v912, int32(10000))
	v935 = v932*int32(-10000) + v912
	v936 = int32(100)
	v937 = base.I32_div_u_s(v935, v936)
	v938 = int32(1)
	v942 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937<<(uint(v938)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v928-int32(3)))) = uint16(v942)
	v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v935-v937*v936)<<(uint(v938)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v928-v938))) = uint16(v953)
	v956 = v911 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v912) {
		v911 = v956
		v912 = v932
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v961 = v956
	v963 = v932
	goto L163
L169:
	;
	goto L168
L170:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1003) {
		goto L175
	} else {
		goto L176
	}
L171:
	;
	v1002 = v961
	v1003 = v963
	goto L170
L172:
	;
	goto L173
L173:
	;
	v985 = int32(65535)
	v987 = int32(100)
	v988 = base.I32_div_u_s(v963&v985, v987)
	v998 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v963-v988*v987)&v985<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v26+v839+v894+(v961^int32(-1))))) = uint16(v998)
	v1002 = v961 | int32(2)
	v1003 = v988
	goto L170
L174:
	;
	v1022 = int32(1)
	v1023 = v832 - v1022
	v1024 = v26 + v839
	*(*uint8)(unsafe.Add(mBase, uint32(v1024))) = uint8(v1021)
	if base.Ui32(int32(2)) <= base.Ui32(v894) {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v1010 = v1003 << (uint(int32(1)) % 32)
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010)+uint32(_consts[1144]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+(v839+v894-v1002)))) = uint8(v1013)
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010)+uint32(_consts[1143]))))
	v1021 = v1017
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1021 = v1003 | int32(48)
	goto L174
L178:
	;
	v1029 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1024)+1)) = uint8(v1029)
	v1033 = v894 + int32(1)
	goto L180
L179:
	;
	v1033 = v1022
	goto L180
L180:
	;
	v1034 = v1033 + v839
	v1035 = v26 + v1034
	v1036 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v1035))) = uint8(v1036)
	v1041 = base.B2i32(v1023 < int32(0))
	if v1023 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1042 = int32(45)
	goto L183
L182:
	;
	v1042 = int32(43)
	goto L183
L183:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1035)+1)) = uint8(v1042)
	if v1023 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1046 = int32(1) - v832
	goto L186
L185:
	;
	v1046 = v1023
	goto L186
L186:
	;
	v1051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1035)+2)) = uint16(v1051)
	v1240 = v1034 + int32(4)
	goto L51
L187:
	;
	if base.Ui32(v1119) < base.Ui32(int32(100)) {
		goto L195
	} else {
		goto L196
	}
L188:
	;
	v1118 = v1060
	v1119 = v814
	goto L187
L189:
	;
	goto L190
L190:
	;
	v1067 = v814
	v1068 = v1060
	goto L191
L191:
	;
	v1084 = v844 + v1058 + v831 - v1068
	v1085 = int32(4)
	v1088 = base.I32_div_u_s(v1067, int32(10000))
	v1091 = v1088*int32(-10000) + v1067
	v1092 = int32(100)
	v1093 = base.I32_div_u_s(v1091, v1092)
	v1094 = int32(1)
	v1098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1093<<(uint(v1094)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1084-v1085))) = uint16(v1098)
	v1109 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1091-v1093*v1092)<<(uint(v1094)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1084-int32(2)))) = uint16(v1109)
	v1112 = v1068 + v1085
	if base.Ui32(int32(99999999)) < base.Ui32(v1067) {
		v1067 = v1088
		v1068 = v1112
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v1118 = v1112
	v1119 = v1088
	goto L187
L193:
	;
	goto L192
L194:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1158) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v1158 = v1119
	v1159 = v1118
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1139 = int32(2)
	v1141 = int32(65535)
	v1143 = int32(100)
	v1144 = base.I32_div_u_s(v1119&v1141, v1143)
	v1154 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1119-v1144*v1143)&v1141<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v844+v1058+v831-v1118-v1139))) = uint16(v1154)
	v1158 = v1144
	v1159 = v1118 | v1139
	goto L194
L198:
	;
	v1177 = int32(1)
	if v1058 == v1177 {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v844+v1058+v831-v1159-int32(2)))) = uint16(v1171)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v1175 = v1158 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v844+v1058))) = uint8(v1175)
	goto L198
L202:
	;
	v1240 = v1217 + int32(base.Ui32(v138)>>(uint(int32(31))%32))
	goto L51
L203:
	;
	if v832&int32(4) != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	if v824 < int32(0) {
		goto L215
	} else {
		goto L216
	}
L206:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v844)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v1182
	v1185 = int32(5)
	goto L208
L207:
	;
	v1185 = v1177
	goto L208
L208:
	;
	if v832&int32(2) != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1188 = v844 + v1185
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1188))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1188-int32(1)))) = uint16(v1191)
	v1196 = v1185 | int32(2)
	goto L211
L210:
	;
	v1196 = v1185
	goto L211
L211:
	;
	if v832&int32(1) != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1199 = v844 + v1196
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1199-int32(1)))) = uint8(v1202)
	goto L214
L213:
	;
	goto L214
L214:
	;
	v1206 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v844+v832))) = uint8(v1206)
	v1217 = v831 + int32(1)
	goto L202
L215:
	;
	v1214 = int32(2) - v824
	goto L217
L216:
	;
	v1214 = v832
	goto L217
L217:
	;
	v1217 = v1214
	goto L202
L218:
	;
	v1247 = v1241
	v1250 = v24
	goto L219
L219:
	;
	v1253 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1247))) = uint8(v1253)
	v1255 = int32(1)
	v1256 = v1247 + v1255
	v1260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(8)+v1250<<(uint(v1255)%32)))))
	v1265 = v1260 & int32(1023)
	v1269 = v1260 << (uint(int32(16)) % 32) & int32(-2147483648)
	v1272 = int32(31)
	v1273 = int32(base.Ui32(v1260)>>(uint(int32(10))%32)) & v1272
	if v1273 != v1272 {
		goto L225
	} else {
		goto L226
	}
L220:
	;
	v2477 = v2472
	goto L6
L221:
	;
	v1352 = int32(0)
	v1369 = base.I32_reinterpret_f32(base.F32_reinterpret_i32(v1346 | v1345<<(uint(int32(13))%32)))
	v1371 = v1369 & int32(8388607)
	v1374 = int32(255)
	v1375 = int32(base.Ui32(v1369)>>(uint(int32(23))%32)) & v1374
	if v1371|v1375 != 0 {
		goto L265
	} else {
		goto L266
	}
L222:
	;
	goto L221
L223:
	;
	v1345 = v1265
	v1346 = v1273<<(uint(int32(23))%32) + v1269 + int32(939524096)
	goto L222
L224:
	;
	if v1260&int32(512) != 0 {
		goto L234
	} else {
		goto L235
	}
L225:
	;
	if v1273 != 0 {
		goto L223
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if v1265 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	if v1265 != 0 {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	v1345 = int32(0)
	v1346 = v1269
	goto L222
L230:
	;
	v1345 = int32(0)
	v1346 = v1269 | int32(2139095040)
	goto L222
L231:
	;
	goto L232
L232:
	;
	v1345 = v1265
	v1346 = v1269 | int32(2143289344)
	goto L222
L233:
	;
	v1345 = v1333 & int32(1022)
	v1346 = v1335 | v1269
	goto L222
L234:
	;
	v1333 = v1265 << (uint(int32(1)) % 32)
	v1335 = int32(939524096)
	goto L233
L235:
	;
	goto L236
L236:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1265) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1333 = v1265 << (uint(int32(2)) % 32)
	v1335 = int32(931135488)
	goto L233
L238:
	;
	goto L239
L239:
	;
	if base.Ui32(int32(127)) < base.Ui32(v1265) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1333 = v1265 << (uint(int32(3)) % 32)
	v1335 = int32(922746880)
	goto L233
L241:
	;
	goto L242
L242:
	;
	if base.Ui32(int32(63)) < base.Ui32(v1265) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1333 = v1265 << (uint(int32(4)) % 32)
	v1335 = int32(914358272)
	goto L233
L244:
	;
	goto L245
L245:
	;
	if base.Ui32(int32(31)) < base.Ui32(v1265) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1333 = v1265 << (uint(int32(5)) % 32)
	v1335 = int32(905969664)
	goto L233
L247:
	;
	goto L248
L248:
	;
	if base.Ui32(int32(15)) < base.Ui32(v1265) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1333 = v1265 << (uint(int32(6)) % 32)
	v1335 = int32(897581056)
	goto L233
L250:
	;
	goto L251
L251:
	;
	if base.Ui32(int32(7)) < base.Ui32(v1265) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1333 = v1265 << (uint(int32(7)) % 32)
	v1335 = int32(889192448)
	goto L233
L253:
	;
	goto L254
L254:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1265) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1333 = v1265 << (uint(int32(8)) % 32)
	v1335 = int32(880803840)
	goto L233
L256:
	;
	goto L257
L257:
	;
	v1328 = base.B2i32(v1265 == int32(1))
	if v1265 == int32(1) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1329 = int32(1024)
	goto L260
L259:
	;
	v1329 = v1265 << (uint(int32(9)) % 32)
	goto L260
L260:
	;
	if v1265 == int32(1) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1332 = int32(864026624)
	goto L263
L262:
	;
	v1332 = int32(872415232)
	goto L263
L263:
	;
	v1333 = v1329
	v1335 = v1332
	goto L233
L264:
	;
	v2472 = v2471 + v1256
	v2474 = v1250 + int32(1)
	if v2474 != v14 {
		v1247 = v2472
		v1250 = v2474
		goto L219
	} else {
		goto L431
	}
L265:
	;
	v1380 = base.B2i32(v1375 != v1374)
	goto L267
L266:
	;
	v1380 = v1352
	goto L267
L267:
	;
	if v1380 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	if v1371 != 0 {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	goto L270
L270:
	;
	if base.Ui32(int32(23)) < base.Ui32(v1375-int32(127)) {
		goto L288
	} else {
		goto L289
	}
L271:
	;
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1139])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1256)+2)) = uint8(v1384)
	v1387 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1140])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1256))) = uint16(v1387)
	v2471 = int32(3)
	goto L264
L272:
	;
	goto L273
L273:
	;
	if v1369 < int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1392 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1256))) = uint8(v1392)
	goto L276
L275:
	;
	goto L276
L276:
	;
	v1396 = v1256 + int32(base.Ui32(v1369)>>(uint(int32(31))%32))
	if v1375 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1396))) = int64(8751735898823355977)
	if v1369 < int32(0) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	v1404 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1396))) = uint8(v1404)
	if v1369 < int32(0) {
		goto L283
	} else {
		goto L284
	}
L280:
	;
	v1403 = int32(9)
	goto L282
L281:
	;
	v1403 = int32(8)
	goto L282
L282:
	;
	v2471 = v1403
	goto L264
L283:
	;
	v1410 = int32(2)
	goto L285
L284:
	;
	v1410 = int32(1)
	goto L285
L285:
	;
	v2471 = v1410
	goto L264
L286:
	;
	v2063 = v2062 + v2055
	v2064 = int32(0)
	if v1369 < v2064 {
		goto L356
	} else {
		goto L357
	}
L287:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v2003) {
		v2045 = v2003
		v2055 = v2013
		v2062 = int32(8)
		goto L286
	} else {
		goto L347
	}
L288:
	;
	v1428 = int32(2)
	v1434 = v1371 << (uint(v1428) % 32)
	if v1375 != 0 {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	v1415 = int32(-1)
	v1417 = int32(150) - v1375
	if v1371&(v1415<<(uint(v1417)%32)^v1415) != 0 {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v2003 = int32(base.Ui32(v1371|int32(8388608)) >> (uint(v1417) % 32))
	v2013 = v1352
	goto L287
L291:
	;
	v1437 = v1434 | int32(33554432)
	goto L293
L292:
	;
	v1437 = v1434
	goto L293
L293:
	;
	v1438 = base.B2i32(v1371 != int32(0)) | base.B2i32(base.Ui32(v1375) < base.Ui32(v1428)) ^ int32(-1) + v1437
	v1440 = v1437 | int32(2)
	if v1375 != 0 {
		goto L297
	} else {
		goto L298
	}
L294:
	;
	v1996 = v1978 + v1984
	v1997 = v1995 + v1980
	if base.Ui32(v1997) <= base.Ui32(int32(99999999)) {
		v2003 = v1997
		v2013 = v1996
		goto L287
	} else {
		goto L346
	}
L295:
	;
	v1914 = int32(0)
	v1915 = int32(10)
	v1916 = base.I32_div_u_s(v1904, v1915)
	v1918 = base.I32_div_u_s(v1906, v1915)
	if base.Ui32(v1918) < base.Ui32(v1916) {
		goto L340
	} else {
		goto L341
	}
L296:
	;
	v1818 = int32(0)
	v1819 = int32(10)
	v1820 = base.I32_div_u_s(v1808, v1819)
	v1822 = base.I32_div_u_s(v1810, v1819)
	if base.Ui32(v1820) <= base.Ui32(v1822) {
		goto L334
	} else {
		goto L335
	}
L297:
	;
	v1444 = v1375 - int32(152)
	goto L299
L298:
	;
	v1444 = int32(-151)
	goto L299
L299:
	;
	if int32(0) <= v1444 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1450 = int32(base.Ui32(v1444*int32(78913)) >> (uint(int32(18)) % 32))
	v1455 = *(*int64)(unsafe.Add(mBase, uint32(v1450<<(uint(int32(3))%32))+uint32(_consts[1141])))
	v1457 = v1455 & int64(4294967295)
	v1458 = base.I64_extend_i32_u(v1438)
	v1460 = int64(32)
	v1462 = base.I32_wrap_i64(int64(base.Ui64(v1457*v1458) >> (uint(v1460) % 64)))
	v1464 = int64(base.Ui64(v1455) >> (uint(v1460) % 64))
	v1465 = v1464 * v1458
	v1467 = v1462 + base.I32_wrap_i64(v1465)
	v1474 = v1450 - v1444
	v1479 = v1474 + int32(base.Ui32(v1450*int32(1217359))>>(uint(int32(19))%32))
	v1480 = int32(5) - v1479
	v1483 = v1479 + int32(27)
	v1485 = (base.B2i32(base.Ui32(v1467) < base.Ui32(v1462))+base.I32_wrap_i64(int64(base.Ui64(v1465)>>(uint(v1460)%64))))<<(uint(v1480)%32) | int32(base.Ui32(v1467)>>(uint(v1483)%32))
	v1486 = base.I64_extend_i32_u(v1440)
	v1490 = base.I32_wrap_i64(int64(base.Ui64(v1457*v1486) >> (uint(v1460) % 64)))
	v1491 = v1486 * v1464
	v1493 = v1490 + base.I32_wrap_i64(v1491)
	v1501 = (base.B2i32(base.Ui32(v1493) < base.Ui32(v1490))+base.I32_wrap_i64(int64(base.Ui64(v1491)>>(uint(v1460)%64))))<<(uint(v1480)%32) | int32(base.Ui32(v1493)>>(uint(v1483)%32))
	v1502 = base.I64_extend_i32_u(v1437)
	v1506 = base.I32_wrap_i64(int64(base.Ui64(v1457*v1502) >> (uint(v1460) % 64)))
	v1507 = v1502 * v1464
	v1509 = v1506 + base.I32_wrap_i64(v1507)
	v1517 = (base.B2i32(base.Ui32(v1509) < base.Ui32(v1506))+base.I32_wrap_i64(int64(base.Ui64(v1507)>>(uint(v1460)%64))))<<(uint(v1480)%32) | int32(base.Ui32(v1509)>>(uint(v1483)%32))
	v1518 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1444) {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	goto L302
L302:
	;
	v1663 = v1444 * int32(-732923)
	v1665 = int32(base.Ui32(v1663) >> (uint(int32(20)) % 32))
	v1666 = v1665 + v1444
	v1670 = *(*int64)(unsafe.Add(mBase, uint32(int32(1876512)-v1666<<(uint(int32(3))%32))))
	v1672 = v1670 & int64(4294967295)
	v1673 = base.I64_extend_i32_u(v1438)
	v1675 = int64(32)
	v1677 = base.I32_wrap_i64(int64(base.Ui64(v1672*v1673) >> (uint(v1675) % 64)))
	v1679 = int64(base.Ui64(v1670) >> (uint(v1675) % 64))
	v1680 = v1679 * v1673
	v1682 = v1677 + base.I32_wrap_i64(v1680)
	v1693 = v1665 - int32(base.Ui32(v1666*int32(-1217359))>>(uint(int32(19))%32))
	v1694 = int32(4) - v1693
	v1697 = v1693 + int32(28)
	v1699 = (base.B2i32(base.Ui32(v1682) < base.Ui32(v1677))+base.I32_wrap_i64(int64(base.Ui64(v1680)>>(uint(v1675)%64))))<<(uint(v1694)%32) | int32(base.Ui32(v1682)>>(uint(v1697)%32))
	v1700 = base.I64_extend_i32_u(v1437)
	v1704 = base.I32_wrap_i64(int64(base.Ui64(v1672*v1700) >> (uint(v1675) % 64)))
	v1705 = v1700 * v1679
	v1707 = v1704 + base.I32_wrap_i64(v1705)
	v1715 = (base.B2i32(base.Ui32(v1707) < base.Ui32(v1704))+base.I32_wrap_i64(int64(base.Ui64(v1705)>>(uint(v1675)%64))))<<(uint(v1694)%32) | int32(base.Ui32(v1707)>>(uint(v1697)%32))
	v1716 = base.I64_extend_i32_u(v1440)
	v1720 = base.I32_wrap_i64(int64(base.Ui64(v1672*v1716) >> (uint(v1675) % 64)))
	v1721 = v1679 * v1716
	v1723 = v1720 + base.I32_wrap_i64(v1721)
	v1731 = (base.B2i32(base.Ui32(v1723) < base.Ui32(v1720))+base.I32_wrap_i64(int64(base.Ui64(v1721)>>(uint(v1675)%64))))<<(uint(v1694)%32) | int32(base.Ui32(v1723)>>(uint(v1697)%32))
	v1733 = v1731 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v1663) {
		goto L324
	} else {
		goto L325
	}
L303:
	;
	v1524 = int32(10)
	v1525 = base.I32_div_u_s(v1501-int32(1), v1524)
	v1527 = base.I32_div_u_s(v1485, v1524)
	if base.Ui32(v1525) <= base.Ui32(v1527) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	v1572 = v1518
	goto L305
L305:
	;
	v1577 = base.I32_rem_u_s(v1437, int32(5))
	if v1577 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L306:
	;
	v1530 = v1450 - int32(1)
	v1535 = *(*int64)(unsafe.Add(mBase, uint32(v1530<<(uint(int32(3))%32))+uint32(_consts[1141])))
	v1539 = int64(32)
	v1541 = base.I32_wrap_i64(int64(base.Ui64(v1535&int64(4294967295)*v1502) >> (uint(v1539) % 64)))
	v1544 = int64(base.Ui64(v1535)>>(uint(v1539)%64)) * v1502
	v1546 = v1541 + base.I32_wrap_i64(v1544)
	v1557 = v1474 + int32(base.Ui32(v1530*int32(1217359))>>(uint(int32(19))%32))
	v1565 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1546) < base.Ui32(v1541))+base.I32_wrap_i64(int64(base.Ui64(v1544)>>(uint(v1539)%64))))<<(uint(int32(6)-v1557)%32)|int32(base.Ui32(v1546)>>(uint(v1557+int32(26))%32)), int32(10))
	v1566 = v1565
	goto L308
L307:
	;
	v1566 = v1518
	goto L308
L308:
	;
	if base.Ui32(int32(33)) < base.Ui32(v1444) {
		v1898 = v1517
		v1900 = v1566
		v1903 = v1450
		v1904 = v1501
		v1906 = v1485
		goto L295
	} else {
		goto L309
	}
L309:
	;
	v1572 = v1566
	goto L305
L310:
	;
	v1582 = v1437
	v1584 = v1518
	goto L313
L311:
	;
	goto L312
L312:
	;
	v1608 = int32(0)
	v1610 = base.I32_rem_u_s(v1440, int32(5))
	if v1610 == v1608 {
		goto L317
	} else {
		goto L318
	}
L313:
	;
	v1600 = v1584 + int32(1)
	v1601 = int32(5)
	v1602 = base.I32_div_u_s(v1582, v1601)
	v1604 = base.I32_rem_u_s(v1602, v1601)
	if v1604 == int32(0) {
		v1582 = v1602
		v1584 = v1600
		goto L313
	} else {
		goto L315
	}
L314:
	;
	if base.Ui32(v1600) < base.Ui32(v1450) {
		v1898 = v1517
		v1900 = v1572
		v1903 = v1450
		v1904 = v1501
		v1906 = v1485
		goto L295
	} else {
		goto L316
	}
L315:
	;
	goto L314
L316:
	;
	v1802 = v1517
	v1804 = v1572
	v1807 = v1450
	v1808 = v1501
	v1810 = v1485
	goto L296
L317:
	;
	v1615 = v1608
	v1619 = v1440
	goto L320
L318:
	;
	v1642 = v1608
	goto L319
L319:
	;
	v1898 = v1517
	v1900 = v1572
	v1903 = v1450
	v1904 = v1501 - base.B2i32(base.Ui32(v1450) <= base.Ui32(v1642))
	v1906 = v1485
	goto L295
L320:
	;
	v1633 = v1615 + int32(1)
	v1634 = int32(5)
	v1635 = base.I32_div_u_s(v1619, v1634)
	v1637 = base.I32_rem_u_s(v1635, v1634)
	if v1637 == int32(0) {
		v1615 = v1633
		v1619 = v1635
		goto L320
	} else {
		goto L322
	}
L321:
	;
	v1642 = v1633
	goto L319
L322:
	;
	goto L321
L323:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v1663) {
		v1898 = v1715
		v1900 = v1780
		v1903 = v1666
		v1904 = v1731
		v1906 = v1699
		goto L295
	} else {
		goto L331
	}
L324:
	;
	v1736 = int32(10)
	v1737 = base.I32_div_u_s(v1733, v1736)
	v1739 = base.I32_div_u_s(v1699, v1736)
	if base.Ui32(v1737) <= base.Ui32(v1739) {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	v1786 = v1352
	goto L326
L326:
	;
	v1802 = v1715
	v1804 = v1786
	v1807 = v1666
	v1808 = v1733
	v1810 = v1699
	goto L296
L327:
	;
	v1742 = int32(1) - v1666
	v1747 = *(*int64)(unsafe.Add(mBase, uint32(v1742<<(uint(int32(3))%32))+uint32(_consts[1142])))
	v1751 = int64(32)
	v1753 = base.I32_wrap_i64(int64(base.Ui64(v1747&int64(4294967295)*v1700) >> (uint(v1751) % 64)))
	v1756 = int64(base.Ui64(v1747)>>(uint(v1751)%64)) * v1700
	v1758 = v1753 + base.I32_wrap_i64(v1756)
	v1771 = v1665 + (int32(base.Ui32(v1742*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v1779 = base.I32_rem_u_s((base.B2i32(base.Ui32(v1758) < base.Ui32(v1753))+base.I32_wrap_i64(int64(base.Ui64(v1756)>>(uint(v1751)%64))))<<(uint(int32(4)-v1771)%32)|int32(base.Ui32(v1758)>>(uint(v1771+int32(28))%32)), int32(10))
	v1780 = v1779
	goto L329
L328:
	;
	v1780 = v1352
	goto L329
L329:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v1663) {
		goto L323
	} else {
		goto L330
	}
L330:
	;
	v1786 = v1780
	goto L326
L331:
	;
	v1792 = int32(-1)
	if v1437&(v1792<<(uint(v1665-int32(1))%32)^v1792) != 0 {
		v1898 = v1715
		v1900 = v1780
		v1903 = v1666
		v1904 = v1731
		v1906 = v1699
		goto L295
	} else {
		goto L332
	}
L332:
	;
	v1802 = v1715
	v1804 = v1780
	v1807 = v1666
	v1808 = v1731
	v1810 = v1699
	goto L296
L333:
	;
	v1885 = v1870 & int32(255)
	v1978 = v1866
	v1980 = v1868
	v1984 = v1807
	v1995 = (v1883|base.B2i32(v1885 != int32(5))|v1868)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1885)) | base.B2i32(v1868 == v1875)
	goto L294
L334:
	;
	v1866 = v1818
	v1868 = v1802
	v1870 = v1804
	v1875 = v1810
	v1883 = int32(0)
	goto L333
L335:
	;
	goto L336
L336:
	;
	v1828 = v1818
	v1829 = v1802
	v1831 = v1804
	v1833 = v1820
	v1835 = v1822
	v1836 = int32(1)
	goto L337
L337:
	;
	v1846 = v1828 + int32(1)
	v1847 = int32(10)
	v1848 = base.I32_div_u_s(v1829, v1847)
	v1851 = v1829 - v1848*v1847
	v1856 = v1836 & base.B2i32(v1831&int32(255) == int32(0))
	v1858 = base.I32_div_u_s(v1833, v1847)
	v1860 = base.I32_div_u_s(v1835, v1847)
	if base.Ui32(v1860) < base.Ui32(v1858) {
		v1828 = v1846
		v1829 = v1848
		v1831 = v1851
		v1833 = v1858
		v1835 = v1860
		v1836 = v1856
		goto L337
	} else {
		goto L339
	}
L338:
	;
	v1866 = v1846
	v1868 = v1848
	v1870 = v1851
	v1875 = v1835
	v1883 = v1856 ^ int32(1)
	goto L333
L339:
	;
	goto L338
L340:
	;
	v1922 = v1914
	v1923 = v1898
	v1924 = v1916
	v1926 = v1918
	goto L343
L341:
	;
	v1953 = v1914
	v1954 = v1898
	v1956 = v1900
	v1962 = v1906
	goto L342
L342:
	;
	v1978 = v1953
	v1980 = v1954
	v1984 = v1903
	v1995 = base.B2i32(v1954 == v1962) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1956&int32(255)))
	goto L294
L343:
	;
	v1940 = v1922 + int32(1)
	v1941 = int32(10)
	v1942 = base.I32_div_u_s(v1923, v1941)
	v1944 = base.I32_div_u_s(v1924, v1941)
	v1946 = base.I32_div_u_s(v1926, v1941)
	if base.Ui32(v1946) < base.Ui32(v1944) {
		v1922 = v1940
		v1923 = v1942
		v1924 = v1944
		v1926 = v1946
		goto L343
	} else {
		goto L345
	}
L344:
	;
	v1953 = v1940
	v1954 = v1942
	v1956 = v1923 - v1942*int32(10)
	v1962 = v1926
	goto L342
L345:
	;
	goto L344
L346:
	;
	v2045 = v1997
	v2055 = v1996
	v2062 = int32(9)
	goto L286
L347:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v2003) {
		v2045 = v2003
		v2055 = v2013
		v2062 = int32(7)
		goto L286
	} else {
		goto L348
	}
L348:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v2003) {
		v2045 = v2003
		v2055 = v2013
		v2062 = int32(6)
		goto L286
	} else {
		goto L349
	}
L349:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v2003) {
		v2045 = v2003
		v2055 = v2013
		v2062 = int32(5)
		goto L286
	} else {
		goto L350
	}
L350:
	;
	if base.Ui32(int32(999)) < base.Ui32(v2003) {
		v2045 = v2003
		v2055 = v2013
		v2062 = int32(4)
		goto L286
	} else {
		goto L351
	}
L351:
	;
	if base.Ui32(int32(99)) < base.Ui32(v2003) {
		v2045 = v2003
		v2055 = v2013
		v2062 = int32(3)
		goto L286
	} else {
		goto L352
	}
L352:
	;
	if base.Ui32(int32(9)) < base.Ui32(v2003) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v2042 = int32(2)
	goto L355
L354:
	;
	v2042 = int32(1)
	goto L355
L355:
	;
	v2045 = v2003
	v2055 = v2013
	v2062 = v2042
	goto L286
L356:
	;
	v2067 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1256))) = uint8(v2067)
	v2070 = int32(1)
	goto L358
L357:
	;
	v2070 = v2064
	goto L358
L358:
	;
	if base.Ui32(v2063+int32(3)) <= base.Ui32(int32(9)) {
		goto L361
	} else {
		goto L362
	}
L359:
	;
	v2291 = int32(0)
	if base.Ui32(v2045) < base.Ui32(int32(10000)) {
		goto L401
	} else {
		goto L402
	}
L360:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2075))) = v2287
	v2289 = v2286
	goto L359
L361:
	;
	v2075 = v1256 + v2070
	v2076 = int32(0)
	if v2063 <= v2076 {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	goto L363
L363:
	;
	if v2055 != 0 {
		goto L369
	} else {
		goto L370
	}
L364:
	;
	v2286 = int32(2) - v2063
	v2287 = int64(3472328296227679792)
	goto L360
L365:
	;
	goto L366
L366:
	;
	if int32(0) <= v2055 {
		v2286 = v2076
		v2287 = int64(3472328296227680304)
		goto L360
	} else {
		goto L367
	}
L367:
	;
	v2289 = int32(1)
	goto L359
L368:
	;
	v2135 = int32(0)
	if base.Ui32(v2119) < base.Ui32(int32(10000)) {
		goto L377
	} else {
		goto L378
	}
L369:
	;
	v2119 = v2045
	v2125 = v2062
	goto L368
L370:
	;
	goto L371
L371:
	;
	v2088 = v2045
	v2093 = v2062
	goto L372
L372:
	;
	if v2088&int32(1) != 0 {
		v2119 = v2088
		v2125 = v2093
		goto L368
	} else {
		goto L374
	}
L373:
	;
	v2119 = v2088
	v2125 = v2093
	goto L368
L374:
	;
	v2112 = base.I32_div_u_s(v2088, int32(10))
	if int32(0)-v2088 == v2112*int32(-10) {
		v2088 = v2112
		v2093 = v2093 - int32(1)
		goto L372
	} else {
		goto L375
	}
L375:
	;
	goto L373
L376:
	;
	if base.Ui32(v2194) < base.Ui32(int32(100)) {
		goto L384
	} else {
		goto L385
	}
L377:
	;
	v2192 = v2135
	v2194 = v2119
	goto L376
L378:
	;
	goto L379
L379:
	;
	v2142 = v2135
	v2143 = v2119
	goto L380
L380:
	;
	v2159 = v1256 + v2070 + v2125 - v2142
	v2163 = base.I32_div_u_s(v2143, int32(10000))
	v2166 = v2163*int32(-10000) + v2143
	v2167 = int32(100)
	v2168 = base.I32_div_u_s(v2166, v2167)
	v2169 = int32(1)
	v2173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2168<<(uint(v2169)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2159-int32(3)))) = uint16(v2173)
	v2184 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2166-v2168*v2167)<<(uint(v2169)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2159-v2169))) = uint16(v2184)
	v2187 = v2142 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v2143) {
		v2142 = v2187
		v2143 = v2163
		goto L380
	} else {
		goto L382
	}
L381:
	;
	v2192 = v2187
	v2194 = v2163
	goto L376
L382:
	;
	goto L381
L383:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2234) {
		goto L388
	} else {
		goto L389
	}
L384:
	;
	v2233 = v2192
	v2234 = v2194
	goto L383
L385:
	;
	goto L386
L386:
	;
	v2216 = int32(65535)
	v2218 = int32(100)
	v2219 = base.I32_div_u_s(v2194&v2216, v2218)
	v2229 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2194-v2219*v2218)&v2216<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1256+v2070+v2125+(v2192^int32(-1))))) = uint16(v2229)
	v2233 = v2192 | int32(2)
	v2234 = v2219
	goto L383
L387:
	;
	v2253 = int32(1)
	v2254 = v2063 - v2253
	v2255 = v1256 + v2070
	*(*uint8)(unsafe.Add(mBase, uint32(v2255))) = uint8(v2252)
	if base.Ui32(int32(2)) <= base.Ui32(v2125) {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	v2241 = v2234 << (uint(int32(1)) % 32)
	v2244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241)+uint32(_consts[1144]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1256+(v2070+v2125-v2233)))) = uint8(v2244)
	v2248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241)+uint32(_consts[1143]))))
	v2252 = v2248
	goto L387
L389:
	;
	goto L390
L390:
	;
	v2252 = v2234 | int32(48)
	goto L387
L391:
	;
	v2260 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2255)+1)) = uint8(v2260)
	v2264 = v2125 + int32(1)
	goto L393
L392:
	;
	v2264 = v2253
	goto L393
L393:
	;
	v2265 = v2264 + v2070
	v2266 = v1256 + v2265
	v2267 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2266))) = uint8(v2267)
	v2272 = base.B2i32(v2254 < int32(0))
	if v2254 < int32(0) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2273 = int32(45)
	goto L396
L395:
	;
	v2273 = int32(43)
	goto L396
L396:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2266)+1)) = uint8(v2273)
	if v2254 < int32(0) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2277 = int32(1) - v2063
	goto L399
L398:
	;
	v2277 = v2254
	goto L399
L399:
	;
	v2282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2277<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2266)+2)) = uint16(v2282)
	v2471 = v2265 + int32(4)
	goto L264
L400:
	;
	if base.Ui32(v2350) < base.Ui32(int32(100)) {
		goto L408
	} else {
		goto L409
	}
L401:
	;
	v2349 = v2291
	v2350 = v2045
	goto L400
L402:
	;
	goto L403
L403:
	;
	v2298 = v2045
	v2299 = v2291
	goto L404
L404:
	;
	v2315 = v2075 + v2289 + v2062 - v2299
	v2316 = int32(4)
	v2319 = base.I32_div_u_s(v2298, int32(10000))
	v2322 = v2319*int32(-10000) + v2298
	v2323 = int32(100)
	v2324 = base.I32_div_u_s(v2322, v2323)
	v2325 = int32(1)
	v2329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2324<<(uint(v2325)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2315-v2316))) = uint16(v2329)
	v2340 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2322-v2324*v2323)<<(uint(v2325)%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2315-int32(2)))) = uint16(v2340)
	v2343 = v2299 + v2316
	if base.Ui32(int32(99999999)) < base.Ui32(v2298) {
		v2298 = v2319
		v2299 = v2343
		goto L404
	} else {
		goto L406
	}
L405:
	;
	v2349 = v2343
	v2350 = v2319
	goto L400
L406:
	;
	goto L405
L407:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2389) {
		goto L412
	} else {
		goto L413
	}
L408:
	;
	v2389 = v2350
	v2390 = v2349
	goto L407
L409:
	;
	goto L410
L410:
	;
	v2370 = int32(2)
	v2372 = int32(65535)
	v2374 = int32(100)
	v2375 = base.I32_div_u_s(v2350&v2372, v2374)
	v2385 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2350-v2375*v2374)&v2372<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2075+v2289+v2062-v2349-v2370))) = uint16(v2385)
	v2389 = v2375
	v2390 = v2349 | v2370
	goto L407
L411:
	;
	v2408 = int32(1)
	if v2289 == v2408 {
		goto L416
	} else {
		goto L417
	}
L412:
	;
	v2402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2389<<(uint(int32(1))%32))+uint32(_consts[1143]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2075+v2289+v2062-v2390-int32(2)))) = uint16(v2402)
	goto L411
L413:
	;
	goto L414
L414:
	;
	v2406 = v2389 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2075+v2289))) = uint8(v2406)
	goto L411
L415:
	;
	v2471 = v2448 + int32(base.Ui32(v1369)>>(uint(int32(31))%32))
	goto L264
L416:
	;
	if v2063&int32(4) != 0 {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	goto L418
L418:
	;
	if v2055 < int32(0) {
		goto L428
	} else {
		goto L429
	}
L419:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2075)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v2075))) = v2413
	v2416 = int32(5)
	goto L421
L420:
	;
	v2416 = v2408
	goto L421
L421:
	;
	if v2063&int32(2) != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v2419 = v2075 + v2416
	v2422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2419))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2419-int32(1)))) = uint16(v2422)
	v2427 = v2416 | int32(2)
	goto L424
L423:
	;
	v2427 = v2416
	goto L424
L424:
	;
	if v2063&int32(1) != 0 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v2430 = v2075 + v2427
	v2433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2430))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2430-int32(1)))) = uint8(v2433)
	goto L427
L426:
	;
	goto L427
L427:
	;
	v2437 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2075+v2063))) = uint8(v2437)
	v2448 = v2062 + int32(1)
	goto L415
L428:
	;
	v2445 = int32(2) - v2055
	goto L430
L429:
	;
	v2445 = v2063
	goto L430
L430:
	;
	v2448 = v2445
	goto L415
L431:
	;
	goto L220
L432:
	;
	F_pfree(m, v10)
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L1
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	return v20
L435:
	;
	goto L434
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
