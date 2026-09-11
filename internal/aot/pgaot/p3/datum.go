package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumRestore(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5 + int32(4)
	if v6 == int32(-2) {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v16)
		if v6 == int32(-1) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20 + int32(4)
			return v21
		} else {
			v26 = F_palloc(m, v6)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v6 != 0 {
					v31 = F__emscripten_memcpy_bulkmem(m, v26, v30, v6)
					mBase = m.M
					v32 = v31
				} else {
					v32 = v26
				}
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v33 + v6
				return v32
			}
		}
	}
}
func F_datumSerialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v6 = int32(0)
	if l1 != 0 {
		v10 = int32(-2)
	} else {
		v10 = int32(-1)
	}
	if l1 != 0 {
		v27 = v10
		v28 = v6
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v29))) = v27
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		v33 = v31 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v33
		if l1 == int32(0) {
			if l2 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38 + int32(4)
				return
			} else {
				if v28 != 0 {
					v42 = F_palloc(m, v27)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_EOH_flatten_into(m, v28, v42, v27)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v27 != 0 {
								v47 = F__emscripten_memcpy_bulkmem(m, v46, v42, v27)
								mBase = m.M
							} else {
							}
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49 + v27
							F_pfree(m, v42)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					if v27 != 0 {
						v54 = F__emscripten_memcpy_bulkmem(m, v33, l0, v27)
						mBase = m.M
					} else {
					}
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56 + v27
					return
				}
			}
		} else {
			return
		}
	} else {
		if l2 != 0 {
			v27 = v10
			v28 = v6
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v27
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			v33 = v31 + int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v33
			if l1 == int32(0) {
				if l2 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38 + int32(4)
					return
				} else {
					if v28 != 0 {
						v42 = F_palloc(m, v27)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_EOH_flatten_into(m, v28, v42, v27)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v27 != 0 {
									v47 = F__emscripten_memcpy_bulkmem(m, v46, v42, v27)
									mBase = m.M
								} else {
								}
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49 + v27
								F_pfree(m, v42)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						if v27 != 0 {
							v54 = F__emscripten_memcpy_bulkmem(m, v33, l0, v27)
							mBase = m.M
						} else {
						}
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56 + v27
						return
					}
				}
			} else {
				return
			}
		} else {
			if l3 != int32(-1) {
				v25 = F_datumGetSize(m, l0, int32(0), l3)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = v25
					v28 = v6
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(v29))) = v27
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					v33 = v31 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v33
					if l1 == int32(0) {
						if l2 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38 + int32(4)
							return
						} else {
							if v28 != 0 {
								v42 = F_palloc(m, v27)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_EOH_flatten_into(m, v28, v42, v27)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										if v27 != 0 {
											v47 = F__emscripten_memcpy_bulkmem(m, v46, v42, v27)
											mBase = m.M
										} else {
										}
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49 + v27
										F_pfree(m, v42)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								if v27 != 0 {
									v54 = F__emscripten_memcpy_bulkmem(m, v33, l0, v27)
									mBase = m.M
								} else {
								}
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56 + v27
								return
							}
						}
					} else {
						return
					}
				}
			} else {
				v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v13 != int32(1) {
					v25 = F_datumGetSize(m, l0, int32(0), l3)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = v25
						v28 = v6
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v27
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						v33 = v31 + int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v33
						if l1 == int32(0) {
							if l2 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38 + int32(4)
								return
							} else {
								if v28 != 0 {
									v42 = F_palloc(m, v27)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_EOH_flatten_into(m, v28, v42, v27)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											if v27 != 0 {
												v47 = F__emscripten_memcpy_bulkmem(m, v46, v42, v27)
												mBase = m.M
											} else {
											}
											v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49 + v27
											F_pfree(m, v42)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									if v27 != 0 {
										v54 = F__emscripten_memcpy_bulkmem(m, v33, l0, v27)
										mBase = m.M
									} else {
									}
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56 + v27
									return
								}
							}
						} else {
							return
						}
					}
				} else {
					v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					if v16&int32(254) != int32(2) {
						v25 = F_datumGetSize(m, l0, int32(0), l3)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = v25
							v28 = v6
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = v27
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v33 = v31 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v33
							if l1 == int32(0) {
								if l2 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38 + int32(4)
									return
								} else {
									if v28 != 0 {
										v42 = F_palloc(m, v27)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											F_EOH_flatten_into(m, v28, v42, v27)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												if v27 != 0 {
													v47 = F__emscripten_memcpy_bulkmem(m, v46, v42, v27)
													mBase = m.M
												} else {
												}
												v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49 + v27
												F_pfree(m, v42)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													return
												}
											}
										}
									} else {
										if v27 != 0 {
											v54 = F__emscripten_memcpy_bulkmem(m, v33, l0, v27)
											mBase = m.M
										} else {
										}
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56 + v27
										return
									}
								}
							} else {
								return
							}
						}
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
						v22 = F_EOH_get_flat_size(m, v21)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v27 = v22
							v28 = v21
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = v27
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v33 = v31 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v33
							if l1 == int32(0) {
								if l2 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38 + int32(4)
									return
								} else {
									if v28 != 0 {
										v42 = F_palloc(m, v27)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											F_EOH_flatten_into(m, v28, v42, v27)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												if v27 != 0 {
													v47 = F__emscripten_memcpy_bulkmem(m, v46, v42, v27)
													mBase = m.M
												} else {
												}
												v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49 + v27
												F_pfree(m, v42)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													return
												}
											}
										}
									} else {
										if v27 != 0 {
											v54 = F__emscripten_memcpy_bulkmem(m, v33, l0, v27)
											mBase = m.M
										} else {
										}
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56 + v27
										return
									}
								}
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
func F_datum_compute_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	v8 = base.B2i32(l3 != int32(-1))
	if l3 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v143 + v145
