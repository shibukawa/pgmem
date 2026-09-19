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
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
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
								F_errmsg(m, int32(_a_F__int_inter_0), int32(0))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F__int_inter_1), int32(152), int32(_a_F__int_inter_2))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F__int_inter_0), int32(0))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F__int_inter_1), int32(153), int32(_a_F__int_inter_2))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
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
									v27 = F_ArrayGetNItemsSafe(m, v24, v11+int32(16))
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
										v47 = F_ArrayGetNItemsSafe(m, v44, v16+int32(16))
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
							v27 = F_ArrayGetNItemsSafe(m, v24, v11+int32(16))
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
								v47 = F_ArrayGetNItemsSafe(m, v44, v16+int32(16))
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
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F__int_inter_0), int32(0))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F__int_inter_1), int32(153), int32(_a_F__int_inter_2))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
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
							v27 = F_ArrayGetNItemsSafe(m, v24, v11+int32(16))
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
								v47 = F_ArrayGetNItemsSafe(m, v44, v16+int32(16))
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
					v27 = F_ArrayGetNItemsSafe(m, v24, v11+int32(16))
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
						v47 = F_ArrayGetNItemsSafe(m, v44, v16+int32(16))
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
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
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
	v169 = m.ExcPending
	if v169 != 0 {
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
	v31 = F_ArrayGetNItemsSafe(m, v28, v30)
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
	v36 = F_ArrayGetNItemsSafe(m, v33, v35)
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
	if v31 != v36 {
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
	v62 = F_ArrayGetNItemsSafe(m, v61, v30)
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
	v80 = F_ArrayGetNItemsSafe(m, v79, v35)
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
	F_errmsg(m, int32(_a_F__int_same_0), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F__int_same_1), int32(68), int32(_a_F__int_same_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
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
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F__int_same_0), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F__int_same_1), int32(69), int32(_a_F__int_same_2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v90 int32
	_ = v90
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v113 float64
	_ = v113
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
				v113 = float64(0.005)
				m.G0 = v12 + int32(32)
				return v113
			} else {
				v24 = int32(4)
				v28 = F_bsearch(m, l0+v24, l1, l3, v24, int32(_a_F_int_query_opr_selec_0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return float64(0)
				} else {
					if v28 == int32(0) {
						v84 = float64(0.005)
						v87 = base.F64_promote_f32(base.F32_mul(l4, float32(0.5)))
						if base.F64_gt(v87, v84) != 0 {
							v113 = v84
						} else {
							v100 = v87
							v103 = float64(0)
							if base.F64_lt(v100, v103) != 0 {
								v113 = v103
							} else {
								if base.F64_gt(v100, float64(1)) == int32(0) {
									v113 = v100
								} else {
									v113 = float64(1)
								}
							}
						}
					} else {
						v34 = *(*float32)(unsafe.Add(mBase, uint32(l2+(v28-l1))))
						v100 = base.F64_promote_f32(v34)
						v103 = float64(0)
						if base.F64_lt(v100, v103) != 0 {
							v113 = v103
						} else {
							if base.F64_gt(v100, float64(1)) == int32(0) {
								v113 = v100
							} else {
								v113 = float64(1)
							}
						}
					}
					m.G0 = v12 + int32(32)
					return v113
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
					v100 = base.F64_sub(float64(1), v38)
					v103 = float64(0)
					if base.F64_lt(v100, v103) != 0 {
						v113 = v103
					} else {
						if base.F64_gt(v100, float64(1)) == int32(0) {
							v113 = v100
						} else {
							v113 = float64(1)
						}
					}
					m.G0 = v12 + int32(32)
					return v113
				case 1, 2, 3, 4:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return float64(0)
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v58
						F_errmsg_internal(m, int32(_a_F_int_query_opr_selec_1), v12+int32(16))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(_a_F_int_query_opr_selec_2), int32(317), int32(_a_F_int_query_opr_selec_3))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
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
						v100 = base.F64_mul(v38, v49)
						v103 = float64(0)
						if base.F64_lt(v100, v103) != 0 {
							v113 = v103
						} else {
							if base.F64_gt(v100, float64(1)) == int32(0) {
								v113 = v100
							} else {
								v113 = float64(1)
							}
						}
						m.G0 = v12 + int32(32)
						return v113
					}
				default:
					if v40 == int32(124) {
						v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						v94 = F_int_query_opr_selec(m, l0+v90<<(uint(int32(3))%32), l1, l2, l3, l4)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return float64(0)
						} else {
							v100 = base.F64_sub(base.F64_add(v38, v94), base.F64_mul(v38, v94))
							v103 = float64(0)
							if base.F64_lt(v100, v103) != 0 {
								v113 = v103
							} else {
								if base.F64_gt(v100, float64(1)) == int32(0) {
									v113 = v100
								} else {
									v113 = float64(1)
								}
							}
							m.G0 = v12 + int32(32)
							return v113
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
							F_errmsg_internal(m, int32(_a_F_int_query_opr_selec_1), v12+int32(16))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_int_query_opr_selec_2), int32(317), int32(_a_F_int_query_opr_selec_3))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
			v73 = m.ExcPending
			if v73 != 0 {
				return float64(0)
			} else {
				v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v74
				F_errmsg_internal(m, int32(_a_F_int_query_opr_selec_4), v12)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return float64(0)
				} else {
					F_errfinish(m, int32(_a_F_int_query_opr_selec_2), int32(324), int32(_a_F_int_query_opr_selec_3))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
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
			*(*int64)(unsafe.Add(mBase, uint32(v12)+7)) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v22
			v26 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v26)
			m.G0 = v9 + int32(16)
			return v12
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
			v30 = v9 + int32(4)
			v33 = F_pg_snprintf(m, v30, int32(12), int32(_a_F_int_to_roman_0), v9)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
				if v35 == int32(0) {
				} else {
					v38 = v35
					v39 = v30
					v42 = v33
					for {
						v44 = base.I32_extend8_s(v38)
						if v44 < int32(49) {
						} else {
							v48 = v44 - int32(49)
							switch v42 - int32(1) {
							case 0:
								v68 = int32(_a_F_int_to_roman_1)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v48<<(uint(int32(2))%32))))
								v73 = F_strlen(m, v12)
								mBase = m.M
								v75 = F_strcpy(m, v73+v12, v72)
								mBase = m.M
							case 1:
								v68 = int32(_a_F_int_to_roman_2)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v48<<(uint(int32(2))%32))))
								v73 = F_strlen(m, v12)
								mBase = m.M
								v75 = F_strcpy(m, v73+v12, v72)
								mBase = m.M
							case 2:
								v68 = int32(_a_F_int_to_roman_3)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v48<<(uint(int32(2))%32))))
								v73 = F_strlen(m, v12)
								mBase = m.M
								v75 = F_strcpy(m, v73+v12, v72)
								mBase = m.M
							case 3:
								v52 = v48
								for {
									v58 = F_strlen(m, v12)
									mBase = m.M
									v60 = int32(77)
									*(*uint16)(unsafe.Add(mBase, uint32(v58+v12))) = uint16(v60)
									if int32(0) < v52 {
										v52 = v52 - int32(1)
										continue
									} else {
										break
									}
									break
								}
							default:
							}
						}
						v82 = int32(1)
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
						if v86 != 0 {
							v38 = v86
							v39 = v39 + v82
							v42 = v42 - v82
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13956(m, l0, l1, l2, int32(_a_F_parseIntFromText_0), int32(1324), int32(1319), int32(1314), int32(_a_F_parseIntFromText_1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
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
	var v41 int32
	_ = v41
	var v49 float64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 float64
	_ = v179
	var v183 float64
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v208 float64
	_ = v208
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
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
	*(*int32)(unsafe.Add(mBase, _c_F_parse_int[0])) = v18
	v24 = F_strtox_2(m, l0, v11+int32(4), v18, int64(2147483648))
	mBase = m.M
	goto L7
L7:
	;
	v26 = base.F64_convert_i32_s(base.I32_wrap_i64(v24))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.B2i32(v29 == int32(46))|base.B2i32(v29 == int32(69))|base.B2i32(v29 == int32(101)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v57 = int32(0)
	if l0 == v55 {
		v242 = v57
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_parse_int[0]))
	if v41 != int32(68) {
		v55 = v28
		v56 = v26
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_parse_int[0])) = int32(0)
	v49 = F_strtod(m, l0, v11+int32(4))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v49
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v55 = v54
	v56 = v49
	goto L8
L15:
	;
	m.G0 = v11 + int32(16)
	return v242
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_parse_int[0]))
	if base.B2i32(v60 == int32(68))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
		v242 = v57
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v74 = v55
	goto L18
