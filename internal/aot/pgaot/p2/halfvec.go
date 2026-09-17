package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HalfvecL1DistanceDefault(m *base.Module, l0 int32, l1 int32, l2 int32) float32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 float32
	_ = v6
	var v12 int32
	_ = v12
	var v14 float32
	_ = v14
	var v16 int32
	_ = v16
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 float32
	_ = v203
	var v205 int32
	_ = v205
	var v212 float32
	_ = v212
	v4 = int32(0)
	v6 = float32(0)
	if v4 < l0 {
		v12 = v4
		v14 = v6
		for {
			v16 = v12 << (uint(int32(1)) % 32)
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v16))))
			v23 = v18 & int32(1023)
			v27 = v18 << (uint(int32(16)) % 32) & int32(-2147483648)
			v30 = int32(31)
			v31 = int32(base.Ui32(v18)>>(uint(int32(10))%32)) & v30
			if v31 != v30 {
				if v31 != 0 {
					v103 = v23
					v104 = v31<<(uint(int32(23))%32) + v27 + int32(939524096)
				} else {
					if v23 != 0 {
						if v18&int32(512) != 0 {
							v91 = v23 << (uint(int32(1)) % 32)
							v93 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v23) {
								v91 = v23 << (uint(int32(2)) % 32)
								v93 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v23) {
									v91 = v23 << (uint(int32(3)) % 32)
									v93 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v23) {
										v91 = v23 << (uint(int32(4)) % 32)
										v93 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v23) {
											v91 = v23 << (uint(int32(5)) % 32)
											v93 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v23) {
												v91 = v23 << (uint(int32(6)) % 32)
												v93 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v23) {
													v91 = v23 << (uint(int32(7)) % 32)
													v93 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v23) {
														v91 = v23 << (uint(int32(8)) % 32)
														v93 = int32(880803840)
													} else {
														v86 = base.B2i32(v23 == int32(1))
														if v23 == int32(1) {
															v87 = int32(1024)
														} else {
															v87 = v23 << (uint(int32(9)) % 32)
														}
														if v23 == int32(1) {
															v90 = int32(864026624)
														} else {
															v90 = int32(872415232)
														}
														v91 = v87
														v93 = v90
													}
												}
											}
										}
									}
								}
							}
						}
						v103 = v91 & int32(1022)
						v104 = v93 | v27
					} else {
						v103 = int32(0)
						v104 = v27
					}
				}
			} else {
				if v23 == int32(0) {
					v103 = int32(0)
					v104 = v27 | int32(2139095040)
				} else {
					v103 = v23
					v104 = v27 | int32(2143289344)
				}
			}
			v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v16))))
			v115 = v110 & int32(1023)
			v119 = v110 << (uint(int32(16)) % 32) & int32(-2147483648)
			v122 = int32(31)
			v123 = int32(base.Ui32(v110)>>(uint(int32(10))%32)) & v122
			if v123 != v122 {
				if v123 != 0 {
					v195 = v115
					v196 = v123<<(uint(int32(23))%32) + v119 + int32(939524096)
				} else {
					if v115 != 0 {
						if v110&int32(512) != 0 {
							v183 = v115 << (uint(int32(1)) % 32)
							v185 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v115) {
								v183 = v115 << (uint(int32(2)) % 32)
								v185 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v115) {
									v183 = v115 << (uint(int32(3)) % 32)
									v185 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v115) {
										v183 = v115 << (uint(int32(4)) % 32)
										v185 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v115) {
											v183 = v115 << (uint(int32(5)) % 32)
											v185 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v115) {
												v183 = v115 << (uint(int32(6)) % 32)
												v185 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v115) {
													v183 = v115 << (uint(int32(7)) % 32)
													v185 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v115) {
														v183 = v115 << (uint(int32(8)) % 32)
														v185 = int32(880803840)
													} else {
														v178 = base.B2i32(v115 == int32(1))
														if v115 == int32(1) {
															v179 = int32(1024)
														} else {
															v179 = v115 << (uint(int32(9)) % 32)
														}
														if v115 == int32(1) {
															v182 = int32(864026624)
														} else {
															v182 = int32(872415232)
														}
														v183 = v179
														v185 = v182
													}
												}
											}
										}
									}
								}
							}
						}
						v195 = v183 & int32(1022)
						v196 = v185 | v119
					} else {
						v195 = int32(0)
						v196 = v119
					}
				}
			} else {
				if v115 == int32(0) {
					v195 = int32(0)
					v196 = v119 | int32(2139095040)
				} else {
					v195 = v115
					v196 = v119 | int32(2143289344)
				}
			}
			v203 = base.F32_add(base.F32_abs(base.F32_sub(base.F32_reinterpret_i32(v104|v103<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v196|v195<<(uint(int32(13))%32)))), v14)
			v205 = v12 + int32(1)
			if v205 != l0 {
				v12 = v205
				v14 = v203
				continue
			} else {
				break
			}
			break
		}
		v212 = v203
	} else {
		v212 = v6
	}
	return v212
}
func F_halfvec_concat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			v18 = v16 + v17
			F_CheckDim_1(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v23 = F_mul_size(m, int32(2), v18)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = F_add_size(m, int32(8), v23)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = F_palloc0(m, v25)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v27)+4)) = uint16(v18)
							*(*int32)(unsafe.Add(mBase, uint32(v27))) = v25 << (uint(int32(2)) % 32)
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
							if int32(0) < v33 {
								v36 = int32(8)
								v41 = int32(0)
								for {
									v48 = int32(1)
									v49 = v41 << (uint(v48) % 32)
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+v36+v49))))
									*(*uint16)(unsafe.Add(mBase, uint32(v27+v36+v49))) = uint16(v52)
									v55 = v41 + v48
									v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
									if v55 < v56 {
										v41 = v55
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
							if int32(0) < v65 {
								v68 = int32(8)
								v73 = int32(0)
								for {
									v80 = int32(1)
									v81 = v73 << (uint(v80) % 32)
									v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v68+v81))))
									*(*uint16)(unsafe.Add(mBase, uint32(v27+v68+v81+v83<<(uint(v80)%32)))) = uint16(v88)
									v91 = v73 + v80
									v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
									if v91 < v92 {
										v73 = v91
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							return v27
						}
					}
				}
			}
		}
	}
}
func F_halfvec_inner_product(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13911(m, l0, int32(_a_F_halfvec_inner_product_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_halfvec_l2_distance(m *base.Module, l0 int32) int32 {
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
						F_errmsg(m, int32(_a_F_halfvec_l2_distance_0), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_l2_distance_1), int32(80), int32(_a_F_halfvec_l2_distance_2))
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
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_halfvec_l2_distance[0]))
				v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v17), v10+v40, v15+v40)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v50 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_promote_f32(v46)))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v50
					}
				}
			}
		}
	}
}
func F_halfvec_negative_inner_product(m *base.Module, l0 int32) int32 {
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
						F_errmsg(m, int32(_a_F_halfvec_negative_inner_product_0), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_negative_inner_product_1), int32(80), int32(_a_F_halfvec_negative_inner_product_2))
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
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_halfvec_negative_inner_product[0]))
				v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v17), v10+v40, v15+v40)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v50 = F_Float8GetDatum(m, base.F64_promote_f32(base.F32_neg(v46)))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v50
					}
				}
			}
		}
	}
}
func F_halfvec_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pq_getmsgint(m, v13, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = F_pq_getmsgint(m, v13, int32(2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = base.I32_extend16_s(v15)
	F_CheckDim_1(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.B2i32(v12 != int32(-1))&base.B2i32(v12 != v22) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v32 = v20 << (uint(int32(16)) % 32)
	if v32 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L9:
	;
	v35 = F_mul_size(m, int32(2), v22)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v37 = F_add_size(m, int32(8), v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = F_palloc0(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)) = uint16(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v37 << (uint(int32(2)) % 32)
	if int32(0) < v22 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	m.G0 = v10 + int32(32)
	return v39
L16:
	;
	v61 = F_pq_getmsgint(m, v13, int32(2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39+int32(8)+v50<<(uint(int32(1))%32)))) = uint16(v61)
	F_CheckElement_1(m, v61&int32(_a_F_halfvec_recv_0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v69 = v50 + int32(1)
	if v69 != v22 {
		v50 = v69
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	F_errmsg(m, int32(_a_F_halfvec_recv_1), v10+int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_halfvec_recv_2), int32(92), int32(_a_F_halfvec_recv_3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32 >> (uint(int32(16)) % 32)
	F_errmsg(m, int32(_a_F_halfvec_recv_4), v10)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_halfvec_recv_2), int32(390), int32(_a_F_halfvec_recv_5))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
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
func F_halfvec_to_sparsevec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
	F_CheckDim_2(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.B2i32(v23 != int32(-1))&base.B2i32(v23 != v24) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = base.B2i32(v24 <= int32(0))
	if v24 <= int32(0) {
		v135 = v2
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L81
	}
L7:
	;
	F_CheckNnz(m, v135, v24)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L19
	}
L8:
	;
	v36 = v19 + int32(8)
	v37 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v24) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = v37
	v43 = v2
	v51 = v2
	goto L12
L10:
	;
	v91 = v37
	v92 = v2
	goto L11
L11:
	;
	v107 = v91
	v108 = v92
	v110 = int32(0)
	goto L16
L12:
	;
	v57 = v36 + v42<<(uint(int32(1))%32)
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	v59 = int32(_a_F_halfvec_to_sparsevec_0)
	v61 = int32(0)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+2)))
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)))
	v81 = v43 + base.B2i32(v58&v59 != v61) + base.B2i32(v64&v59 != v61) + base.B2i32(v70&v59 != v61) + base.B2i32(v76&v59 != v61)
	v82 = int32(4)
	v83 = v42 + v82
	v85 = v51 + v82
	if v85 != v24&int32(_a_F_halfvec_to_sparsevec_1) {
		v42 = v83
		v43 = v81
		v51 = v85
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v24&int32(3) == int32(0) {
		v135 = v81
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v91 = v83
	v92 = v81
	goto L11
L16:
	;
	v120 = int32(1)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+v107<<(uint(v120)%32)))))
	v128 = v108 + base.B2i32(v123&int32(_a_F_halfvec_to_sparsevec_0) != int32(0))
	v132 = v110 + v120
	if v132 != v24&int32(3) {
		v107 = v107 + v120
		v108 = v128
		v110 = v132
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v135 = v128
	goto L7
L18:
	;
	goto L17
L19:
	;
	v151 = F_mul_size(m, int32(4), v135)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v153 = F_add_size(m, int32(16), v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v156 = F_mul_size(m, int32(4), v135)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v158 = F_add_size(m, v153, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v160 = F_palloc0(m, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v158 << (uint(int32(2)) % 32)
	if v34 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L78
	}
L26:
	;
	v170 = v160 + int32(16)
	v176 = int32(0)
	v178 = v176
	v181 = v176
	goto L29
L27:
	;
	goto L28
L28:
	;
	m.G0 = v16 + int32(16)
	return v160
L29:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+int32(8)+v178<<(uint(int32(1))%32)))))
	if v194&int32(_a_F_halfvec_to_sparsevec_0) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v197 <= v181 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	v294 = v181
	goto L33
