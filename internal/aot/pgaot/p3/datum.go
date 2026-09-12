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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v8 = base.B2i32(l3 != int32(-1))
	if l3 != int32(-1) {
		if l3 != int32(-1) {
			switch l2 - int32(99) {
			case 0:
				v42 = l0
			case 1:
				v42 = (l0 + int32(7)) & int32(-8)
			default:
				v42 = (l0 + int32(1)) & int32(-2)
			case 6:
				v42 = (l0 + int32(3)) & int32(-4)
			}
			if int32(0) < l3 {
				v87 = v42
				v89 = l3
				return v87 + v89
			} else {
				if l3 != int32(-1) {
					v84 = F_strlen(m, l1)
					mBase = m.M
					v87 = v42
					v89 = v84 + int32(1)
					return v87 + v89
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v48 = v42
					v49 = v47
					if v49&int32(255) == int32(1) {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
						if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v87 = v48
							v89 = int32(6)
							return v87 + v89
						} else {
							v62 = int32(18)
							if v55&int32(255) == v62 {
								v68 = v62
							} else {
								v68 = int32(2)
							}
							return v68 + v48
						}
					} else {
						if v49&int32(1) != 0 {
							return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
						}
					}
				}
			}
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v24&int32(1) != 0 {
				v48 = l0
				v49 = v24
				if v49&int32(255) == int32(1) {
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v87 = v48
						v89 = int32(6)
						return v87 + v89
					} else {
						v62 = int32(18)
						if v55&int32(255) == v62 {
							v68 = v62
						} else {
							v68 = int32(2)
						}
						return v68 + v48
					}
				} else {
					if v49&int32(1) != 0 {
						return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
					}
				}
			} else {
				switch l2 - int32(99) {
				case 0:
					v42 = l0
				case 1:
					v42 = (l0 + int32(7)) & int32(-8)
				default:
					v42 = (l0 + int32(1)) & int32(-2)
				case 6:
					v42 = (l0 + int32(3)) & int32(-4)
				}
				if int32(0) < l3 {
					v87 = v42
					v89 = l3
					return v87 + v89
				} else {
					if l3 != int32(-1) {
						v84 = F_strlen(m, l1)
						mBase = m.M
						v87 = v42
						v89 = v84 + int32(1)
						return v87 + v89
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v48 = v42
						v49 = v47
						if v49&int32(255) == int32(1) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v87 = v48
								v89 = int32(6)
								return v87 + v89
							} else {
								v62 = int32(18)
								if v55&int32(255) == v62 {
									v68 = v62
								} else {
									v68 = int32(2)
								}
								return v68 + v48
							}
						} else {
							if v49&int32(1) != 0 {
								return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
							}
						}
					}
				}
			}
		}
	} else {
		if l4 == int32(112) {
			if l3 != int32(-1) {
				switch l2 - int32(99) {
				case 0:
					v42 = l0
				case 1:
					v42 = (l0 + int32(7)) & int32(-8)
				default:
					v42 = (l0 + int32(1)) & int32(-2)
				case 6:
					v42 = (l0 + int32(3)) & int32(-4)
				}
				if int32(0) < l3 {
					v87 = v42
					v89 = l3
					return v87 + v89
				} else {
					if l3 != int32(-1) {
						v84 = F_strlen(m, l1)
						mBase = m.M
						v87 = v42
						v89 = v84 + int32(1)
						return v87 + v89
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v48 = v42
						v49 = v47
						if v49&int32(255) == int32(1) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v87 = v48
								v89 = int32(6)
								return v87 + v89
							} else {
								v62 = int32(18)
								if v55&int32(255) == v62 {
									v68 = v62
								} else {
									v68 = int32(2)
								}
								return v68 + v48
							}
						} else {
							if v49&int32(1) != 0 {
								return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
							}
						}
					}
				}
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v24&int32(1) != 0 {
					v48 = l0
					v49 = v24
					if v49&int32(255) == int32(1) {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
						if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v87 = v48
							v89 = int32(6)
							return v87 + v89
						} else {
							v62 = int32(18)
							if v55&int32(255) == v62 {
								v68 = v62
							} else {
								v68 = int32(2)
							}
							return v68 + v48
						}
					} else {
						if v49&int32(1) != 0 {
							return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
						}
					}
				} else {
					switch l2 - int32(99) {
					case 0:
						v42 = l0
					case 1:
						v42 = (l0 + int32(7)) & int32(-8)
					default:
						v42 = (l0 + int32(1)) & int32(-2)
					case 6:
						v42 = (l0 + int32(3)) & int32(-4)
					}
					if int32(0) < l3 {
						v87 = v42
						v89 = l3
						return v87 + v89
					} else {
						if l3 != int32(-1) {
							v84 = F_strlen(m, l1)
							mBase = m.M
							v87 = v42
							v89 = v84 + int32(1)
							return v87 + v89
						} else {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							v48 = v42
							v49 = v47
							if v49&int32(255) == int32(1) {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
								if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v87 = v48
									v89 = int32(6)
									return v87 + v89
								} else {
									v62 = int32(18)
									if v55&int32(255) == v62 {
										v68 = v62
									} else {
										v68 = int32(2)
									}
									return v68 + v48
								}
							} else {
								if v49&int32(1) != 0 {
									return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
								}
							}
						}
					}
				}
			}
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v11&int32(3) != 0 {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v24&int32(1) != 0 {
					v48 = l0
					v49 = v24
					if v49&int32(255) == int32(1) {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
						if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v87 = v48
							v89 = int32(6)
							return v87 + v89
						} else {
							v62 = int32(18)
							if v55&int32(255) == v62 {
								v68 = v62
							} else {
								v68 = int32(2)
							}
							return v68 + v48
						}
					} else {
						if v49&int32(1) != 0 {
							return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
						}
					}
				} else {
					switch l2 - int32(99) {
					case 0:
						v42 = l0
					case 1:
						v42 = (l0 + int32(7)) & int32(-8)
					default:
						v42 = (l0 + int32(1)) & int32(-2)
					case 6:
						v42 = (l0 + int32(3)) & int32(-4)
					}
					if int32(0) < l3 {
						v87 = v42
						v89 = l3
						return v87 + v89
					} else {
						if l3 != int32(-1) {
							v84 = F_strlen(m, l1)
							mBase = m.M
							v87 = v42
							v89 = v84 + int32(1)
							return v87 + v89
						} else {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							v48 = v42
							v49 = v47
							if v49&int32(255) == int32(1) {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
								if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v87 = v48
									v89 = int32(6)
									return v87 + v89
								} else {
									v62 = int32(18)
									if v55&int32(255) == v62 {
										v68 = v62
									} else {
										v68 = int32(2)
									}
									return v68 + v48
								}
							} else {
								if v49&int32(1) != 0 {
									return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
								}
							}
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v18 = int32(base.Ui32(v14)>>(uint(int32(2))%32)) - int32(3)
				if base.Ui32(int32(127)) < base.Ui32(v18) {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					if v24&int32(1) != 0 {
						v48 = l0
						v49 = v24
						if v49&int32(255) == int32(1) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v87 = v48
								v89 = int32(6)
								return v87 + v89
							} else {
								v62 = int32(18)
								if v55&int32(255) == v62 {
									v68 = v62
								} else {
									v68 = int32(2)
								}
								return v68 + v48
							}
						} else {
							if v49&int32(1) != 0 {
								return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
							}
						}
					} else {
						switch l2 - int32(99) {
						case 0:
							v42 = l0
						case 1:
							v42 = (l0 + int32(7)) & int32(-8)
						default:
							v42 = (l0 + int32(1)) & int32(-2)
						case 6:
							v42 = (l0 + int32(3)) & int32(-4)
						}
						if int32(0) < l3 {
							v87 = v42
							v89 = l3
							return v87 + v89
						} else {
							if l3 != int32(-1) {
								v84 = F_strlen(m, l1)
								mBase = m.M
								v87 = v42
								v89 = v84 + int32(1)
								return v87 + v89
							} else {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
								v48 = v42
								v49 = v47
								if v49&int32(255) == int32(1) {
									v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
									if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v87 = v48
										v89 = int32(6)
										return v87 + v89
									} else {
										v62 = int32(18)
										if v55&int32(255) == v62 {
											v68 = v62
										} else {
											v68 = int32(2)
										}
										return v68 + v48
									}
								} else {
									if v49&int32(1) != 0 {
										return int32(base.Ui32(v49&int32(254))>>(uint(int32(1))%32)) + v48
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										return int32(base.Ui32(v79)>>(uint(int32(2))%32)) + v48
									}
								}
							}
						}
					}
				} else {
					return l0 + v18
				}
			}
		}
	}
}
