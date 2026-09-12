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
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 float32
	_ = v205
	var v207 int32
	_ = v207
	var v214 float32
	_ = v214
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
			v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v16))))
			v116 = v111 & int32(1023)
			v120 = v111 << (uint(int32(16)) % 32) & int32(-2147483648)
			v123 = int32(31)
			v124 = int32(base.Ui32(v111)>>(uint(int32(10))%32)) & v123
			if v124 != v123 {
				if v124 != 0 {
					v196 = v116
					v197 = v124<<(uint(int32(23))%32) + v120 + int32(939524096)
				} else {
					if v116 != 0 {
						if v111&int32(512) != 0 {
							v184 = v116 << (uint(int32(1)) % 32)
							v186 = int32(939524096)
						} else {
							if base.Ui32(int32(255)) < base.Ui32(v116) {
								v184 = v116 << (uint(int32(2)) % 32)
								v186 = int32(931135488)
							} else {
								if base.Ui32(int32(127)) < base.Ui32(v116) {
									v184 = v116 << (uint(int32(3)) % 32)
									v186 = int32(922746880)
								} else {
									if base.Ui32(int32(63)) < base.Ui32(v116) {
										v184 = v116 << (uint(int32(4)) % 32)
										v186 = int32(914358272)
									} else {
										if base.Ui32(int32(31)) < base.Ui32(v116) {
											v184 = v116 << (uint(int32(5)) % 32)
											v186 = int32(905969664)
										} else {
											if base.Ui32(int32(15)) < base.Ui32(v116) {
												v184 = v116 << (uint(int32(6)) % 32)
												v186 = int32(897581056)
											} else {
												if base.Ui32(int32(7)) < base.Ui32(v116) {
													v184 = v116 << (uint(int32(7)) % 32)
													v186 = int32(889192448)
												} else {
													if base.Ui32(int32(3)) < base.Ui32(v116) {
														v184 = v116 << (uint(int32(8)) % 32)
														v186 = int32(880803840)
													} else {
														v179 = base.B2i32(v116 == int32(1))
														if v116 == int32(1) {
															v180 = int32(1024)
														} else {
															v180 = v116 << (uint(int32(9)) % 32)
														}
														if v116 == int32(1) {
															v183 = int32(864026624)
														} else {
															v183 = int32(872415232)
														}
														v184 = v180
														v186 = v183
													}
												}
											}
										}
									}
								}
							}
						}
						v196 = v184 & int32(1022)
						v197 = v186 | v120
					} else {
						v196 = int32(0)
						v197 = v120
					}
				}
			} else {
				if v116 == int32(0) {
					v196 = int32(0)
					v197 = v120 | int32(2139095040)
				} else {
					v196 = v116
					v197 = v120 | int32(2143289344)
				}
			}
			v205 = base.F32_add(base.F32_abs(base.F32_sub(base.F32_reinterpret_i32(v104|v103<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v197|v196<<(uint(int32(13))%32)))), v14)
			v207 = v12 + int32(1)
			if v207 != l0 {
				v12 = v207
				v14 = v205
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
	var v82 int32
	_ = v82
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
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+(v9+v36)))))
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
									v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
									v82 = int32(1)
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v68+v73<<(uint(v82)%32)))))
									*(*uint16)(unsafe.Add(mBase, uint32(v27+v68+(v73+v80)<<(uint(v82)%32)))) = uint16(v88)
									v91 = v73 + v82
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
						F_errmsg(m, int32(488250), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512461), int32(80), int32(154322))
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
				v45 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
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
						F_errmsg(m, int32(488250), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512461), int32(80), int32(154322))
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
				v45 = *(*int32)(unsafe.Add(mBase, _consts[1403]))
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
						F_errmsg(m, int32(488250), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512461), int32(80), int32(154322))
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
				v45 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
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
	F_CheckElement_1(m, v61&int32(65535))
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
	F_errmsg(m, int32(477735), v10+int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(512461), int32(92), int32(294762))
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
	F_errmsg(m, int32(477766), v10)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(512461), int32(390), int32(37691))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	F_CheckDim_2(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.B2i32(v24 != int32(-1))&base.B2i32(v24 != v25) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = base.B2i32(v25 <= int32(0))
	if v25 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L83
	}
L7:
	;
	F_CheckNnz(m, v136, v25)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L21
	}
L8:
	;
	v136 = v2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v37 = v25 & int32(3)
	v39 = v20 + int32(8)
	v40 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v25) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v40
	v46 = v2
	v47 = v2
	goto L14
L12:
	;
	v91 = v40
	v92 = v2
	goto L13
L13:
	;
	if v37 == int32(0) {
		v136 = v92
		goto L7
	} else {
		goto L17
	}