L33:
	;
	v299 = v178 + int32(1)
	if v299 != v24 {
		v178 = v299
		v181 = v294
		goto L29
	} else {
		goto L77
	}
L34:
	;
	v200 = v181 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v170+v200))) = v178
	v205 = base.I32_extend16_s(v194) & int32(-2147483648)
	v207 = v194 & int32(1023)
	v210 = int32(31)
	v211 = int32(base.Ui32(v194)>>(uint(int32(10))%32)) & v210
	if v211 != v210 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170+v135<<(uint(int32(2))%32)+v200))) = v283 | v282<<(uint(int32(13))%32)
	v294 = v181 + int32(1)
	goto L33
L36:
	;
	v282 = v207
	v283 = v211<<(uint(int32(23))%32) + v205 + int32(939524096)
	goto L35
L37:
	;
	if v194&int32(512) != 0 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	if v211 != 0 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v207 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v207 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v282 = int32(0)
	v283 = v205
	goto L35
L43:
	;
	v282 = int32(0)
	v283 = v205 | int32(2139095040)
	goto L35
L44:
	;
	goto L45
L45:
	;
	v282 = v207
	v283 = v205 | int32(2143289344)
	goto L35
L46:
	;
	v282 = v271 & int32(1022)
	v283 = v273 | v205
	goto L35