L2:
	;
	if l1&int32(3) == int32(0) {
		v107 = l1
		goto L31
	} else {
		goto L32
	}
L3:
	;
	if v49&int32(255) == int32(1) {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	switch l2 - int32(99) {
	case 0:
		v42 = l0
		goto L13
	case 1:
		goto L15
	default:
		goto L14
	case 6:
		goto L16
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v24&int32(1) != 0 {
		v48 = l0
		v49 = v24
		goto L3
	} else {
		goto L12
	}
L6:
	;
	if l3 != int32(-1) {
		goto L4
	} else {
		goto L11
	}
L7:
	;
	if l4 == int32(112) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v11&int32(3) != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(base.Ui32(v14)>>(uint(int32(2))%32)) - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v18) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	return l0 + v18
L11:
	;
	goto L5
L12:
	;
	goto L4
L13:
	;
	if int32(0) < l3 {
		v143 = v42
		v145 = l3
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v42 = (l0 + int32(1)) & int32(-2)
	goto L13
L15:
	;
	v42 = (l0 + int32(7)) & int32(-8)
	goto L13
L16:
	;
	v42 = (l0 + int32(3)) & int32(-4)
	goto L13
L17:
	;
	if l3 != int32(-1) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v48 = v42
	v49 = v47
	goto L3
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v143 = v48
		v145 = int32(6)
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v49&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v62 = int32(18)
	if v55&int32(255) == v62 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = v62
	goto L25
L24:
	;
	v68 = int32(2)
	goto L25
L25:
	;
	return v68 + v48
L26:
	;
	return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
L27:
	;
	goto L28
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
L29:
	;
	v143 = v42
	v145 = v140 + int32(1)
	goto L1
L30:
	;
	v140 = v132 - l1
	goto L29
L31:
	;
	v111 = v107
	goto L40
L32:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v91 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v140 = int32(0)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v96 = l1
	goto L36
L36:
	;
	v100 = v96 + int32(1)
	if v100&int32(3) == int32(0) {
		v107 = v100
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v132 = v100
	goto L30
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v105 != 0 {
		v96 = v100
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v120 = int32(-2139062144)
	if (int32(16843008)-v117|v117)&v120 == v120 {
		v111 = v111 + int32(4)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v126 = v111
	goto L43
L42:
	;
	goto L41
L43:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 != 0 {
		v126 = v126 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v132 = v126
	goto L30
L45:
	;
	goto L44
}
