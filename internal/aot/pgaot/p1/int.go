package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__int_inter(m *base.Module, l0 int32) int32 {
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
								F_errmsg(m, int32(151793), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(492670), int32(152), int32(214496))
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
											F_errmsg(m, int32(151793), int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492670), int32(153), int32(214496))
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
											v64 = F_inner_int_inter(m, v11, v16)
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
									v64 = F_inner_int_inter(m, v11, v16)
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
									F_errmsg(m, int32(151793), int32(0))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(492670), int32(153), int32(214496))
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
									v64 = F_inner_int_inter(m, v11, v16)
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
							v64 = F_inner_int_inter(m, v11, v16)
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
func F__int_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_copy(m, v14)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum_copy(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L48
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L6:
	;
	v23 = F_array_contains_nulls(m, v15)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v23 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v26 = F_array_contains_nulls(m, v20)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v30 = v15 + int32(16)
	v31 = F_ArrayGetNItems(m, v28, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	if v26 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v35 = v20 + int32(16)
	v36 = F_ArrayGetNItems(m, v33, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v38 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L20
L19:
	;
	v48 = v38
	goto L20
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v49 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v59 = (v52<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L23
L22:
	;
	v59 = v49
	goto L23
L23:
	;
	if v36 != v31 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_pfree(m, v15)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L42
	}
L25:
	;
	v136 = int32(0)
	goto L24
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v62 = F_ArrayGetNItems(m, v61, v30)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v64)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = v66
	goto L30
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L30
L30:
	;
	F_isort(m, v74+v15, v62, v12+int32(15))
	mBase = m.M
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v80 = F_ArrayGetNItems(m, v79, v35)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v84 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v92 = v84
	goto L34
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v92 = (v85<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L34
L34:
	;
	F_isort(m, v92+v20, v80, v12+int32(14))
	mBase = m.M
	v97 = int32(0)
	if v31 <= v97 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v136 = int32(1)
	goto L24
L36:
	;
	goto L37
L37:
	;
	v103 = v97
	goto L38
L38:
	;
	v113 = v103 << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v15+v48+v113)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+(v20+v59))))
	if v115 != v117 {
		goto L25
	} else {
		goto L40
	}
L39:
	;
	v136 = v119
	goto L24