L47:
	;
	v271 = v207 << (uint(int32(1)) % 32)
	v273 = int32(939524096)
	goto L46
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(int32(255)) < base.Ui32(v207) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v271 = v207 << (uint(int32(2)) % 32)
	v273 = int32(931135488)
	goto L46
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(int32(127)) < base.Ui32(v207) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v271 = v207 << (uint(int32(3)) % 32)
	v273 = int32(922746880)
	goto L46
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(63)) < base.Ui32(v207) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v271 = v207 << (uint(int32(4)) % 32)
	v273 = int32(914358272)
	goto L46
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(int32(31)) < base.Ui32(v207) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v271 = v207 << (uint(int32(5)) % 32)
	v273 = int32(905969664)
	goto L46
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(int32(15)) < base.Ui32(v207) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v271 = v207 << (uint(int32(6)) % 32)
	v273 = int32(897581056)
	goto L46
L63:
	;
	goto L64
L64:
	;
	if base.Ui32(int32(7)) < base.Ui32(v207) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v271 = v207 << (uint(int32(7)) % 32)
	v273 = int32(889192448)
	goto L46
L66:
	;
	goto L67
L67:
	;
	if base.Ui32(int32(3)) < base.Ui32(v207) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v271 = v207 << (uint(int32(8)) % 32)
	v273 = int32(880803840)
	goto L46
L69:
	;
	goto L70
L70:
	;
	v266 = base.B2i32(v207 == int32(1))
	if v207 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v267 = int32(1024)
	goto L73
L72:
	;
	v267 = v207 << (uint(int32(9)) % 32)
	goto L73
L73:
	;
	if v207 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v270 = int32(864026624)
	goto L76
L75:
	;
	v270 = int32(872415232)
	goto L76
L76:
	;
	v271 = v267
	v273 = v270
	goto L46
L77:
	;
	goto L30
L78:
	;
	F_errmsg_internal(m, int32(_a_F_halfvec_to_sparsevec_2), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_halfvec_to_sparsevec_3), int32(676), int32(_a_F_halfvec_to_sparsevec_4))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v23
	F_errmsg(m, int32(_a_F_halfvec_to_sparsevec_5), v16)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_halfvec_to_sparsevec_3), int32(62), int32(_a_F_halfvec_to_sparsevec_6))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