L18:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if base.B2i32(base.Ui32(v77-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v77 == int32(32)) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v242 = v231
	goto L15
L20:
	;
	v74 = v74 + int32(1)
	goto L18
L21:
	;
	if v77 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L19
L23:
	;
	v231 = int32(1)
	if l1 == int32(0) {
		v242 = v231
		goto L15
	} else {
		goto L61
	}
L24:
	;
	if l3 == int32(0) {
		v242 = v57
		goto L15
	} else {
		goto L57
	}
L25:
	;
	v88 = l2 & int32(2130706432)
	if v88 == int32(0) {
		v242 = v57
		goto L15
	} else {
		goto L28
	}
L26:
	;
	v210 = v56
	goto L27
L27:
	;
	v211 = base.F64_nearest(v210)
	if base.F64_gt(v211, float64(2.147483647e+09))|base.F64_lt(v211, float64(-2.147483648e+09)) == int32(0) {
		goto L23
	} else {
		goto L55
	}
L28:
	;
	v98 = m.G0
	v100 = v98 - int32(16)
	m.G0 = v100
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	switch v104 {
	case 0, 9, 10, 11, 12, 13, 32:
		v122 = v74
		v123 = v100 + int32(12)
		goto L30
	default:
		goto L31
	}
L29:
	;
	if v195 == int32(0) {
		goto L24
	} else {
		goto L54
	}
L30:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v125)
	v132 = v122
	goto L34
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+12)) = uint8(v104)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	switch v110 {
	case 0, 9, 10, 11, 12, 13, 32:
		v122 = v74 + int32(1)
		v123 = v100 + int32(13)
		goto L30
	default:
		goto L32
	}
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+13)) = uint8(v110)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+2)))
	switch v116 {
	case 0, 9, 10, 11, 12, 13, 32:
		v122 = v74 + int32(2)
		v123 = v100 + int32(14)
		goto L30
	default:
		goto L33
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+14)) = uint8(v116)
	v122 = v74 + int32(3)
	v123 = v100 + int32(15)
	goto L30