L40:
	;
	v119 = int32(1)
	v121 = v103 + v119
	if v31 != v121 {
		v103 = v121
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	F_pfree(m, v20)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	m.G0 = v12 + int32(16)
	return v136
L44:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(151793), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(492670), int32(68), int32(374195))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(151793), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(492670), int32(69), int32(374195))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_int_query_opr_selec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 float32
	_ = v34
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v95 float64
	_ = v95
	var v98 int32
	_ = v98
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v120 float64
	_ = v120
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	F_check_stack_depth(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return float64(0)
	} else {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
		switch v18 - int32(2) {
		case 0:
			if l1 == int32(0) {
				v120 = float64(0.005)
				m.G0 = v12 + int32(32)
				return v120
			} else {
				v24 = int32(4)
				v28 = F_bsearch(m, l0+v24, l1, l3, v24, int32(6947))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return float64(0)
				} else {
					if v28 == int32(0) {
						v92 = float64(0.005)
						v95 = base.F64_promote_f32(base.F32_mul(l4, float32(0.5)))
						if base.F64_gt(v95, v92) != 0 {
							v120 = v92
						} else {
							v109 = v95
							v111 = float64(0)
							if base.F64_lt(v109, v111) != 0 {
								v120 = v111
							} else {
								if base.F64_gt(v109, float64(1)) == int32(0) {
									v120 = v109
								} else {
									v120 = float64(1)
								}
							}
						}
					} else {
						v34 = *(*float32)(unsafe.Add(mBase, uint32(l2+(v28-l1))))
						v109 = base.F64_promote_f32(v34)
						v111 = float64(0)
						if base.F64_lt(v109, v111) != 0 {
							v120 = v111
						} else {
							if base.F64_gt(v109, float64(1)) == int32(0) {
								v120 = v109
							} else {
								v120 = float64(1)
							}
						}
					}
					m.G0 = v12 + int32(32)
					return v120
				}
			}
		case 1:
			v38 = F_int_query_opr_selec(m, l0-int32(8), l1, l2, l3, l4)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return float64(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				switch v40 - int32(33) {
				case 0:
					v109 = base.F64_sub(float64(1), v38)
					v111 = float64(0)
					if base.F64_lt(v109, v111) != 0 {
						v120 = v111
					} else {
						if base.F64_gt(v109, float64(1)) == int32(0) {
							v120 = v109
						} else {
							v120 = float64(1)
						}
					}
					m.G0 = v12 + int32(32)
					return v120
				case 1, 2, 3, 4:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return float64(0)
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v58
						F_errmsg_internal(m, int32(478970), v12+int32(16))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(491618), int32(317), int32(488726))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 5:
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					v49 = F_int_query_opr_selec(m, l0+v45<<(uint(int32(3))%32), l1, l2, l3, l4)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return float64(0)
					} else {
						v109 = base.F64_mul(v38, v49)
						v111 = float64(0)
						if base.F64_lt(v109, v111) != 0 {
							v120 = v111
						} else {
							if base.F64_gt(v109, float64(1)) == int32(0) {
								v120 = v109
							} else {
								v120 = float64(1)
							}
						}
						m.G0 = v12 + int32(32)
						return v120
					}
				default:
					if v40 == int32(124) {
						v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						v102 = F_int_query_opr_selec(m, l0+v98<<(uint(int32(3))%32), l1, l2, l3, l4)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return float64(0)
						} else {
							v109 = base.F64_sub(base.F64_add(v38, v102), base.F64_mul(v38, v102))
							v111 = float64(0)
							if base.F64_lt(v109, v111) != 0 {
								v120 = v111
							} else {
								if base.F64_gt(v109, float64(1)) == int32(0) {
									v120 = v109
								} else {
									v120 = float64(1)
								}
							}
							m.G0 = v12 + int32(32)
							return v120
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return float64(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v58
							F_errmsg_internal(m, int32(478970), v12+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(491618), int32(317), int32(488726))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return float64(0)
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
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return float64(0)
			} else {
				v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v78
				F_errmsg_internal(m, int32(58452), v12)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return float64(0)
				} else {
					F_errfinish(m, int32(491618), int32(324), int32(488726))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return float64(0)
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
func F_int_to_intset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(28))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = int64(4294967297)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(98784247808)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4294967408)
		return v4
	}
}
func F_int_to_roman(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v12 = F_palloc(m, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v16)
		if base.Ui32(l0-int32(4000)) <= base.Ui32(int32(-4000)) {
			v22 = int64(2531906049332683555)
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v12)+7)) = v22
			v26 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v26)
			m.G0 = v9 + int32(16)
			return v12
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
			v33 = F_pg_snprintf(m, v9+int32(4), int32(12), int32(485191), v9)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
				if v35 == int32(0) {
				} else {
					v40 = v35
					v44 = v33
					v45 = v9 + int32(4)
					for {
						v46 = base.I32_extend8_s(v40)
						if v46 < int32(49) {
						} else {
							v50 = v46 - int32(49)
							switch v44 - int32(1) {
							case 0:
								v70 = int32(1649664)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v50<<(uint(int32(2))%32))))
								v75 = F_strlen(m, v12)
								mBase = m.M
								v77 = F_strcpy(m, v75+v12, v74)
								mBase = m.M
							case 1:
								v70 = int32(1649616)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v50<<(uint(int32(2))%32))))
								v75 = F_strlen(m, v12)
								mBase = m.M
								v77 = F_strcpy(m, v75+v12, v74)
								mBase = m.M
							case 2:
								v70 = int32(1649568)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v50<<(uint(int32(2))%32))))
								v75 = F_strlen(m, v12)
								mBase = m.M
								v77 = F_strcpy(m, v75+v12, v74)
								mBase = m.M
							case 3:
								v54 = v50
								for {
									v60 = F_strlen(m, v12)
									mBase = m.M
									v62 = int32(77)
									*(*uint16)(unsafe.Add(mBase, uint32(v60+v12))) = uint16(v62)
									if int32(0) < v54 {
										v54 = v54 - int32(1)
										continue
									} else {
										break
									}
									break
								}
							default:
							}
						}
						v84 = int32(1)
						v87 = v45 + v84
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
						if v88 != 0 {
							v40 = v88
							v44 = v44 - v84
							v45 = v87
							continue
						} else {
							break
						}
						break
					}
				}
				m.G0 = v9 + int32(16)
				return v12
			}
		}
	}
}
func F_parseIntFromText(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = F_strlen(m, l0)
	mBase = m.M
	if v11 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L20
	} else {
		goto L36
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L32
	}
