package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__int_contains(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_copy(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			if v18 != 0 {
				v19 = F_array_contains_nulls(m, v11)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(152590), int32(0))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495732), int32(38), int32(149284))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
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
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v21 != 0 {
							v22 = F_array_contains_nulls(m, v16)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								if v22 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(152590), int32(0))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495732), int32(39), int32(149284))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
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
									v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
									mBase = m.M
									v28 = m.ExcPending
									if v28 != 0 {
										return int32(0)
									} else {
										v29 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
										v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										if v31 != 0 {
											v39 = v31
										} else {
											v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
											v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v39+v11, v27, v8+int32(15))
										mBase = m.M
										v44 = F__int_unique(m, v11)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
											v49 = F_ArrayGetNItems(m, v46, v16+int32(16))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												v51 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v51)
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
												if v53 != 0 {
													v61 = v53
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
													v61 = (v54<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v61+v16, v49, v8+int32(14))
												mBase = m.M
												v66 = F__int_unique(m, v16)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v68 = F_inner_int_contains(m, v44, v66)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v44)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v66)
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(16)
																return v68
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								if v31 != 0 {
									v39 = v31
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v39+v11, v27, v8+int32(15))
								mBase = m.M
								v44 = F__int_unique(m, v11)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
									v49 = F_ArrayGetNItems(m, v46, v16+int32(16))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v51 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v51)
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
										if v53 != 0 {
											v61 = v53
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
											v61 = (v54<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v61+v16, v49, v8+int32(14))
										mBase = m.M
										v66 = F__int_unique(m, v16)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = F_inner_int_contains(m, v44, v66)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v44)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v66)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(16)
														return v68
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
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				if v21 != 0 {
					v22 = F_array_contains_nulls(m, v16)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						if v22 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(152590), int32(0))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495732), int32(39), int32(149284))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
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
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								if v31 != 0 {
									v39 = v31
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v39+v11, v27, v8+int32(15))
								mBase = m.M
								v44 = F__int_unique(m, v11)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
									v49 = F_ArrayGetNItems(m, v46, v16+int32(16))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v51 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v51)
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
										if v53 != 0 {
											v61 = v53
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
											v61 = (v54<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v61+v16, v49, v8+int32(14))
										mBase = m.M
										v66 = F__int_unique(m, v16)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = F_inner_int_contains(m, v44, v66)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v44)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v66)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(16)
														return v68
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						if v31 != 0 {
							v39 = v31
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v39+v11, v27, v8+int32(15))
						mBase = m.M
						v44 = F__int_unique(m, v11)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
							v49 = F_ArrayGetNItems(m, v46, v16+int32(16))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v51)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
								if v53 != 0 {
									v61 = v53
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
									v61 = (v54<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v61+v16, v49, v8+int32(14))
								mBase = m.M
								v66 = F__int_unique(m, v16)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = F_inner_int_contains(m, v44, v66)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v44)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v66)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(16)
												return v68
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
func F__int_different(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(6366), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int32(0))
	}
}
func F__int_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_copy(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			if v18 != 0 {
				v19 = F_array_contains_nulls(m, v11)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(152590), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495732), int32(131), int32(272026))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v21 != 0 {
							v22 = F_array_contains_nulls(m, v16)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								if v22 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(152590), int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495732), int32(132), int32(272026))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
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
									v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
									mBase = m.M
									v28 = m.ExcPending
									if v28 != 0 {
										return int32(0)
									} else {
										v29 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
										v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										if v31 != 0 {
											v39 = v31
										} else {
											v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
											v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v39+v11, v27, v8+int32(15))
										mBase = m.M
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
										v47 = F_ArrayGetNItems(m, v44, v16+int32(16))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											v49 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v49)
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
											if v51 != 0 {
												v59 = v51
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
												v59 = (v52<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v59+v16, v47, v8+int32(14))
											mBase = m.M
											v64 = F_inner_int_union(m, v11, v16)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v11)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v16)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(16)
														return v64
													}
												}
											}
										}
									}
								}
							}
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								if v31 != 0 {
									v39 = v31
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v39+v11, v27, v8+int32(15))
								mBase = m.M
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
								v47 = F_ArrayGetNItems(m, v44, v16+int32(16))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v49)
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
									if v51 != 0 {
										v59 = v51
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
										v59 = (v52<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									F_isort(m, v59+v16, v47, v8+int32(14))
									mBase = m.M
									v64 = F_inner_int_union(m, v11, v16)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v11)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v16)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(16)
												return v64
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				if v21 != 0 {
					v22 = F_array_contains_nulls(m, v16)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						if v22 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(152590), int32(0))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495732), int32(132), int32(272026))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
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
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								if v31 != 0 {
									v39 = v31
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v39+v11, v27, v8+int32(15))
								mBase = m.M
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
								v47 = F_ArrayGetNItems(m, v44, v16+int32(16))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v49)
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
									if v51 != 0 {
										v59 = v51
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
										v59 = (v52<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									F_isort(m, v59+v16, v47, v8+int32(14))
									mBase = m.M
									v64 = F_inner_int_union(m, v11, v16)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v11)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v16)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(16)
												return v64
											}
										}
									}
								}
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v27 = F_ArrayGetNItems(m, v24, v11+int32(16))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v29)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						if v31 != 0 {
							v39 = v31
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v39+v11, v27, v8+int32(15))
						mBase = m.M
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
						v47 = F_ArrayGetNItems(m, v44, v16+int32(16))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v49)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
							if v51 != 0 {
								v59 = v51
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
								v59 = (v52<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							F_isort(m, v59+v16, v47, v8+int32(14))
							mBase = m.M
							v64 = F_inner_int_union(m, v11, v16)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v16)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return v64
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