L34:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if base.B2i32(base.Ui32(v137-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v137 == int32(32)) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L29
L36:
	;
	v132 = v132 + int32(1)
	goto L34
L37:
	;
	if v137 != 0 {
		v195 = v125
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L35
L39:
	;
	m.G0 = v100 + int32(16)
	goto L38
L40:
	;
	if v88&int32(251658240) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v151 = int32(_a_F_parse_int_0)
	goto L43
L42:
	;
	v151 = int32(_a_F_parse_int_1)
	goto L43
L43:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v152 == int32(0) {
		v195 = v125
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v156 = v125
	goto L45
L45:
	;
	v166 = v151 + v156<<(uint(int32(4))%32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v88 != v167 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v195 = int32(0)
	goto L39
L47:
	;
	v188 = v156 + int32(1)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v188<<(uint(int32(4))%32)))))
	if v192 != 0 {
		v156 = v188
		goto L45
	} else {
		goto L53
	}
L48:
	;
	v171 = F_strcmp(m, v100+int32(12), v166)
	mBase = m.M
	if v171 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v166)+8))
	v173 = base.F64_mul(v56, v172)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+16)))
	if v174 == int32(0) {
		v183 = v173
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v183
	v195 = int32(1)
	goto L39
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v166)+20))
	if v88 != v177 {
		v183 = v173
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v166)+24))
	v183 = base.F64_mul(v179, base.F64_nearest(base.F64_div(v173, v179)))
	goto L50
L53:
	;
	goto L46
L54:
	;
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v210 = v208
	goto L27
L55:
	;
	if l3 == int32(0) {
		v242 = v57
		goto L15
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_parse_int_2)
	v242 = v57
	goto L15
L57:
	;
	if l2&int32(251658240) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_parse_int_3)
	v242 = v57
	goto L15
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_parse_int_4)
	v242 = v57
	goto L15
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_trunc_sat_f64_s(v211)
	goto L22
}