L3:
	;
	if v55 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v55 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = v10
	v19 = l0
	v20 = v11
	v21 = v17
	goto L11
L8:
	;
	v43 = l0
	v47 = int32(0)
	goto L9
L9:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v55 = v47 - v48
	goto L3
L10:
	;
	v43 = v38
	v47 = v40
	goto L9
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 != v23 {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v38 = v32
	v40 = int32(0)
	goto L10
L13:
	;
	if v23 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v28 = v20 - int32(1)
	if v28 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v31 = int32(1)
	v32 = v19 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v33 != 0 {
		v18 = v18 + v31
		v19 = v32
		v20 = v28
		v21 = v33
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v6 + int32(-4)
	v61 = v11 + v10
	v65 = F_sscanf(m, v61, int32(485191), v6+int32(-32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L20
	} else {
		goto L28
	}
L20:
	;
	return int32(0)
L21:
	;
	if v65 != int32(1) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v71 = int32(10)
	v72 = F___strchrnul(m, v61, v71)
	mBase = m.M
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 == v71 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v78 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v78 = v72
	goto L26
L25:
	;
	v78 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	m.G0 = v8 - int32(-64)
	return v84
L28:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l2
	F_errmsg(m, int32(698005), v6+int32(-16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(492015), int32(1314), int32(63683))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
	F_errmsg(m, int32(698005), v6+int32(-48))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(492015), int32(1319), int32(63683))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
	F_errmsg(m, int32(698005), v8)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(492015), int32(1324), int32(63683))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v24 int64
	_ = v24
	var v26 float64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 float64
	_ = v168
	var v172 float64
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v199 float64
	_ = v199
	var v201 float64
	_ = v201
	var v202 float64
	_ = v202
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v18
	v24 = F_strtox_2(m, l0, v11+int32(4), v18, int64(2147483648))
	mBase = m.M
	goto L7
L7:
	;
	v26 = base.F64_convert_i32_s(base.I32_wrap_i64(v24))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 == int32(46) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v53 = int32(0)
	if l0 == v51 {
		v240 = v53
		goto L16
	} else {
		goto L17
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v45 = F_strtod(m, l0, v11+int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	if v29 == int32(69) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v29 == int32(101) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v37 != int32(68) {
		v51 = v28
		v52 = v26
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	return int32(0)
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v51 = v50
	v52 = v45
	goto L8
L16:
	;
	m.G0 = v11 + int32(16)
	return v240
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v56 == int32(68) {
		v240 = v53
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v52)&int64(9223372036854775807)) {
		v240 = v53
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v69 = v51
	goto L20
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if base.Ui32(v72-int32(9)) < base.Ui32(int32(5)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = v69 + int32(1)
	goto L20
L23:
	;
	if v72 == int32(32) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v222 = int32(1)
	if l1 == int32(0) {
		v240 = v222
		goto L16
	} else {
		goto L63
	}
L26:
	;
	if l3 == int32(0) {
		v240 = v53
		goto L16
	} else {
		goto L59
	}
L27:
	;
	v80 = l2 & int32(2130706432)
	if v80 == int32(0) {
		v240 = v53
		goto L16
	} else {
		goto L30
	}
L28:
	;
	v201 = v52
	goto L29
L29:
	;
	v202 = base.F64_nearest(v201)
	if base.F64_gt(v202, float64(2.147483647e+09))|base.F64_lt(v202, float64(-2.147483648e+09)) == int32(0) {
		goto L25
	} else {
		goto L57
	}
L30:
	;
	v90 = m.G0
	v92 = v90 - int32(16)
	m.G0 = v92
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	switch v96 {
	case 0, 9, 10, 11, 12, 13, 32:
		v115 = v69
		v116 = v92 + int32(12)
		goto L32
	default:
		goto L33
	}
L31:
	;
	if v187 == int32(0) {
		goto L26
	} else {
		goto L56
	}
L32:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v117)
	v126 = v115
	goto L36
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+12)) = uint8(v96)
	v101 = v69 + int32(1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	switch v102 {
	case 0, 9, 10, 11, 12, 13, 32:
		v115 = v101
		v116 = v92 + int32(13)
		goto L32
	default:
		goto L34
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+13)) = uint8(v102)
	v107 = v69 + int32(2)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	switch v108 {
	case 0, 9, 10, 11, 12, 13, 32:
		v115 = v107
		v116 = v92 + int32(14)
		goto L32
	default:
		goto L35
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+14)) = uint8(v108)
	v115 = v69 + int32(3)
	v116 = v92 + int32(15)
	goto L32
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if base.Ui32(v129-int32(9)) < base.Ui32(int32(5)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = v126 + int32(1)
	goto L36
L39:
	;
	if v129 == int32(32) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if v129 != 0 {
		v187 = v117
		goto L41
	} else {
		goto L42
	}
L41:
	;
	m.G0 = v92 + int32(16)
	goto L31
L42:
	;
	if v80&int32(251658240) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v140 = int32(1747136)
	goto L45
L44:
	;
	v140 = int32(1747552)
	goto L45
L45:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v141 == int32(0) {
		v187 = v117
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v148 = v117
	goto L47
L47:
	;
	v155 = v140 + v148<<(uint(int32(4))%32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v80 != v156 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v187 = int32(0)
	goto L41
L49:
	;
	v177 = v148 + int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v177<<(uint(int32(4))%32)))))
	if v181 != 0 {
		v148 = v177
		goto L47
	} else {
		goto L55
	}
L50:
	;
	v160 = F_strcmp(m, v92+int32(12), v155)
	mBase = m.M
	if v160 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v155)+8))
	v162 = base.F64_mul(v52, v161)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+16)))
	if v163 == int32(0) {
		v172 = v162
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v172
	v187 = int32(1)
	goto L41
L53:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	if v80 != v166 {
		v172 = v162
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v155)+24))
	v172 = base.F64_mul(v168, base.F64_nearest(base.F64_div(v162, v168)))
	goto L52
L55:
	;
	goto L48
L56:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v201 = v199
	goto L29
L57:
	;
	if l3 == int32(0) {
		v240 = v53
		goto L16
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(623379)
	v240 = v53
	goto L16
L59:
	;
	if l2&int32(251658240) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(650991)
	v240 = v53
	goto L16
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(650612)
	v240 = v53
	goto L16
L63:
	;
	if base.F64_lt(base.F64_abs(v202), float64(2.147483648e+09)) != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v230
	v240 = v222
	goto L16
L65:
	;
	v228 = base.I32_trunc_f64_s(v202)
	v230 = v228
	goto L64
L66:
	;
	goto L67
L67:
	;
	v230 = int32(-2147483648)
	goto L64
}