L14:
	;
	v61 = v39 + v45<<(uint(int32(1))%32)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
	v63 = int32(32767)
	v65 = int32(0)
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+2)))
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+6)))
	v85 = v46 + base.B2i32(v62&v63 != v65) + base.B2i32(v68&v63 != v65) + base.B2i32(v74&v63 != v65) + base.B2i32(v80&v63 != v65)
	v86 = int32(4)
	v87 = v45 + v86
	v89 = v47 + v86
	if v89 != v25&int32(32764) {
		v45 = v87
		v46 = v85
		v47 = v89
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v91 = v87
	v92 = v85
	goto L13
L16:
	;
	goto L15
L17:
	;
	v107 = v91
	v108 = v92
	v110 = v2
	goto L18
L18:
	;
	v121 = int32(1)
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v107<<(uint(v121)%32)))))
	v129 = v108 + base.B2i32(v124&int32(32767) != int32(0))
	v133 = v110 + v121
	if v133 != v37 {
		v107 = v107 + v121
		v108 = v129
		v110 = v133
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v136 = v129
	goto L7
L20:
	;
	goto L19
L21:
	;
	v153 = F_mul_size(m, int32(4), v136)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v155 = F_add_size(m, int32(16), v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v158 = F_mul_size(m, int32(4), v136)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v160 = F_add_size(m, v155, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v162 = F_palloc0(m, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v160 << (uint(int32(2)) % 32)
	if v35 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L80
	}
L28:
	;
	v172 = v162 + int32(16)
	v178 = int32(0)
	v180 = v178
	v185 = v178
	goto L31
L29:
	;
	goto L30
L30:
	;
	m.G0 = v17 + int32(16)
	return v162
L31:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(8)+v180<<(uint(int32(1))%32)))))
	if v197&int32(32767) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	if v200 <= v185 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	v299 = v185
	goto L35
L35:
	;
	v303 = v180 + int32(1)
	if v303 != v25 {
		v180 = v303
		v185 = v299
		goto L31
	} else {
		goto L79
	}
L36:
	;
	v203 = v185 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v172+v203))) = v180
	v208 = base.I32_extend16_s(v197) & int32(-2147483648)
	v210 = v197 & int32(1023)
	v213 = int32(31)
	v214 = int32(base.Ui32(v197)>>(uint(int32(10))%32)) & v213
	if v214 != v213 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203+(v172+v136<<(uint(int32(2))%32))))) = v287 | v286<<(uint(int32(13))%32)
	v299 = v185 + int32(1)
	goto L35
L38:
	;
	v286 = v210
	v287 = v214<<(uint(int32(23))%32) + v208 + int32(939524096)
	goto L37
L39:
	;
	if v197&int32(512) != 0 {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	if v214 != 0 {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v210 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v210 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v286 = int32(0)
	v287 = v208
	goto L37
L45:
	;
	v286 = int32(0)
	v287 = v208 | int32(2139095040)
	goto L37
L46:
	;
	goto L47
L47:
	;
	v286 = v210
	v287 = v208 | int32(2143289344)
	goto L37
L48:
	;
	v286 = v274 & int32(1022)
	v287 = v276 | v208
	goto L37
L49:
	;
	v274 = v210 << (uint(int32(1)) % 32)
	v276 = int32(939524096)
	goto L48
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(int32(255)) < base.Ui32(v210) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v274 = v210 << (uint(int32(2)) % 32)
	v276 = int32(931135488)
	goto L48
L53:
	;
	goto L54
L54:
	;
	if base.Ui32(int32(127)) < base.Ui32(v210) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v274 = v210 << (uint(int32(3)) % 32)
	v276 = int32(922746880)
	goto L48
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(int32(63)) < base.Ui32(v210) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v274 = v210 << (uint(int32(4)) % 32)
	v276 = int32(914358272)
	goto L48
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(int32(31)) < base.Ui32(v210) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v274 = v210 << (uint(int32(5)) % 32)
	v276 = int32(905969664)
	goto L48
L62:
	;
	goto L63
L63:
	;
	if base.Ui32(int32(15)) < base.Ui32(v210) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v274 = v210 << (uint(int32(6)) % 32)
	v276 = int32(897581056)
	goto L48
L65:
	;
	goto L66
L66:
	;
	if base.Ui32(int32(7)) < base.Ui32(v210) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v274 = v210 << (uint(int32(7)) % 32)
	v276 = int32(889192448)
	goto L48
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(3)) < base.Ui32(v210) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v274 = v210 << (uint(int32(8)) % 32)
	v276 = int32(880803840)
	goto L48
L71:
	;
	goto L72
L72:
	;
	v269 = base.B2i32(v210 == int32(1))
	if v210 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v270 = int32(1024)
	goto L75
L74:
	;
	v270 = v210 << (uint(int32(9)) % 32)
	goto L75
L75:
	;
	if v210 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v273 = int32(864026624)
	goto L78
L77:
	;
	v273 = int32(872415232)
	goto L78
L78:
	;
	v274 = v270
	v276 = v273
	goto L48
L79:
	;
	goto L32
L80:
	;
	F_errmsg_internal(m, int32(464098), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(512592), int32(676), int32(502193))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v24
	F_errmsg(m, int32(477735), v17)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(512592), int32(62), int32(294762))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
